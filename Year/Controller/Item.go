package Controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"tools_more/Core"
)

type itemData struct {
	ID         string `gorm:""`
	Title      string `gorm:"title"`
	Content    string `gorm:"content"`
	Background string `gorm:"background"`
	Img        string `gorm:"img"`
	FuImg      string `gorm:"fu_img"`
}

func (itemData) TableName() string {
	return "node"
}

func Item(context *gin.Context) {
	var param = context.Params

	var item itemData

	var id = param.ByName("ID")
	//var openId = param.ByName("OpenID")

	Core.GetDB("zhufu")

	Core.DB.Select("id,title, content, background, img, fu_img").Where("id = ? AND  state = 1", id).First(&item)

	context.HTML(http.StatusOK, "year_item.html", gin.H{
		"Data": item,
	})
}
