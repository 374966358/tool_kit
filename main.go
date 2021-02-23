package main

import (
	"github.com/gin-gonic/gin"

	"tools_more/Demo"
	"tools_more/Travel"
	"tools_more/Wage"
	"tools_more/Year"
)

func main() {

	r := gin.Default()

	Wage.IncAction(r)
	Demo.IncAction(r)
	Travel.IncAction(r)
	// Ciming.IncAction(r)
	Year.IncAction(r)

	// **代表所有子目录 *代表子目录下的所有模板
	r.Static("/static", "./Statics")

	r.LoadHTMLGlob("Views/*/*")

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run()
}
