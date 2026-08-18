package admin

import (
	"fmt"
	"gormdemo01/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	//查询数据库
	/*userList := []models.User{}
	models.DB.Find(&userList)
	c.JSON(200, gin.H{
		"result": userList,
	})*/
	//查询age大于20的用户
	userList := []models.User{}
	models.DB.Where("age>20").Find(&userList)
	c.JSON(200, gin.H{
		"result": userList,
	})
}
func (con UserController) Add(c *gin.Context) {
	user := &models.User{
		Id:       6,
		Username: "Gorm",
		Age:      22,
		Email:    "123@qq.com",
		AddTime:  int(models.GetUnix()),
	}
	fmt.Println(user)
	models.DB.Create(&user)
	c.JSON(200, "增加数据成功")
}

func (con UserController) Edit(c *gin.Context) {
	//1.保存所有字段
	/*//查询id等于6的数据
	user := models.User{Id: 6}
	models.DB.Find(&user)
	fmt.Println(user)
	//修改数据
	user.Username = "哈哈哈"
	user.Email = "321@qq.com"
	user.AddTime = int(models.GetUnix())
	models.DB.Save(&user)
	*/
	//2.更新单列
	//user := models.User{}
	//models.DB.Model(&user).Where("id=?", 6).Update("username", "哈哈哈哈哈哈哈哈哈")
	user := models.User{}
	models.DB.Where("id=?", 6).Find(&user)
	user.Username = "哈哈哈"
	user.Email = "321@qq.com"
	user.AddTime = int(models.GetUnix())
	models.DB.Save(&user)
	c.String(200, "更新用户成功")
}
func (con UserController) Delete(c *gin.Context) {
	user := models.User{Id: 6}
	models.DB.Delete(&user)
	c.String(200, "删除用户成功")
}
