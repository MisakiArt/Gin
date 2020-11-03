package main

import (
	logic "GinTest/Logic"
	"GinTest/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type testForm struct {
	Message string `form:"message"  binding:"required"  json:"message"`
	Nick string `form:"nick"  json:"nick" binding:"required" `
}


func test(c *gin.Context) {
	var form testForm
	if c.ShouldBind(&form) !=nil {
		c.JSON(-1, gin.H{"message": "params error"})
	}
	fmt.Println(form)
	var tag []models.Tags
	models.GetAll(&tag)
	for _,v :=range tag {
		fmt.Println(v)
	}


	access:= []string{"jsf-UserAccountSettingList","jsf-FollowersDownloadList","jsf-UserUserProfileList"}
	accessStruct :=logic.CheckAccess{Items: access}
	accessString,_ := json.Marshal(accessStruct)

	logicS := logic.LogicStruct{
		Authorization: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1OTU4MTUwNDgsInVzZXJpZCI6IjU0OSIsInRpbWVvdXQiOjUwNDAwLCJyZXF1ZXN0X3V1aWQiOiIxNTk1ODE1MDQ4LjE4ODgtMTg3NSJ9.idR4xjGDWnfSCzVU5kAeVpuRNwGHdR0LqynD5ELmxN4",
		JCustomerUUID: "428cff2385ebe526"}
	responseForm := logic.CheckAccessResponse{}
	 logicS.HttpPost(string(accessString),"https://devapi.jingsocial.com/api/user/users/checkuseraccess",&responseForm)
	logic.SuccessResponse(200,"",responseForm,c)
}


func main() {

	r := gin.Default()
	r.POST("/ping", test)
	_=r.Run(":9508") // 监听并在 0.0.0.0:8080 上启动服务
}