package main

import (
	"errors"
	"github.com/levigross/grequests"
	"log"
	"time"
)

// server酱返回的数据结构
type ServerResponse struct {
	Errno   int    `json:"errno"`
	Errmsg  string `json:"errmsg"`
	Dataset string `json:"dataset"`
}

// 发送到微信端
func sendToWechat(title, content string) error {
	param := map[string]string{"text": title, "desp": content}
	ro := &grequests.RequestOptions{
		Params:         param,
		RequestTimeout: 10 * time.Second}
	url := "https://sc.ftqq.com/" + sendKey + ".send"

	if response, err := grequests.Get(url, ro); err != nil {
		log.Println("发送请求错误", err.Error())
		return err
	} else {
		var serverResponse ServerResponse
		if err := response.JSON(&serverResponse); err != nil {
			log.Println("解析数据至json失败")
			return err
		}
		// 0 代表发送成功
		if serverResponse.Errno != 0 {
			log.Println("发送失败,", serverResponse.Errmsg)
			return errors.New(serverResponse.Errmsg)
		} else {
			log.Println("发送成功，返回值:", response.String())
			return nil
		}
	}
}
