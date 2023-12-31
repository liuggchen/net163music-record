package u

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/asmcos/requests"
)

type Options struct {
	Crypto  string
	Ua      string
	Cookies []*http.Cookie
	Token   string
	Url     string
}

func chooseUserAgent(ua string) string {
	userAgentList := []string{
		"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
		"Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 5.1.1; Nexus 6 Build/LYZ28E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89;GameHelper",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
		"Mozilla/5.0 (iPad; CPU OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:46.0) Gecko/20100101 Firefox/46.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:46.0) Gecko/20100101 Firefox/46.0",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/13.10586",
	}

	rand.Seed(time.Now().UnixNano())
	index := 0
	if ua == "" {
		index = rand.Intn(len(userAgentList))
	} else if ua == "mobile" {
		index = rand.Intn(8)
	} else {
		index = rand.Intn(7) + 7
	}
	return userAgentList[index]
}

func CreateRequest(method string, url string, data map[string]string, options *Options, resPoint interface{}) ([]*http.Cookie, error) {
	req := requests.Requests()
	req.Header.Set("User-Agent", chooseUserAgent(options.Ua))
	csrfToken := ""
	music_U := ""

	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if strings.Contains(url, "music.163.com") {
		req.Header.Set("Referer", "https://music.163.com")
	}
	if options.Cookies != nil {
		for _, cookie := range options.Cookies {
			req.SetCookie(cookie)
			if cookie.Name == "__csrf" {
				csrfToken = cookie.Value
			}
			if cookie.Name == "MUSIC_U" {
				music_U = cookie.Value
			}
		}
	}
	if options.Crypto == "weapi" {
		data["csrf_token"] = csrfToken
		data = Weapi(data)
		reg, _ := regexp.Compile(`/\w*api/`)
		url = reg.ReplaceAllString(url, "/weapi/")
	} else if options.Crypto == "linuxapi" {
		linuxApiData := make(map[string]interface{}, 3)
		linuxApiData["method"] = method
		reg, _ := regexp.Compile(`/\w*api/`)
		linuxApiData["url"] = reg.ReplaceAllString(url, "/api/")
		linuxApiData["params"] = data
		data = Linuxapi(linuxApiData)
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36")
		url = "https://music.163.com/api/linux/forward"
	} else if options.Crypto == "eapi" {
		eapiData := make(map[string]interface{})
		for key, value := range data {
			eapiData[key] = value
		}
		rand.Seed(time.Now().UnixNano())
		header := map[string]string{
			"osver":       "",
			"deviceId":    "",
			"mobilename":  "",
			"appver":      "6.1.1",
			"versioncode": "140",
			"buildver":    strconv.FormatInt(time.Now().Unix(), 10),
			"resolution":  "1920x1080",
			"os":          "android",
			"channel":     "",
			"requestId":   strconv.FormatInt(time.Now().Unix()*1000, 10) + strconv.Itoa(rand.Intn(1000)),
			"MUSIC_U":     music_U,
		}

		for key, value := range header {
			req.SetCookie(&http.Cookie{Name: key, Value: value, Path: "/"})
		}
		eapiData["header"] = header
		data = Eapi(options.Url, eapiData)
		reg, _ := regexp.Compile(`/\w*api/`)
		url = reg.ReplaceAllString(url, "/eapi/")
	}
	var resp *requests.Response
	var err error
	if method == "POST" {
		var form requests.Datas = data
		resp, err = req.Post(url, form)
	} else {
		resp, err = req.Get(url)
	}

	if err != nil {
		return nil, err
	}
	cookies := resp.Cookies()

	body := resp.Content()
	b := bytes.NewReader(body)
	var out bytes.Buffer
	r, err := zlib.NewReader(b)
	// 数据被压缩 进行解码
	if err == nil {
		io.Copy(&out, r)
		body = out.Bytes()
	}

	if err := json.Unmarshal(body, resPoint); err != nil {
		return nil, fmt.Errorf("json decode error: %w, body: %s", err, body)
	}

	return cookies, nil
}
