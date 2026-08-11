package admin

import "github.com/gin-gonic/gin"

type ArticleController struct{}

func (con ArticleController) Index(c *gin.Context) {
	c.String(200, "文章列表")
}
func (con ArticleController) Add(c *gin.Context) {
	c.String(200, "add-文章")
}
func (con ArticleController) Edit(c *gin.Context) {
	c.String(200, "edit-文章")
}
