package Travel

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	"tools_more/Common"
	"tools_more/Travel/Model"
)

type travel struct {
	WeChat      *Common.WeChat
	Title       string
	Keywords    string
	Description string
}

func IncAction(r *gin.Engine) {
	r.GET("/travel", func(context *gin.Context) {
		data := travel{
			WeChat:      Common.GetWeChat("http://6xcloud.com/travel"),
			Title:       "疫情出行各地政策查询",
			Keywords:    "返乡政策,返乡,疫情,出行,北京,上海,深圳,广州",
			Description: "收集各地返乡政策，提供回城帮助，疫情期间出行，先查询，再行动，即刻查询目的地的政策",
		}

		context.HTML(http.StatusOK, "travel_index.html", data)
	})

	r.GET("/travel/index", func(context *gin.Context) {
		var formP string = context.DefaultQuery("form_p", "北京")
		var formC string = context.DefaultQuery("form_c", "北京")
		var toP string = context.DefaultQuery("to_p", "")
		var toC string = context.DefaultQuery("to_c", "")

		buf, err := Model.DoMchrApi("api/QuarkGo/getHomeData?fromp=" + url.QueryEscape(formP) + "&fromc=" + url.QueryEscape(formC) + "&top=" + url.QueryEscape(toP) + "&toc=" + url.QueryEscape(toC))

		if err != nil {
			panic(err.Error())
			return
		}

		var dps Model.MyJsonName

		_ = json.Unmarshal(buf, &dps)

		context.JSON(http.StatusOK, dps)
	})

	r.GET("/travel/tel", func(context *gin.Context) {
		var province string = context.DefaultQuery("province", "北京")
		var city string = context.DefaultQuery("city", "北京")

		buf, err := Model.DoMchrApi("api/QuarkGo/tellist?province=" + url.QueryEscape(province) + "&city=" + url.QueryEscape(city))

		if err != nil {
			panic(err.Error())
			return
		}

		var dps Model.MyTel

		_ = json.Unmarshal(buf, &dps)

		context.JSON(http.StatusOK, dps)
	})
}
