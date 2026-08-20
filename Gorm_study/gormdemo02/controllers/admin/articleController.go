package admin

import (
	"gormdemo02/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	BaseController
}

func (con ArticleController) Index(c *gin.Context) {
	articleList := []models.Article{}

	models.DB.Find(&articleList)

	c.JSON(200, gin.H{
		"result": articleList,
	})

}
func (con ArticleController) Add(c *gin.Context) {
	c.String(http.StatusOK, "-add--文章-")
}
func (con ArticleController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "-Edit---文章---")
}
