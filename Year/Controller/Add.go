package Controller

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	Node "tools_more/Year/Model"
)

func Add(context *gin.Context) {
	var addData Node.AddStruct

	err := context.ShouldBind(&addData)

	if err != nil {
		context.JSON(http.StatusOK, gin.H{"code": 500, "msg": "请正确提交信息"})
	} else {
		addData.OpenId, _ = url.QueryUnescape(addData.OpenId)

		err := addData.Insert()

		if err != nil {
			context.JSON(http.StatusOK, gin.H{"code": 500, "msg": "添加失败请稍后再试"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"code": 200, "msg": "添加成功"})
	}
}
