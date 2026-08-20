package admin

import (
	"gormdemo02/models"

	"github.com/gin-gonic/gin"
)

// 测试的结构体
type NavJson struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func (NavJson) TableName() string {
	return "nav"
}

type NavController struct {
	BaseController
}

func (con NavController) Index(c *gin.Context) {

	// 查询全部数据
	// navList := []models.Nav{}
	// models.DB.Find(&navList)
	// c.JSON(200, gin.H{
	// 	"result": navList,
	// })

	// 查询一条数据

	// navResult := models.Nav{Id: 21}
	// models.DB.Find(&navResult)
	// c.JSON(200, gin.H{
	// 	"result": navResult,
	// })

	/*
		Where条件
		=
		<
		>
		<=
		>=
		!=
		IS NOT NULL
		IS NULL
		BETWEEN AND
		NOT BETWEEN AND
		IN
		OR
		AND
		NOT
		LIKE
	*/

	//查询Id大于3的数据
	// navList := []models.Nav{}
	// models.DB.Where("id>3").Find(&navList)
	// c.JSON(200, gin.H{
	// 	"result": navList,
	// })

	//查询 id大于3 并且 id小于9的数据

	// var a = 3
	// var b = 9
	// navList := []models.Nav{}
	// models.DB.Where("id>? AND id<?", a, b).Find(&navList)
	// c.JSON(200, gin.H{
	// 	"result": navList,
	// })

	//使用in查询 id 在 3, 5, 6中的数据

	// navList := []models.Nav{}
	// models.DB.Where("id in (?)", []int{3, 5, 6}).Find(&navList)
	// c.JSON(200, gin.H{
	// 	"result": navList,
	// })

	//使用like查询标题里面包含 会 的内容
	// navList := []models.Nav{}
	// models.DB.Where("title like ?", "%会%").Find(&navList)
	// c.JSON(200, gin.H{
	// 	"result": navList,
	// })

	//查询 id在3和9之间的数据 使用 between and

	// navList := []models.Nav{}
	// models.DB.Where("id between ? and ?", 3, 9).Find(&navList)
	// c.JSON(200, gin.H{
	// 	"result": navList,
	// })

	//Or 查询id=2 或者 id=3的数据
	// navList := []models.Nav{}
	// models.DB.Where("id =? OR id =?", 2, 3).Find(&navList)
	// c.JSON(200, gin.H{
	// 	"result": navList,
	// })

	// navList := []models.Nav{}
	// models.DB.Where("id =?", 2).Or("id = ?", 3).Or("id = ?", 4).Find(&navList)
	// c.JSON(200, gin.H{
	// 	"result": navList,
	// })

	//使用Select指定返回的字段

	// navList := []NavJson{}
	// models.DB.Select("id,title").Find(&navList)
	// c.JSON(200, gin.H{
	// 	"result": navList,
	// })

	//Order排序 、Limit 、Offset

	// navList := []models.Nav{}
	// models.DB.Order("id desc").Order("sort asc").Find(&navList)
	// c.JSON(200, gin.H{
	// 	"result": navList,
	// })

	// navList := []models.Nav{}
	// models.DB.Order("id desc").Order("sort asc").Limit(2).Find(&navList)
	// c.JSON(200, gin.H{
	// 	"result": navList,
	// })

	//分页
	// navList := []models.Nav{}
	// models.DB.Order("id desc").Offset(2).Limit(2).Find(&navList)
	// c.JSON(200, gin.H{
	// 	"result": navList,
	// })

	//Count 统计总数
	// navList := []models.Nav{}
	// var num int64
	// models.DB.Find(&navList).Count(&num)
	// c.JSON(200, gin.H{
	// 	"result": num,
	// })

	//使用原生 sql 删除 user 表中的一条数据

	// models.DB.Exec("delete from user where id=?", 5)

	//使用原生 sql 修改 user 表中的一条数据

	// models.DB.Exec("update user set username='itying gin grom' where id=?", 6)

	//使用原生 sql查询 uid=2 的数据

	// userList := []models.User{}
	// models.DB.Raw("select * from user where id >2").Scan(&userList)

	// c.JSON(200, gin.H{
	// 	"result": userList,
	// })

	//使用原生 查询 User 表中所有的数据

	// userList := []models.User{}
	// models.DB.Raw("select * from user").Scan(&userList)

	// c.JSON(200, gin.H{
	// 	"result": userList,
	// })

	//统计 user 表的数量

	var num int
	models.DB.Raw("select count(1) from user").Scan(&num)
	c.JSON(200, gin.H{
		"result": num,
	})

	// c.String(200, "Nav Index")
}
