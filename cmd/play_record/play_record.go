package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"net163music-record/internal/model"
	"net163music-record/internal/u"
)

func init() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = location
}

func main() {

	bytes, err := os.ReadFile("uid.txt")
	if err != nil {
		log.Fatalf("read uid.txt error: %s", err.Error())
	}

	User := &UserRecordService{UId: string(bytes)}
	r, err := User.All()
	if err != nil {
		log.Fatalf("err: %s", err.Error())
	}

	file, err := os.Create("data/" + time.Now().Format("20060102") + ".txt")
	if err != nil {
		log.Fatalf("create file error: %s", err.Error())
	}
	defer file.Close()

	var todayData = make(map[int]model.ItemDayData)
	for _, item := range r {
		file.WriteString(fmt.Sprintf("%-5d%-15d%s\n", item.PlayCount, item.Song.Id, item.Song.Name))
		todayData[item.Song.Id] = model.ItemDayData{
			Id:    item.Song.Id,
			Name:  item.Song.Name,
			Count: item.PlayCount,
		}
	}

	yFile, err := os.Open("data/" + time.Now().AddDate(0, 0, -1).Format("20060102") + ".txt")
	if err != nil {
		log.Fatalf("open yesterday file error: %s", err.Error())
	}
	reader := bufio.NewReader(yFile)

	var yesterdayData = make(map[int]model.ItemDayData)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("readLine error: %s", err.Error())
		}
		newLine := spaceReg.ReplaceAllString(string(line), " ")
		split := strings.SplitN(newLine, " ", 3)
		c, _ := strconv.Atoi(split[0])
		musicId, _ := strconv.Atoi(split[1])
		yesterdayData[musicId] = model.ItemDayData{
			Id:    musicId,
			Name:  split[2],
			Count: c,
		}
	}

	var todayNewItems []model.ItemDayData
	var todayMoreItems []model.ItemDayData
	for mId, item := range todayData {
		if v, ok := yesterdayData[mId]; !ok {
			todayNewItems = append(todayNewItems, item)
		} else if item.Count > v.Count {
			todayMoreItems = append(todayMoreItems, model.ItemDayData{
				Id:    mId,
				Name:  item.Name,
				Count: item.Count - v.Count,
			})
		}
	}

	newFile, err := os.OpenFile("data/new.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("create file error: %s", err.Error())
	}
	defer newFile.Close()

	log.SetOutput(io.MultiWriter(newFile, os.Stdout))
	log.SetFlags(log.Lmsgprefix)

	log.Printf("====== %s ======\n", time.Now().Format("2006-01-02 15:04:05"))
	log.Printf("新增歌曲： %d首\n", len(todayNewItems))
	if len(todayNewItems) > 0 {
		sort.Slice(todayNewItems, func(i, j int) bool {
			return todayNewItems[i].Count > todayNewItems[j].Count
		})
	}
	for _, item := range todayNewItems {
		log.Printf("%-5d%-15d %s\n", item.Count, item.Id, item.Name)
	}

	log.Printf("今日接着听歌曲： %d首\n", len(todayMoreItems))
	if len(todayMoreItems) > 0 {
		sort.Slice(todayMoreItems, func(i, j int) bool {
			return todayMoreItems[i].Count > todayMoreItems[j].Count
		})
	}
	for _, item := range todayMoreItems {
		log.Printf("%-5d%-15d %s\n", item.Count, item.Id, item.Name)
	}
	log.Printf("")

}

var spaceReg = regexp.MustCompile(" +")

type UserRecordService struct {
	UId string `json:"uid" form:"uid"`
}

func (s *UserRecordService) Week() ([]model.WeekDataItem, error) {
	var res = new(model.PlayRecord)
	if err := s.req("1", res); err != nil {
		return nil, err
	}
	return res.WeekData, nil
}

func (s *UserRecordService) All() ([]model.AllDataItem, error) {
	var res = new(model.PlayRecord)
	if err := s.req("0", res); err != nil {
		return nil, err
	}
	return res.AllData, nil
}

func (s *UserRecordService) req(t string, resPoint interface{}) error {
	options := &u.Options{
		Crypto: "weapi",
		//Cookies: cookies,
	}
	data := map[string]string{
		"uid":  s.UId,
		"type": t,
	}
	_, err := u.CreateRequest("POST", `https://music.163.com/weapi/v1/play/record`, data, options, resPoint)
	if err != nil {
		return err
	}
	return nil
}
