package Wage

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tools_more/Wage/Conf"
	"tools_more/Wage/Cortroller"
)

func IncAction(r *gin.Engine) {
	//r.LoadHTMLGlob("Wage/View/*")
	// r.Static("/static", "./Wage/Static")
	r.GET("/", func(c *gin.Context) {
		var formdata Conf.Formdata
		formparam := Conf.Beijing()
		formdata.City = "bj"
		//formdata.Money = 0
		formdata.Param = formparam
		formdata.List = ""
		c.HTML(http.StatusOK, "wage_index.html", Cortroller.Math(formdata))
	})
	r.GET("/MP_verify_an5Thf5lXORqKP6w.txt", func(c *gin.Context) {
		c.File("MP_verify_an5Thf5lXORqKP6w.txt")
		c.HTML(http.StatusOK, "wage_sitemap.html", "")
	})
	r.GET("/sitemap", func(c *gin.Context) {
		type Sitemap struct {
			List []struct {
				Title string
				Link  string
			}
		}
		c.HTML(http.StatusOK, "wage_sitemap.html", Sitemap{})
	})
	r.GET("/city/:city/:money", func(c *gin.Context) {
		city := c.Param("city")
		money := c.Param("money")
		var formdata Conf.Formdata
		var formparam Conf.FormParam
		if city == "beijing" {
			formparam = Conf.Beijing()
		} else {
			formparam = Conf.Other()
		}
		c.ShouldBind(&formparam)
		formdata.City = city
		formdata.Money, _ = strconv.ParseFloat(money, 64)
		formdata.Param = formparam
		c.HTML(http.StatusOK, "wage_index.html", Cortroller.Math(formdata))
	})
	//r.GET("/wageApi", func(c *gin.Context) {
	//	var formdata Conf.Formdata
	//	formparam := Conf.Beijing()
	//	formdata.Param = formparam
	//	c.JSON(200, Cortroller.Math(formdata))
	//})
}
