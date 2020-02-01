package main

const baseUrl = "https://lab.ahusmart.com/nCoV/api/"
const sendKey = "SCUxxx"

// 初始化历史数据记录器
var history = History{
	RumorsId:           0,
	PubDate:            0,
	ProvinceUpdateTime: 0,
	ChinaUpdateTime:    0,
}

func main() {

	go SubChina("default", 3600)
	go SubProvince("安徽省", "default", 1800)
	go SubCity("合肥", "default", 100)
	go SubNews("all", "default", 100)
	go SubRumors("default", 600)
	select {}
}
