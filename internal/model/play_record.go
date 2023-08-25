package model

type ItemDayData struct {
	Id    int
	Name  string
	Count int
}

type PlayRecord struct {
	AllData  []AllDataItem  `json:"allData"`
	WeekData []WeekDataItem `json:"weekData"`
	Code     int            `json:"code"`
}

type AllDataItem struct {
	PlayCount int `json:"playCount"`
	Score     int `json:"score"`
	Song      struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
		Pst  int    `json:"pst"`
		T    int    `json:"t"`
		Ar   []struct {
			Id    int           `json:"id"`
			Name  string        `json:"name"`
			Tns   []interface{} `json:"tns"`
			Alias []interface{} `json:"alias"`
		} `json:"ar"`
		Alia []string    `json:"alia"`
		Pop  float64     `json:"pop"`
		St   int         `json:"st"`
		Rt   *string     `json:"rt"`
		Fee  int         `json:"fee"`
		V    int         `json:"v"`
		Crbt interface{} `json:"crbt"`
		Cf   string      `json:"cf"`
		Al   struct {
			Id     int      `json:"id"`
			Name   string   `json:"name"`
			PicUrl string   `json:"picUrl"`
			Tns    []string `json:"tns"`
			PicStr string   `json:"pic_str,omitempty"`
			Pic    int64    `json:"pic"`
		} `json:"al"`
		Dt int `json:"dt"`
		H  struct {
			Br   int     `json:"br"`
			Fid  int     `json:"fid"`
			Size int     `json:"size"`
			Vd   float64 `json:"vd"`
		} `json:"h"`
		M struct {
			Br   int     `json:"br"`
			Fid  int     `json:"fid"`
			Size int     `json:"size"`
			Vd   float64 `json:"vd"`
		} `json:"m"`
		L struct {
			Br   int     `json:"br"`
			Fid  int     `json:"fid"`
			Size int     `json:"size"`
			Vd   float64 `json:"vd"`
		} `json:"l"`
		A                    interface{}   `json:"a"`
		Cd                   string        `json:"cd"`
		No                   int           `json:"no"`
		RtUrl                interface{}   `json:"rtUrl"`
		Ftype                int           `json:"ftype"`
		RtUrls               []interface{} `json:"rtUrls"`
		DjId                 int           `json:"djId"`
		Copyright            int           `json:"copyright"`
		SId                  int           `json:"s_id"`
		Mark                 int64         `json:"mark"`
		OriginCoverType      int           `json:"originCoverType"`
		OriginSongSimpleData *struct {
			SongId  int    `json:"songId"`
			Name    string `json:"name"`
			Artists []struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
			} `json:"artists"`
			AlbumMeta struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
			} `json:"albumMeta"`
		} `json:"originSongSimpleData"`
		Single          int         `json:"single"`
		NoCopyrightRcmd interface{} `json:"noCopyrightRcmd"`
		Rurl            interface{} `json:"rurl"`
		Rtype           int         `json:"rtype"`
		Mst             int         `json:"mst"`
		Cp              int         `json:"cp"`
		Mv              int         `json:"mv"`
		PublishTime     int64       `json:"publishTime"`
		Privilege       struct {
			Id                 int         `json:"id"`
			Fee                int         `json:"fee"`
			Payed              int         `json:"payed"`
			St                 int         `json:"st"`
			Pl                 int         `json:"pl"`
			Dl                 int         `json:"dl"`
			Sp                 int         `json:"sp"`
			Cp                 int         `json:"cp"`
			Subp               int         `json:"subp"`
			Cs                 bool        `json:"cs"`
			Maxbr              int         `json:"maxbr"`
			Fl                 int         `json:"fl"`
			Toast              bool        `json:"toast"`
			Flag               int         `json:"flag"`
			PreSell            bool        `json:"preSell"`
			PlayMaxbr          int         `json:"playMaxbr"`
			DownloadMaxbr      int         `json:"downloadMaxbr"`
			MaxBrLevel         string      `json:"maxBrLevel"`
			PlayMaxBrLevel     string      `json:"playMaxBrLevel"`
			DownloadMaxBrLevel string      `json:"downloadMaxBrLevel"`
			PlLevel            string      `json:"plLevel"`
			DlLevel            string      `json:"dlLevel"`
			FlLevel            string      `json:"flLevel"`
			Rscl               interface{} `json:"rscl"`
			FreeTrialPrivilege struct {
				ResConsumable  bool        `json:"resConsumable"`
				UserConsumable bool        `json:"userConsumable"`
				ListenType     interface{} `json:"listenType"`
			} `json:"freeTrialPrivilege"`
			ChargeInfoList []struct {
				Rate          int         `json:"rate"`
				ChargeUrl     interface{} `json:"chargeUrl"`
				ChargeMessage interface{} `json:"chargeMessage"`
				ChargeType    int         `json:"chargeType"`
			} `json:"chargeInfoList"`
		} `json:"privilege"`
		Tns []string `json:"tns,omitempty"`
		Eq  string   `json:"eq,omitempty"`
	} `json:"song"`
}

type WeekDataItem struct {
	PlayCount int `json:"playCount"`
	Score     int `json:"score"`
	Song      struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
		Pst  int    `json:"pst"`
		T    int    `json:"t"`
		Ar   []struct {
			Id    int           `json:"id"`
			Name  string        `json:"name"`
			Tns   []interface{} `json:"tns"`
			Alias []interface{} `json:"alias"`
		} `json:"ar"`
		Alia []string    `json:"alia"`
		Pop  float64     `json:"pop"`
		St   int         `json:"st"`
		Rt   *string     `json:"rt"`
		Fee  int         `json:"fee"`
		V    int         `json:"v"`
		Crbt interface{} `json:"crbt"`
		Cf   string      `json:"cf"`
		Al   struct {
			Id     int      `json:"id"`
			Name   string   `json:"name"`
			PicUrl string   `json:"picUrl"`
			Tns    []string `json:"tns"`
			PicStr string   `json:"pic_str,omitempty"`
			Pic    int64    `json:"pic"`
		} `json:"al"`
		Dt int `json:"dt"`
		H  struct {
			Br   int     `json:"br"`
			Fid  int     `json:"fid"`
			Size int     `json:"size"`
			Vd   float64 `json:"vd"`
		} `json:"h"`
		M struct {
			Br   int     `json:"br"`
			Fid  int     `json:"fid"`
			Size int     `json:"size"`
			Vd   float64 `json:"vd"`
		} `json:"m"`
		L struct {
			Br   int     `json:"br"`
			Fid  int     `json:"fid"`
			Size int     `json:"size"`
			Vd   float64 `json:"vd"`
		} `json:"l"`
		A                    interface{}   `json:"a"`
		Cd                   string        `json:"cd"`
		No                   int           `json:"no"`
		RtUrl                interface{}   `json:"rtUrl"`
		Ftype                int           `json:"ftype"`
		RtUrls               []interface{} `json:"rtUrls"`
		DjId                 int           `json:"djId"`
		Copyright            int           `json:"copyright"`
		SId                  int           `json:"s_id"`
		Mark                 int64         `json:"mark"`
		OriginCoverType      int           `json:"originCoverType"`
		OriginSongSimpleData *struct {
			SongId  int    `json:"songId"`
			Name    string `json:"name"`
			Artists []struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
			} `json:"artists"`
			AlbumMeta struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
			} `json:"albumMeta"`
		} `json:"originSongSimpleData"`
		Single          int         `json:"single"`
		NoCopyrightRcmd interface{} `json:"noCopyrightRcmd"`
		Rtype           int         `json:"rtype"`
		Rurl            interface{} `json:"rurl"`
		Mst             int         `json:"mst"`
		Cp              int         `json:"cp"`
		Mv              int         `json:"mv"`
		PublishTime     int64       `json:"publishTime"`
		Privilege       struct {
			Id                 int         `json:"id"`
			Fee                int         `json:"fee"`
			Payed              int         `json:"payed"`
			St                 int         `json:"st"`
			Pl                 int         `json:"pl"`
			Dl                 int         `json:"dl"`
			Sp                 int         `json:"sp"`
			Cp                 int         `json:"cp"`
			Subp               int         `json:"subp"`
			Cs                 bool        `json:"cs"`
			Maxbr              int         `json:"maxbr"`
			Fl                 int         `json:"fl"`
			Toast              bool        `json:"toast"`
			Flag               int         `json:"flag"`
			PreSell            bool        `json:"preSell"`
			PlayMaxbr          int         `json:"playMaxbr"`
			DownloadMaxbr      int         `json:"downloadMaxbr"`
			MaxBrLevel         string      `json:"maxBrLevel"`
			PlayMaxBrLevel     string      `json:"playMaxBrLevel"`
			DownloadMaxBrLevel string      `json:"downloadMaxBrLevel"`
			PlLevel            string      `json:"plLevel"`
			DlLevel            string      `json:"dlLevel"`
			FlLevel            string      `json:"flLevel"`
			Rscl               interface{} `json:"rscl"`
			FreeTrialPrivilege struct {
				ResConsumable  bool        `json:"resConsumable"`
				UserConsumable bool        `json:"userConsumable"`
				ListenType     interface{} `json:"listenType"`
			} `json:"freeTrialPrivilege"`
			ChargeInfoList []struct {
				Rate          int         `json:"rate"`
				ChargeUrl     interface{} `json:"chargeUrl"`
				ChargeMessage interface{} `json:"chargeMessage"`
				ChargeType    int         `json:"chargeType"`
			} `json:"chargeInfoList"`
		} `json:"privilege"`
		Tns []string `json:"tns,omitempty"`
		Eq  string   `json:"eq,omitempty"`
	} `json:"song"`
}
