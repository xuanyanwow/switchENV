package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)


var OpenChannel = make(chan int)


type Response struct {
	Code  int       `json:"code"`
	Data  map[string]interface{}  `json:"data"`
	Msg   string    `json:"msg"`
}

func Json(code int, data map[string]interface{}, msg string) string {
	res := &Response{
		Code: code,
		Data: data,
		Msg: msg,
	}
	content, err := json.Marshal(res)

	if err != nil {
		panic(err)
	}

	return string(content)
}

func Success(data map[string]interface{}, msg string) string {
	return Json(200, data,msg)
}

func FailParam(data map[string]interface{}, msg string) string {
	return Json(400, data,msg)
}

func Fail(data map[string]interface{}, msg string) string {
	return Json(500, data,msg)
}

func End(c *gin.Context, string string){
	c.String(200, string)
}