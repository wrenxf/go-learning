package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}
func main() {
	r := gin.Default()
	//自定义模板函数 注意要把这个函数放到加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})
	//加载模板文件，放到配置路由前面
	r.LoadHTMLGlob("templates/**/*")
	//配置静态web目录 第一个参数表示路由，第二个参数表示映射目录
	r.Static("/static", "./static")
	//前台
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "首页")
	})
	//后台
	r.GET("/admine", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
		})
	})
	r.Run()
}
