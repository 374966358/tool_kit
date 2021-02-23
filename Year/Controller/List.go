package Controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"tools_more/Core"
)

type listStruct struct {
	ID         int    `gorm:""`
	Alias      string `gorm:""`
	Title      string `gorm:""`
	Background string `gorm:""`
	Content    string `gorm:""`
	Img        string `gorm:""`
	FuImg      string `gorm:""`
	State      int32  `gorm:""`
}

func (listStruct) TableName() string {
	return "node"
}

func List(context *gin.Context) {
	var openId = context.Param("OpenID")

	var listData []listStruct

	Core.GetDB("zhufu")

	Core.DB.Select("id, alias, title,background, content, img, fu_img, state").Where("open_id = ?", openId).Find(&listData)

	context.HTML(http.StatusOK, "year_list.html", gin.H{
		"lists": listData,
	})
}
