package main

// 由于支持markdown语法，因此换两行才是换行
var chinaOverview_template = `更新时间：%v

确诊病例: %v

疑似病例：%v

治愈病例：%v

死亡病例：%v

![地图](%v)
`

// 辟谣信息展示模板
var rumors_template = `%v

%v

`

//省份信息展示模板
var provinceNews_template = `%v
 
更新时间: %v

来源:[%v](%v)
`

//省份概览数据
var provinceOverviewData_template = `
## %v

确诊数: %v

疑似数: %v

治愈数: %v

死亡数: %v

`

// 指定城市概览数据
var cityOverviewData_template = `
## %v

确诊数: %v

疑似数: %v

治愈数: %v

死亡数: %v

`

// 省份详情表头
var Province_Header_template = `| 省份 &emsp;| 确诊数 &emsp; | 疑似数 &emsp;|治愈数 &emsp; |死亡数 &emsp; |
|  :----: |  :----: |:----: | :----:| :----: |
`

// 多个城市详情表头
var City_Header_template = `| 城市 &emsp;| 确诊数 &emsp; | 疑似数 &emsp;|治愈数 &emsp; |死亡数 &emsp; |
|  :----: |  :----: |:----: | :----:| :----: |
`

//单条城市信息
var singleProvinceOrCity = `|   %v &emsp;|    %v &emsp;    |    %v &emsp;   |   %v &emsp;   |   %v &emsp;   |
`
