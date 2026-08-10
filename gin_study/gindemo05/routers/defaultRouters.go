package routers

import "github.com/gin-gonic/gin"

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", func(c *gin.Context) {
			c.String(200, "首页")
		})
		defaultRouters.GET("/news", func(c *gin.Context) {
			c.String(200, "新闻")
		})
	}
}
