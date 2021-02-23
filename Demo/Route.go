package Demo

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"tools_more/Demo/Model"
)

func IncAction(r *gin.Engine) {

	r.GET("/other/index", func(c *gin.Context) {
		buf, err := Model.DoMchrApi("api/QuarkGo/getHomeData?fromp=" + url.QueryEscape("浙江") + "&fromc=" + url.QueryEscape("杭州") + "&top=" + url.QueryEscape("辽宁") + "&toc=" + url.QueryEscape("丹东"))
		if err != nil {
			panic(err)
		}
		var dps Model.MyJsonName
		json.Unmarshal(buf, &dps)
		//c.JSON(200, dps)
		c.HTML(http.StatusOK, "demo_index.html", dps)
	})
}

// 各地区数量接口：https://m.sm.cn/api/rest?format=json&method=Huoshenshan.local&province=河北&city=河北
