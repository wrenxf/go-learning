package routers

import (
	"gindemo06/controllers/defaults"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", defaults.DefaultController{}.Index)
		defaultRouters.GET("/news", defaults.DefaultController{}.News)
	}
}
