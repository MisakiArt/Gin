package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)
type AuthSt struct {
	Authorization string
	JCustomerUUID string
}

type CheckAccess struct {
	Items []string `json:"items"`
}

type CheckAccessResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data []string `json:"data"`
}



func SuccessResponse(code int,message string, m interface{},c *gin.Context)  {
	c.JSON(200, gin.H{
		"code":code,
		"message":message,
		"data":m,
	})
}



/**
str 变成字符串的消息体
url 请求地址

 */
func (l *AuthSt) HttpPost(str string,url string,responseForm interface{}){
	var jsonstr = []byte(str)
	buffer:=bytes.NewBuffer(jsonstr)
	request,_:=http.NewRequest("POST",url,buffer)
	request.Header.Set("Content-Type","application/json;charset=UTF-8")
	request.Header.Set("Authorization",l.Authorization)
	request.Header.Set("J-CustomerUUID",l.JCustomerUUID)
	clilent :=http.Client{}
	resp,_ :=clilent.Do(request.WithContext(context.TODO()))
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll%v", err)
	}
	_=json.Unmarshal(respBytes,responseForm)
}