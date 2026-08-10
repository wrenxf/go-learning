package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

type UserInfo struct {
	Username string `json:"username"  form:"username"`
	Password string `json:"password"  form:"password"`
}
type artile struct {
	Title   string `json:"title"  xml:"title"`
	Content string `json:"content"  xml:"content"`
}

func main() {
	r := gin.Default()
	//自定义模板函数 注意要把这个函数放到加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})
	//加载模板文件，放到配置路由前面
	r.LoadHTMLGlob("templates/**/*")
	//配置静态web目录 第一个参数表示路由，第二个参数表示映射目录
	r.Static("/static", "./static")
	//get请求传值
	r.GET("/", func(c *gin.Context) {
		username := c.Query("username")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})
	//post演示
	r.GET("user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/user.html", gin.H{})
	})
	//获取表单post过来的数据
	r.POST("doAddUser", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	//获取get post传值绑定到结构体
	r.GET("/getUser", func(c *gin.Context) {
		user := &UserInfo{}
		if err := c.ShouldBind(&user); err == nil {
			fmt.Printf("%#v", user)
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	})
	//获取post xml数据
	r.POST("/xml", func(c *gin.Context) {
		article := &artile{}
		xmlSliceData, _ := c.GetRawData() //获取c.request.Body读取请求数据
		if err := xml.Unmarshal(xmlSliceData, &article); err == nil {
			c.JSON(http.StatusOK, article)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}

	})
	//动态路由传值 list/123
	r.GET("/list/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		c.String(http.StatusOK, "%v", cid)
	})
	//后台
	r.GET("/admine", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
		})
	})
	r.Run()
}
