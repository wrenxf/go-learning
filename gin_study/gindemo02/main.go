package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	//配置模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "值：%v", "首页")
	})
	r.GET("/json1", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "你好gin",
		})
	})
	r.GET("/json2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "你好gin",
		})
	})
	r.GET("/json3", func(c *gin.Context) {
		a := Article{
			Title:   "标题",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSON(200, a)
	})
	r.GET("/jsonp", func(c *gin.Context) {
		a := Article{
			Title:   "标题",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSONP(200, a)
	})
	r.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{
			"success": true,
			"msg":     "你好gin我是xml",
		})
	})
	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html", gin.H{
			"title": "我是后台数据",
		})
	})
	r.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "我是后台数据",
			"price": "20",
		})
	})
	r.Run()
}
