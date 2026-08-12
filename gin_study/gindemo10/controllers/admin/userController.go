package admin

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	c.String(200, "用户列表--")
	//con.success(c)
}
func (con UserController) Add(c *gin.Context) {
	//c.String(200, "用户列表-add---")
	c.HTML(http.StatusOK, "admin/useradd.html", gin.H{})
}

/*func (con UserController) Edit(c *gin.Context) {
	c.String(200, "用户列表-Edit------")
}*/

func (con UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")
	file, err := c.FormFile("face")
	//file.Filenaem获取文件名称        aaa.jpg      ./static/upload/aaa.jpg
	dst := path.Join("./static/upload", file.Filename)
	if err == nil {

		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
		"dst":      dst,
	})
}
func (con UserController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useredit.html", gin.H{})
}
func (con UserController) DoEdit(c *gin.Context) {
	username := c.PostForm("username")

	file1, err1 := c.FormFile("face1")
	dst1 := path.Join("./static/upload", file1.Filename)
	if err1 == nil {
		c.SaveUploadedFile(file1, dst1)
	}
	file2, err2 := c.FormFile("face2")
	dst2 := path.Join("./static/upload", file2.Filename)
	if err2 == nil {
		c.SaveUploadedFile(file2, dst2)
	}
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
		"dst1":     dst1,
		"dst2":     dst2,
	})

}
