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

	go SubChina("default", 300)

	go SubProvince("安徽省", "default", 300)
	go SubProvince("上海市", "default", 300)

	go SubCity("合肥", "default", 300)
	go SubCity("万州区", "default", 300)

	go SubNews("all", "default", 300)
	go SubNews("安徽省", "default", 300)

	go SubRumors("default", 300)
	select {}
}
