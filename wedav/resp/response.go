package resp

import (
	"encoding/json"
	"fmt"
)

//用于响应返回结构体
type R struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//用于创建响应结果，结构体
type D map[string]interface{}

var Error_message = "{\"code\":500,\"message\":\"服务响应错误\",\"data\":\"null\"}"

//创建一个返回结果
func NewR(code int, message string, data interface{}) *R {
	return &R{Code: code, Message: message, Data: data}
}

//将结果转换未JSON数据
func (r *R) JSON() string {
	resp, err := json.Marshal(r)
	if err != nil {
		return Error_message
	}
	return string(resp)
}

var (
	ERROR_PASSWORD_CODE int = 50001
	ERRROR_NAME_CODE    int = 50002
	ERROR_NET_CODE      int = 50004
)

//异常code解析说明
var errors map[int]string = map[int]string{
	ERROR_PASSWORD_CODE: "请求异常",
}

func Erron(code int) string {
	if _, ok := errors[code]; ok {
		return errors[code]
	}
	return fmt.Sprintf("Error %v code  not found", code)
}
