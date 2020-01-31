package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"log"
	"time"
)

//存放判断重复信息的关键信息
type History struct {
	ChinaUpdateTime    int64    `json:"chinaUpdateTime"`    // 全国概览数数据更新时间
	ProvinceUpdateTime int64    `json:"provinceUpdateTime"` // 省份详细数据更新时间
	RumorsId           int64    `json:"rumorsId"`           // 辟谣信息ID
	PubDate            int64    `json:"pubDate"`            // 新闻的发布时间
	CityData           CityData `json:"cityData"`           // 城市数据

}

// 省份详细数据响应信息
type ProvinceResponse struct {
	Results []ProvinceData `json:"results"`
	Success bool           `json:"success"`
}

// 省份详细数据
type ProvinceData struct {
	ProvinceName      string     `json:"provinceName"`
	ProvinceShortName string     `json:"provinceShortName"`
	ConfirmedCount    int        `json:"confirmedCount"`
	SuspectedCount    int        `json:"suspectedCount"`
	CuredCount        int        `json:"curedCount"`
	DeadCount         int        `json:"deadCount"`
	Cities            []CityData `json:"cities"`
	UpdateTime        int64      `json:"updateTime"`
}

// 城市详细数据
type CityData struct {
	CityName       string `json:"cityName"`
	ConfirmedCount int64  `json:"confirmedCount"`
	SuspectedCount int64  `json:"suspectedCount"`
	CuredCount     int64  `json:"curedCount"`
	DeadCount      int64  `json:"deadCount"`
}

// 返回信息
type ReturnMsg struct {
	Title   string      `json:"title"`   // 系统默认提供的title
	Content string      `json:"content"` // 内容部分
	Id      interface{} `json:"id"`      // 唯一的标识
}

// 解析并发送全国概览数据
func parseChinaOverview(jsonStr string) (ReturnMsg, bool) {

	var rsp ReturnMsg

	confirmedCount := gjson.Get(jsonStr, "results.0.confirmedCount").Int()
	suspectedCount := gjson.Get(jsonStr, "results.0.suspectedCount").Int()
	curedCount := gjson.Get(jsonStr, "results.0.curedCount").Int()
	deadCount := gjson.Get(jsonStr, "results.0.deadCount").Int()
	updateTime := gjson.Get(jsonStr, "results.0.updateTime").Int()
	dailyPic := gjson.Get(jsonStr, "results.0.dailyPic").String()

	if updateTime <= history.ChinaUpdateTime {
		log.Println("已发送过全国概览数据,不再发送")
		return rsp, false
	} else {
		t := time.Unix(updateTime/1e3, 0)
		updateTime1 := t.Format("2006-01-02 15:04:05")

		content := fmt.Sprintf(chinaOverview_template, updateTime1, confirmedCount, suspectedCount, curedCount, deadCount, dailyPic)
		log.Println("待发送全国概览数据的内容为:", content)
		rsp.Title = "全国概况数据更新"
		rsp.Content = content
		rsp.Id = updateTime
		return rsp, true
	}

}

// 解析并发送辟谣信息
func parseRumors(jsonStr string) (ReturnMsg, bool) {
	var rsp ReturnMsg
	id := gjson.Get(jsonStr, "results.0.id").Int()
	title := gjson.Get(jsonStr, "results.0.title").String()
	mainSummary := gjson.Get(jsonStr, "results.0.mainSummary").String()
	body := gjson.Get(jsonStr, "results.0.body").String()

	// 判断是否已经发送过此id
	if id <= history.RumorsId {
		log.Println("已发送过此ID辟谣信息:", title)
		return rsp, false
	} else {
		content := fmt.Sprintf(rumors_template, mainSummary, body)
		log.Println("标题为:", title, "待发送辟谣内容为:", content)

		rsp.Content = content
		rsp.Title = title
		rsp.Id = id
		return rsp, true
	}

}

// 解析并发送省份新闻
func parseProvinceNews(jsonStr string) (ReturnMsg, bool) {
	var rsp ReturnMsg

	pubDate := gjson.Get(jsonStr, "results.0.pubDate").Int()
	title := gjson.Get(jsonStr, "results.0.title").String()
	summary := gjson.Get(jsonStr, "results.0.summary").String()
	infoSource := gjson.Get(jsonStr, "results.0.infoSource").String()
	sourceUrl := gjson.Get(jsonStr, "results.0.sourceUrl").String()

	// 判断是否已经发送过
	if pubDate <= history.PubDate {
		log.Println("已发送过此省份消息:", title)
		return rsp, false
	}
	t := time.Unix(pubDate/1e3, 0)
	updateTime := t.Format("2006-01-02 15:04:05")

	content := fmt.Sprintf(provinceNews_template, summary, updateTime, infoSource, sourceUrl)
	log.Println("标题为:", title, "待发送省份新闻内容为:", content)

	rsp.Id = pubDate
	rsp.Content = content
	rsp.Title = title
	return rsp, true
}

// 解析并发送所有区域的新闻信息
func parseAreaNews(jsonStr string) (ReturnMsg, bool) {
	var rsp ReturnMsg

	pubDate := gjson.Get(jsonStr, "results.0.pubDate").Int()
	title := gjson.Get(jsonStr, "results.0.title").String()
	summary := gjson.Get(jsonStr, "results.0.summary").String()
	infoSource := gjson.Get(jsonStr, "results.0.infoSource").String()
	sourceUrl := gjson.Get(jsonStr, "results.0.sourceUrl").String()

	// 判断是否已经发送过
	if pubDate <= history.PubDate {
		log.Println("已发送过此区域消息:", title)
		return rsp, false
	} else {
		t := time.Unix(pubDate/1e3, 0)
		updateTime := t.Format("2006-01-02 15:04:05")

		content := fmt.Sprintf(provinceNews_template, summary, updateTime, infoSource, sourceUrl)
		log.Println("标题为:", title, "待发送全部区域新闻内容为:", content)

		rsp.Id = pubDate
		rsp.Content = content
		rsp.Title = title
		return rsp, true
	}

}

// 解析并发送指定城市的概况信息
func parseCityOverviewData(jsonStr string) (ReturnMsg, bool) {
	var rsp ReturnMsg

	cityName := gjson.Get(jsonStr, "results.0.cityName").String()
	confirmedCount := gjson.Get(jsonStr, "results.0.confirmedCount").Int()
	suspectedCount := gjson.Get(jsonStr, "results.0.suspectedCount").Int()
	curedCount := gjson.Get(jsonStr, "results.0.curedCount").Int()
	deadCount := gjson.Get(jsonStr, "results.0.deadCount").Int()

	cityData := CityData{
		CityName:       cityName,
		ConfirmedCount: confirmedCount,
		SuspectedCount: suspectedCount,
		CuredCount:     curedCount,
		DeadCount:      deadCount,
	}

	// 判断是否已经发送过
	if cityData == history.CityData {
		log.Println("已发送过此区域消息:", cityName)
		return rsp, false
	} else {

		content := fmt.Sprintf(cityOverviewData_template, cityName, confirmedCount, suspectedCount, curedCount, deadCount)
		log.Println("待发送城市概况内容为:", content)

		rsp.Content = content
		rsp.Title = cityName + "统计数据更新"
		rsp.Id = cityData

		return rsp, true
	}

}

// 解析并发送指定省份的详细信息
func parseChinaProvincesData(jsonStr string) (ReturnMsg, bool) {
	var rsp ReturnMsg
	var provinceResponse ProvinceResponse
	if err := json.Unmarshal([]byte(jsonStr), &provinceResponse); err != nil {
		log.Println("数据结构不正确", err.Error())
		return rsp, false
	} else {
		// 存放完整的content
		content := ""
		provinceData := provinceResponse.Results[0]

		pname := provinceData.ProvinceName
		confirmedCount := provinceData.ConfirmedCount
		suspectedCount := provinceData.SuspectedCount
		curedCount := provinceData.CuredCount
		deadCount := provinceData.DeadCount
		updateTime := provinceData.UpdateTime

		// 判断是否发送过此消息
		if updateTime <= history.ProvinceUpdateTime {
			log.Println("已发送过此省份的详细信息:", pname)
			return rsp, false
		} else {
			overviewData := fmt.Sprintf(provinceOverviewData_template, pname, confirmedCount, suspectedCount, curedCount, deadCount)
			content += overviewData
			content += City_Header_template
			cities := provinceData.Cities
			// 循环添加各城市信息
			for _, c := range cities {
				cityName := c.CityName
				confirmedCount := c.ConfirmedCount
				suspectedCount := c.SuspectedCount
				curedCount := c.CuredCount
				deadCount := c.DeadCount
				cityData := fmt.Sprintf(singleProvinceOrCity, cityName, confirmedCount, suspectedCount, curedCount, deadCount)
				content += cityData
			}

			log.Println("待发送省份详情数据内容为:", content)
			rsp.Title = pname + "详情数据更新"
			rsp.Content = content
			rsp.Id = updateTime
			return rsp, true
		}
	}
}
