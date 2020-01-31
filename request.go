package main

import (
	"errors"
	"fmt"
	"github.com/levigross/grequests"
	"log"
	"time"
)

// http 响应信息结构
type Response struct {
	Results []map[string]interface{} `json:"results"`
	Success bool                     `json:"success"`
}

// 获取最新一条辟谣或防治信息 ok
func GetRumors() (string, error) {
	url := baseUrl + "rumors?num=1"
	fmt.Println(url)
	return CommonRequest(url)
}

// 获取指定省份的最新一条新闻 ok
func GetProvinceNews(provinceName string) (string, error) {
	url := baseUrl + "news?num=1&province=" + provinceName
	return CommonRequest(url)
}

// 获取所有地区的最新一条新闻 ok
func GetAreaNews() (string, error) {
	url := baseUrl + "news?num=1"
	return CommonRequest(url)
}

// 获取指定省份的所有统计数据 ok
func GetProvinceOverviewData(provinceName string) (string, error) {
	url := baseUrl + "area?latest=1&province=" + provinceName
	return CommonRequest(url)
}

// 获取指定省份指定城市的统计数据(该功能可能不够完善，一些自治区的州) ok
func GetCityOverviewData(cityName string) (string, error) {
	url := baseUrl + "city?cityName=" + cityName
	return CommonRequest(url)
}

// 获取全国所有省份的最新统计数据 todo 未完成
func GetChinaProvincesData() (string, error) {
	url := baseUrl + "country?countryName=" + "中国"
	return CommonRequest(url)
}

// 获取全国最新概览数据 ok
func GetChinaOverviewData() (string, error) {
	url := baseUrl + "overall?latest=1"
	return CommonRequest(url)
}

// 通用请求
func CommonRequest(url string) (string, error) {
	ro := &grequests.RequestOptions{
		RequestTimeout: 10 * time.Second}

	if response, err := grequests.Get(url, ro); err != nil {
		log.Println("发送请求错误", err.Error())
		return "", err
	} else {
		var rsp Response
		rspStr := response.String()
		if err := response.JSON(&rsp); err != nil {
			return "解析json错误", err
		} else {
			if rsp.Success == false {

				return "不正确的返回结果", errors.New("不正确的返回结果")
			} else {
				log.Println("响应字符串为:", rspStr)
				return rspStr, nil
			}

		}

	}
}
