package main

import (
	"log"
	"time"
)

// 订阅谣言信息，参数title,自定义微信接收到的主题名称，建议用默认值
func SubRumors(title string, interval int) {
	for {
		pushRumors(title)
		time.Sleep(time.Duration(interval) * time.Second)
	}

}

// 订阅辟谣信息，参数title,自定义微信接收到的主题名称，建议用默认值
func SubNews(area string, title string, interval int) {
	if area == "all" || area == "default" || area == "" {
		for {
			pushAreaNews(title)
			time.Sleep(time.Duration(interval) * time.Second)
		}
	} else { // 发送省份新闻
		for {
			pushProvinceNews(area, title)
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}

}

// 订阅全国数据
func SubChina(title string, interval int) {
	for {
		pushChinaOverview(title)
		time.Sleep(time.Duration(interval) * time.Second)
	}

}

// 订阅省份详细信息，参数title,自定义微信接收到的主题名称，建议用默认值
func SubProvince(provinceName string, title string, interval int) {
	for {
		pushProvinceOverview(provinceName, title)
		time.Sleep(time.Duration(interval) * time.Second)
	}

}

// 订阅城市数据
func SubCity(cityName string, title string, interval int) {
	for {
		pushCityOverview(cityName, title)
		time.Sleep(time.Duration(interval) * time.Second)
	}

}

// 推送辟谣信息
func pushRumors(title string) {
	if str, err := GetRumors(); err != nil {
		log.Println("从接口获取数据失败", err.Error())
		return
	} else {
		if rsp, ok := parseRumors(str); ok {
			if title == "default" || title == "" {
				title = rsp.Title
			}
			if err := sendToWechat(title, rsp.Content); err == nil {
				history.RumorsId = rsp.Id.(int64)
			}
		}
	}
}

// 推送所有的新闻消息
func pushAreaNews(title string) {
	if str, err := GetAreaNews(); err != nil {
		log.Println("从接口获取数据失败", err.Error())
		return
	} else {
		if rsp, ok := parseAreaNews(str); ok {
			if title == "default" || title == "" {
				title = rsp.Title
			}
			if err := sendToWechat(title, rsp.Content); err == nil {
				history.PubDate = rsp.Id.(int64)
			}
		}
	}
}

// 推送省份新闻消息
func pushProvinceNews(provinceName, title string) {
	if str, err := GetProvinceNews(provinceName); err != nil {
		log.Println("从接口获取数据失败", err.Error())
		return
	} else {
		if rsp, ok := parseProvinceNews(str); ok {
			if title == "default" || title == "" {
				title = rsp.Title
			}
			if err := sendToWechat(title, rsp.Content); err == nil {
				history.PubDate = rsp.Id.(int64)
			}
		}
	}
}

// 推送全国概况数据
func pushChinaOverview(title string) {
	if str, err := GetChinaOverviewData(); err != nil {
		log.Println("从接口获取数据失败", err.Error())
		return
	} else {
		if rsp, ok := parseChinaOverview(str); ok {
			if title == "default" || title == "" {
				title = rsp.Title
			}
			if err := sendToWechat(title, rsp.Content); err == nil {
				history.ChinaUpdateTime = rsp.Id.(int64)
			}
		}
	}
}

// 推送省份详细信息
func pushProvinceOverview(provinceName, title string) {
	if str, err := GetProvinceOverviewData(provinceName); err != nil {
		log.Println("从接口获取数据失败", err.Error())
		return
	} else {
		if rsp, ok := parseChinaProvincesData(str); ok {
			if title == "default" || title == "" {
				title = rsp.Title
			}
			if err := sendToWechat(title, rsp.Content); err == nil {
				history.ProvinceUpdateTime = rsp.Id.(int64)
			}
		}
	}
}

// 推送城市详细信息
func pushCityOverview(cityName, title string) {
	if str, err := GetCityOverviewData(cityName); err != nil {
		log.Println("从接口获取数据失败", err.Error())
		return
	} else {
		if rsp, ok := parseCityOverviewData(str); ok {
			if title == "default" || title == "" {
				title = rsp.Title
			}
			if err := sendToWechat(title, rsp.Content); err == nil {
				history.CityData = rsp.Id.(CityData)
			}
		}
	}
}
