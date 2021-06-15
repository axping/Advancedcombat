package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	//需要处理消息
	message := "你好，我的小猫咪"

	//编码消息
	encodingMessage := base64.StdEncoding.EncodeToString([]byte(message))

	//输出编码消息
	fmt.Println(encodingMessage)

	fmt.Println("解码消息===============")
	//解码消息
	data, err := base64.StdEncoding.DecodeString(encodingMessage)
	if err != nil {
		panic("base64 解析异常")
	}

	fmt.Println(string(data))

}
