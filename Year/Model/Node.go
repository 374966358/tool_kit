package Node

import (
	"tools_more/Core"
)

type AddStruct struct {
	OpenId     string `gorm:"" form:"open_id"`
	UniId      string `gorm:""`
	Alias      string `gorm:"" form:"alias"`
	Title      string `gorm:"" form:"title"`
	Content    string `gorm:"" form:"content"`
	Background string `gorm:"" form:"background_image"`
	Img        string `gorm:"" form:"blessing_image"`
	FuImg      string `gorm:"" form:"fu_image"`
	Addtime    string `gorm:""`
}

func (AddStruct) TableName() string {
	return "node"
}

func (a AddStruct) Insert() error {
	Core.GetDB("zhufu")

	return Core.DB.Create(a).Error
}
