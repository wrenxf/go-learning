package main

import "github.com/gin-gonic/gin"

func main() {
	//创建一个默认的路由
	r := gin.Default()
	//配置路由
	r.GET("/", func(c *gin.Context) {
		c.String(200, "值%v", "你好Gin")
	})
	r.GET("/news", func(c *gin.Context) {
		c.String(200, "我是新闻页面")
	})
	r.POST("/add", func(c *gin.Context) {
		c.String(200, "我是post请求返回的数据")
	})
	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "这是一个put请求,主要用于编辑数据")
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "这是一个delete请求,主要用于删除数据")
	})
	//启动一个web服务(默认为8080端口)
	//r.Run()

	//启动一个web服务指定特定端口
	r.Run(":8000")
}
