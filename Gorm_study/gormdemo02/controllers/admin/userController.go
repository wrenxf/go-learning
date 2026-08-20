package admin

import (
	"fmt"
	"gormdemo02/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {

	//查询数据库
	// userList := []models.User{}
	// models.DB.Find(&userList)
	// c.JSON(200, gin.H{
	// 	"result": userList,
	// })

	//查询age大于20的用户
	userList := []models.User{}
	models.DB.Where("age>20").Find(&userList)
	c.JSON(200, gin.H{
		"result": userList,
	})

}
func (con UserController) Add(c *gin.Context) {
	//增加数据
	user := models.User{
		Username: "itying GORM",
		Age:      22,
		Email:    "222@qq.con",
		AddTime:  int(models.GetUnix()),
	}

	models.DB.Create(&user)
	fmt.Println(user)
	c.String(200, "增加数据成功")
}
func (con UserController) Edit(c *gin.Context) {
	//保存所有字段

	// 查询id等于6的数据
	// user := models.User{Id: 6}
	// models.DB.Find(&user)
	// //更新数据
	// user.Username = "哈哈"
	// user.Email = "itying@qqq.com"
	// user.AddTime = int(models.GetUnix())
	// models.DB.Save(&user)

	//更新单个列
	// user := models.User{}
	// models.DB.Model(&user).Where("id = ?", 6).Update("username", "哈哈哈哈哈哈")

	//另一种更新数据的方法
	user := models.User{}
	models.DB.Where("id = ?", 6).Find(&user)
	user.Username = "哈哈"
	user.Email = "aaa@qqq.com"
	user.AddTime = int(models.GetUnix())
	models.DB.Save(&user)

	c.String(200, "修改用户成功")
}
func (con UserController) Delete(c *gin.Context) {

	//删除id=6的数据
	// user := models.User{Id: 6}
	// models.DB.Delete(&user)

	//删除username=gorm的数据
	user := models.User{}
	models.DB.Where("username = ?", "gorm").Delete(&user)
	c.String(200, "删除用户")
}
