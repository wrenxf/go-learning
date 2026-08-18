# 使用代理

```
https://goproxy.cn
```

使用 `https://goproxy.cn` 的核心原因非常直接：**它是目前在中国大陆访问速度最快、最稳定的 Go 模块代理。**

简单来说，如果你不使用这个代理，当你在 Go 代码中引入第三方库（比如 Gin 框架）时，`go get` 命令会尝试直接去 GitHub 等境外网站下载代码。由于网络环境的特殊性，这个过程往往会非常慢，甚至直接连接超时失败。

# 初始化 Go 模块

```go
go mod init gin_study
```

在开始学习 Gin 框架前执行 `go mod init gin_study`，是**初始化 Go 模块**的标准操作，目的是为项目建立依赖管理体系，让后续引入 Gin 等第三方库时能自动下载、版本可控、避免冲突。这一步是现代 Go 开发的“地基”，没有它，你无法顺利使用任何外部框架。

# go mod 常用指令

| 指令                   | 作用说明                                               | 常用示例                                 |
| ---------------------- | ------------------------------------------------------ | ---------------------------------------- |
| `go mod init <名称>`   | 初始化模块，创建 `go.mod` 文件（项目第一步必做）。     | `go mod init gindemo01`                  |
| `go mod tidy`          | 自动添加缺失依赖、移除无用依赖，保持文件整洁。         | `go mod tidy`                            |
| `go get <包>@<版本>`   | 下载指定包并更新到 `go.mod`，可指定版本号。            | `go get github.com/gin-gonic/gin@v1.9.1` |
| `go get -u`            | 将项目中所有依赖升级到最新的兼容版本。                 | `go get -u`                              |
| `go mod download`      | 仅下载依赖到本地缓存，不修改代码或文件。               | `go mod download`                        |
| `go list -m all`       | 列出当前模块及所有直接/间接依赖的版本信息。            | `go list -m all`                         |
| `go mod graph`         | 打印完整的依赖关系图，用于排查冲突。                   | `go mod graph`                           |
| `go mod vendor`        | 将所有依赖复制到项目内的 `vendor` 目录（离线构建用）。 | `go mod vendor`                          |
| `go mod edit -replace` | 在 `go.mod` 中添加替换规则（常用于本地调试）。         | `go mod edit -replace=old=new`           |

# 一、Gin 介绍

Gin 是一个 Go (Golang) 编写的轻量级 http web 框架，运行速度非常快，如果你是性能和高效的追求者，我们推荐你使用 Gin 框架。

Gin 最擅长的就是 Api 接口的高并发，如果项目的规模不大，业务相对简单，这个时候我们也推荐您使用 Gin。

当某个接口的性能遭到较大挑战的时候，这个还是可以考虑使用 Gin 重写接口。

Gin 也是一个流行的 golang Web 框架，Github Strat 量已经超过了 50k。

Gin 的官网：https://gin-gonic.com/zh-cn/

Gin Github 地址：https://github.com/gin-gonic/gin

# 二、Gin 环境搭建

要安装 Gin 软件包，需要先安装 Go 并设置 Go 工作区。

**现代 Go 工具链（1.16+）默认要求所有项目必须在模块模式下运行。**

1. **进入你的项目目录**
   先创建一个项目文件夹并进入：

   ```
   mkdir my-gin-project
   cd my-gin-project
   ```

2. **初始化 Go 模块**
   在项目根目录执行：

   ```
   go mod init my-gin-project
   ```

   这会生成一个 `go.mod` 文件，它是 Go 项目的依赖管理核心。

3. **安装 Gin 框架**
   现在可以安全地安装 Gin：

   ```
   go get -u github.com/gin-gonic/gin
   ```

## 1.下载并安装 gin：

```
$ go get -u github.com/gin-gonic/gin
```

## 2.将 gin 引入到代码中：

```
import "github.com/gin-gonic/gin" 
```

## 3.（可选）如果使用诸如 http.StatusOK 之类的常量，则需要引入 net/http 包：

```
import "net/http" 
```

http.StatusOK等价于200的状态码

## 4.新建 Main.go 配置路由

```go
package main
import (
    "github.com/gin-gonic/gin"
)
func main() {
    // 创建一个默认的路由引擎
    r := gin.Default()
    // 配置路由
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{ // c.JSON：返回 JSON 格式的数据
            "message": "Hello world!", })
    })
    // 启动 HTTP 服务，默认在 0.0.0.0:8080 启动服务
    r.Run()
}
```

## 5.运行你的项目

```
$ go run main.go
```

## 6.要改变默认启动的端口

```
r.Run(":9000")
```

如果 go get 失败请参考：
http://bbs.itying.com/topic/5ed08edee7c0790f8475e276

# 三、golang 程序的热加载

所谓热加载就是当我们对代码进行修改时，程序能够自动重新加载并执行，这在我们开发中是非常便利的，可以快速进行代码测试，省去了每次手动重新编译

beego 中我们可以使用官方给我们提供的 bee 工具来热加载项目，但是 gin 中并没有官方提供的热加载工具，这个时候我们要实现热加载就可以借助第三方的工具。

工具 1（推荐）：https://github.com/gravityblast/fresh

```
go get github.com/pilu/fresh
D:\gin_demo>fresh
```

工具 2：https://github.com/codegangsta/gin

```
go get -u github.com/codegangsta/gin
D:\gin_demo>gin run main.go
```

## 遇到的问题

![image1](assets/image1.png)

你遇到的 `fresh: The term 'fresh' is not recognized...` 错误，是因为 **Fresh 工具虽然已下载依赖包，但其可执行文件（fresh.exe）未被正确安装到系统 PATH 环境变量中**，导致终端无法识别该命令。这是 Go 1.16+ 版本后使用 `go get` 安装命令行工具的常见问题。

### **解决方案（推荐按顺序操作）**

#### **第一步：使用** `go install` **安装 Fresh（关键步骤）**

在终端中执行以下命令，将 Fresh 编译并安装到 Go 的 bin 目录：

```
go install github.com/pilu/fresh@latest
```

> ✅ 这一步会将 `fresh.exe` 安装到 `$GOPATH/bin` 或 `$GOBIN` 目录下（通常是 `C:\Users\你的用户名\go\bin`）。

#### **第二步：确认安装路径并添加到系统 PATH**

1. 打开 PowerShell 或 CMD，执行：

   ```
   go env GOPATH
   ```

   你会看到类似 `C:\Users\你的用户名\go` 的路径。

2. 将 `C:\Users\你的用户名\go\bin` 添加到系统环境变量 `PATH` 中：

   - 右键“此电脑” → “属性” → “高级系统设置” → “环境变量”。
   - 在“用户变量”或“系统变量”中找到 `Path`，点击“编辑”。
   - 新建一条，填入 `C:\Users\你的用户名\go\bin`（请替换为你的实际路径）。
   - 点击“确定”保存。

3. **重启终端**（PowerShell/CMD），使环境变量生效。

#### **第三步：验证安装成功**

在新终端中输入：

```
fresh --version
```

如果输出版本号（如 `v0.0.0-20240621171608-8d1fef547a99`），说明安装成功。

------

### **为什么** `go get` **不行？**

从 Go 1.17 开始，`go get` 不再自动将可执行文件安装到 `GOPATH/bin`，它只负责下载依赖包。要安装命令行工具，必须使用 `go install`。

------

### **额外建议：配置 GO111MODULE**

如果你的项目是模块模式（有 `go.mod`），建议显式设置：

```
go env -w GO111MODULE=on
```

这可以避免一些兼容性问题。

# 四、Gin 框架中的路由

## 4.1、路由概述

路由（Routing）是由一个 URI（或者叫路径）和一个特定的 HTTP 方法（GET、POST 等）组成的，涉及到应用如何响应客户端对某个网站节点的访问。

RESTful API 是目前比较成熟的一套互联网应用程序的 API 设计理论，所以我们设计我们的路由的时候建议参考 RESTful API 指南。

在 RESTful 架构中，每个网址代表一种资源，不同的请求方式表示执行不同的操作：

| GET（SELECT）        | 从服务器取出资源（一项或多项）                     |
| -------------------- | :------------------------------------------------- |
| **POST（CREATE）**   | **在服务器新建一个资源**                           |
| **PUT（UPDATE）**    | **在服务器更新资源（客户端提供改变后的完整资源）** |
| **DELETE（DELETE）** | **从服务器删除资源**                               |

### **📌 一句话总结**

- **GET** = 查（只读，不改变数据）
- **POST** = 增（创建新数据）
- **PUT** = 改（替换/更新现有数据）
- **DELETE** = 删（移除数据）

## 4.2、简单的路由配置

**简单的路由配置**(可以通过 postman 测试)
当用 GET 请求访问一个网址的时候，做什么事情：

```go
r.GET("网址", func(c *gin.Context) {
    c.String(200, "Get")
})
```

当用 POST 访问一个网址的时候，做什么事情：

```go
r.POST("网址", func(c *gin.Context) {
    c.String(200, "POST")
})
```

当用 PUT 访问一个网址的时候，执行的操作：

```go
r.PUT("网址", func(c *gin.Context) {
    c.String(200, "PUT")
})
```

当用 DELETE 访问一个网址的时候，执行的操作：

```go
r.DELETE("网址", func(c *gin.Context) {
    c.String(200, "DELETE")
})
```

**路由里面获取 Get 传值**
域名/news?aid=20

```go
r.GET("/news", func(c *gin.Context) {
    aid := c.Query("aid")
    c.String(200, "aid=%s", aid)
})
```

**动态路由**
域名/user/20

```go
r.GET("/user/:uid", func(c *gin.Context) {
    uid := c.Param("uid")
    c.String(200, "userID=%s", uid)
})
```

## 4.3、 c.String() c.JSON() c.JSONP() c.XML() c.HTML()

| 方法         | 返回格式 | Content-Type             | 主要用途                              |
| ------------ | -------- | ------------------------ | ------------------------------------- |
| `c.String()` | 纯文本   | `text/plain`             | 简单消息、调试输出                    |
| `c.JSON()`   | JSON     | `application/json`       | API 数据返回，前后端分离最常用        |
| `c.JSONP()`  | JSONP    | `application/javascript` | 解决跨域请求（通过回调函数包裹 JSON） |
| `c.XML()`    | XML      | `application/xml`        | 旧系统对接、RSS 订阅、SOAP 服务       |
| `c.HTML()`   | HTML     | `text/html`              | 服务端渲染页面（需提前加载模板）      |

**返回一个字符串**

```go
r.GET("/news", func(c *gin.Context) {
    aid := c.Query("aid")
    c.String(200, "aid=%s", aid)
})
```

**返回一个 JSON 数据**

```go
func main() {
    r := gin.Default()
    // gin.H 是 map[string]interface{}的缩写
    r.GET("/someJSON", func(c *gin.Context) {
        // 方式一：自己拼接 JSON
        c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
    })
    r.GET("/moreJSON", func(c *gin.Context) {
        // 方法二：使用结构体
        var msg struct {
            Name string `json:"user"` Message string
            Age int
        }
        msg.Name = "IT 营学院" msg.Message = "Hello world!" msg.Age = 18
        c.JSON(http.StatusOK, msg)
    })
    r.Run(":8080")
}
```

**JSOPN**

```go
func main() {
    r := gin.Default()
    r.GET("/JSONP", func(c *gin.Context) {
        data := map[string]interface{}{ "foo": "bar", }
        // /JSONP?callback=x
        // 将输出：x({\"foo\":\"bar\"})
        c.JSONP(http.StatusOK, data)
    })

    // 监听并在 0.0.0.0:8080 上启动服务
    r.Run(":8080")
}
```

**返回 XML 数据**

```go
func main() {
    r := gin.Default()
    // gin.H 是 map[string]interface{}的缩写
    r.GET("/someXML", func(c *gin.Context) {
        // 方式一：自己拼接 JSON
        c.XML(http.StatusOK, gin.H{"message": "Hello world!"})
    })
    r.GET("/moreXML", func(c *gin.Context) {
        // 方法二：使用结构体
        type MessageRecord struct {
            Name string
            Message string
            Age int
        }
        var msg MessageRecord
        msg.Name = "IT 营学院" msg.Message = "Hello world!" msg.Age = 18
        c.XML(http.StatusOK, msg)
    })
    r.Run(":8080")
}
```

**渲染模板**

```go
router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "default/index.html", map[string]interface{}{ 
        "title": "前台首页"
    })
})
```

# 五、Gin HTML 模板渲染

## 5.1、全部模板放在一个目录里面的配置方法

### 1、我们首先在项目根目录新建 templates 文件夹，然后在文件夹中新建 index.html

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <h1>这是一个 html 模板</h1>
        <h3>{{.title}}</h3>
    </body>
</html>
```

### 2、Gin 框架中使用 c.HTML 可以渲染模板，渲染模板前需要使用 LoadHTMLGlob()或者LoadHTMLFiles()方法加载模板。

```go
router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "default/index.html", map[string]interface{}{ "title": "前台首页"
                                                                      })
})
router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{ "title": "Main website", })
})
```

```go
package main
import ( 
    "net/http"
    "github.com/gin-gonic/gin"
)
func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")
    //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{ "title": "Main website", })
    })
    router.Run(":8080")
}
```

| 特性     | `r.LoadHTMLGlob`           | `r.LoadHTMLFiles`        |
| -------- | -------------------------- | ------------------------ |
| 加载方式 | 批量、按模式匹配           | 手动、逐个指定           |
| 灵活性   | 高，适合目录化管理         | 低，但非常精确           |
| 适用场景 | 模板文件较多，且按目录组织 | 模板文件较少，或位置分散 |

```go
// 加载 templates 目录下的所有文件
r.LoadHTMLGlob("templates/*")

// 加载 templates 目录下所有 .html 后缀的文件
r.LoadHTMLGlob("templates/*.html")

// 加载 templates 目录及其所有子目录下的 .html 文件
r.LoadHTMLGlob("templates/**/*.html")
```

```go
// 加载两个指定的文件
r.LoadHTMLFiles("templates/index.html", "templates/user/profile.html")

// 也可以只加载一个文件
r.LoadHTMLFiles("templates/home.html")
```



## 5.2、模板放在不同目录里面的配置方法

Gin 框架中如果不同目录下面有同名模板的话我们需要使用下面方法加载模板
**注意：定义模板的时候需要通过 define 定义名称**
templates/admin/index.html
<!-- 相当于给模板定义一个名字 define end 成对出现-->

```html
{{ define "admin/index.html" }}
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <h1>后台模板</h1>
        <h3>{{.title}}</h3>
    </body>
</html>
{{ end }}
```

templates/default/index.html
<!-- 相当于给模板定义一个名字 define end 成对出现-->

```html
{{ define "default/index.html" }}
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <h1>前台模板</h1>
        <h3>{{.title}}</h3>
    </body>
</html>
{{end}}
```

**业务逻辑**

```go
package main
import ( 
    "net/http"
    "github.com/gin-gonic/gin"
)
func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/**/*")
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "default/index.html", gin.H{ "title": "前台首页", })
    })
    router.GET("/admin", func(c *gin.Context) {
        c.HTML(http.StatusOK, "admin/index.html", gin.H{ "title": "后台首页", })
    })
    router.Run(":8080")
}
```

**注意：如果模板在多级目录里面的话需要这样配置 r.LoadHTMLGlob("templates///*") /** 表示目录**

## 5.3、gin 模板基本语法

### 1、{{.}} 输出数据

模板语法都包含在{{和}}中间，其中{{.}}中的点表示当前对象。
当我们传入一个结构体对象时，我们可以根据.来访问结构体的对应字段。例如：
**业务逻辑**

```go
package main
import ( 
    "net/http"
    "github.com/gin-gonic/gin"
)
type UserInfo struct {
    Name string
    Gender string
    Age int
}
func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/**/*")
    user := UserInfo{
        Name: "张三", Gender: "男", Age: 18, }
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "default/index.html", map[string]interface{}{ 
            "title": "前台首页", 
            "user": user, 
        })

    })
    router.Run(":8080")
}
```

**模板**
<!-- 相当于给模板定义一个名字 define end 成对出现-->

```html
{{ define "default/index.html" }}
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <h1>前台模板</h1>
        <h3>{{.title}}</h3>
        <h4>{{.user.Name}}</h4>
        <h4>{{.user.Age}}</h4>
    </body>
</html>
{{end}}
```

### 2、注释

```go
{{/* a comment */}}
```

注释，执行时会忽略。可以多行。注释不能嵌套，并且必须紧贴分界符始止。

### 3、变量

我们还可以在模板中声明变量，用来保存传入模板的数据或其他语句生成的结果。具体语法
如下：

```html
<h4>{{$obj := .title}}</h4>
<h4>{{$obj}}</h4>
```

### 4、移除空格

有时候我们在使用模板语法的时候会不可避免的引入一下空格或者换行符，这样模板最终渲染出来的内容可能就和我们想的不一样，这个时候可以使用{{-语法去除模板内容左侧的所有空白符号， 使用-}}去除模板内容右侧的所有空白符号。
例如：

```html
{{- .Name -}}
```

**注意**：-要紧挨{{和}}，同时与模板值之间需要使用空格分隔。

### 5、比较函数

布尔函数会将任何类型的零值视为假，其余视为真。
下面是定义为函数的二元比较运算的集合：
eq 如果 arg1 == arg2 则返回真
ne 如果 arg1 != arg2 则返回真
lt 如果 arg1 < arg2 则返回真
le 如果 arg1 <= arg2 则返回真
gt 如果 arg1 > arg2 则返回真
ge 如果 arg1 >= arg2 则返回真

### 6、条件判断

Go 模板语法中的条件判断有以下几种:

```html
{{if pipeline}} T1 {{end}}
{{if pipeline}} T1 {{else}} T0 {{end}}
{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
{{if gt .score 60}}
及格
{{else}}
不及格
{{end}}
{{if gt .score 90}}
优秀
{{else if gt .score 60}}
及格
{{else}}
不及格
{{end}}
```

### 7、range

Go 的模板语法中使用 range 关键字进行遍历，有以下两种写法，其中 pipeline 的值必须是数组、切片、字典或者通道。

```html
{{range $key,$value := .obj}}
{{$value}}
{{end}}
```

如果 pipeline 的值其长度为 0，不会有任何输出

```html
{{$key,$value := .obj}}
{{$value}}
{{else}}
pipeline 的值其长度为 0
{{end}}
```

如果 pipeline 的值其长度为 0，则会执行 T0。

```go
router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "default/index.html", map[string]interface{}{ 
        "hobby": []string{"吃饭", "睡觉", "写代码"}, 
    })
})
{{range $key,$value := .hobby}}
<p>{{$value}}</p>
{{end}}
```

### 8、With

```go
user := UserInfo{
    Name: "张三", Gender: "男", Age: 18, }
router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "default/index.html", map[string]interface{}{ 
        "user": user, 
    })
})
```

以前要输出数据：

```html
<h4>{{.user.Name}}</h4>
<h4>{{.user.Gender}}</h4>
<h4>{{.user.Age}}</h4>
```

现在要输出数据：

```html
{{with .user}}
<h4>姓名：{{.Name}}</h4>
<h4>性别：{{.user.Gender}}</h4>
<h4>年龄：{{.Age}}</h4>
{{end}}
```

简单理解：相当于 var .=.user

### 9、预定义函数 （了解）

执行模板时，函数从两个函数字典中查找：首先是模板函数字典，然后是全局函数字典。一般不在模板内定义函数，而是使用 Funcs 方法添加函数到模板里。

and
函数返回它的第一个 empty 参数或者最后一个参数；
就是说"and x y"等价于"if x then y else x"；所有参数都会执行；
or
返回第一个非 empty 参数或者最后一个参数；
亦即"or x y"等价于"if x then x else y"；所有参数都会执行；
not
返回它的单个参数的布尔值的否定
len
返回它的参数的整数类型长度
index
执行结果为第一个参数以剩下的参数为索引/键指向的值；
如"index x 1 2 3"返回 x[1][2][3]的值；每个被索引的主体必须是数组、切片或者字典。
print
即 fmt.Sprint
printf
即 fmt.Sprintf
println
即 fmt.Sprintln
html
返回与其参数的文本表示形式等效的转义 HTML。
这个函数在 html/template 中不可用。
urlquery
以适合嵌入到网址查询中的形式返回其参数的文本表示的转义值。
这个函数在 html/template 中不可用。
js
返回与其参数的文本表示形式等效的转义 JavaScript。
call
执行结果是调用第一个参数的返回值，该参数必须是函数类型，其余参数作为调用该函数的参数；
如"call .X.Y 1 2"等价于 go 语言里的 dot.X.Y(1, 2)；
其中 Y 是函数类型的字段或者字典的值，或者其他类似情况；
call 的第一个参数的执行结果必须是函数类型的值（和预定义函数如 print 明显不同）；
该函数类型值必须有 1 到 2 个返回值，如果有 2 个则后一个必须是 error 接口类型；
如果有 2 个返回值的方法返回的 error 非 nil，模板执行会中断并返回给调用模板执行者该错误；

```html
{{len .title}}
{{index .hobby 2}}
```

### 10、自定义模板函数

```go
router.SetFuncMap(template.FuncMap{ "formatDate": formatAsDate, })
```

```go
package main
import ( 
    "fmt"
    "html/template"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
)
func formatAsDate(t time.Time) string {
    year, month, day := t.Date()
    return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
func main() {
    router := gin.Default()
    //注册全局模板函数 注意顺序，注册模板函数需要在加载模板上面
    router.SetFuncMap(template.FuncMap{ "formatDate": formatAsDate, })
    //加载模板
    router.LoadHTMLGlob("templates/**/*")
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "default/index.html", map[string]interface{}{ "title": "前台首页", "now": time.Now(), })
    })
    router.Run(":8080")
}
{{.now | formatDate}}
或者
{{formatDate .now }}
```

## 5.4、嵌套 template

### 1、新建 templates/deafult/page_header.html

```html
{{ define "default/page_header.html" }}
<h1>这是一个头部</h1>
{{end}}
```

### 2、外部引入

注意：
1、引入的名字为 page_header.html 中定义的名字
2、引入的时候注意最后的点（.）

```html
{{template "default/page_header.html" .}}
```

```html
<!-- 相当于给模板定义一个名字 define end 成对出现-->
{{ define "default/index.html" }}
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        {{template "default/page_header.html" .}}
    </body>
</html>
{{end}}
```

# 六、静态文件服务

当我们渲染的 HTML 文件中引用了静态文件时,我们需要配置静态 web 服务

r.Static("/static", "./static") 前面的/static 表示路由 后面的./static 表示路径

```go
func main() {
    r := gin.Default()
    r.Static("/static", "./static")
    r.LoadHTMLGlob("templates/**/*")
    // ... r.Run(":8080")
}
<link rel="stylesheet" href="/static/css/base.css" />
```

# 七、路由详解

路由（Routing）是由一个 URI（或者叫路径）和一个特定的 HTTP 方法（GET、POST 等）组成的，涉及到应用如何响应客户端对某个网站节点的访问。

前面章节我们给大家介绍了路由基础以及路由配置，这里我们详细给大家讲讲路由传值、路由返回值

## 7.1、GET POST 以及获取 Get Post 传值

### 7.1.1、Get 请求传值

GET 	/user?uid=20&page=1

```go
router.GET("/user", func(c *gin.Context) {
    uid := c.Query("uid")
    page := c.DefaultQuery("page", "0")
    c.String(200, "uid=%v page=%v", uid, page)
})
```

### 7.1.2、动态路由传值

域名/user/20

```go
r.GET("/user/:uid", func(c *gin.Context) {
    uid := c.Param("uid")
    c.String(200, "userID=%s", uid)
})
```

### 7.1.3、Post 请求传值 获取 form 表单数据

定义一个 add_user.html 的页面

```html
{{ define "default/add_user.html" }}
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <form action="/doAddUser" method="post">
            用户名：<input type="text" name="username" />
            密码: <input type="password" name="password" />
            <input type="submit" value="提交">
        </form>
    </body>
</html>
{{end}}
```

通过 c.PostForm 接收表单传过来的数据

```go
router.GET("/addUser", func(c *gin.Context) {
    c.HTML(200, "default/add_user.html", gin.H{})
})
router.POST("/doAddUser", func(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")
    age := c.DefaultPostForm("age", "20")
    c.JSON(200, gin.H{ "usernmae": username, "password": password, "age": age, })
})
```

### 7.1.4、获取 GET POST 传递的数据绑定到结构体

为了能够更方便的获取请求相关参数，提高开发效率，我们可以基于请求的 Content-Type识别请求数据类型并利用反射机制自动提取请求中 QueryString、form 表单、JSON、XML 等参数到结构体中。 下面的示例代码演示了.ShouldBind()强大的功能，它能够基于请求自动提取 JSON、form 表单和 QueryString 类型的数据，并把值绑定到指定的结构体对象。

```go
//注意首字母大写
type Userinfo struct {
    Username string `form:"username" json:"user"` 
    Password string `form:"password" json:"password"` }
```

**Get 传值绑定到结构体**

/?username=zhangsan&password=123456

```go
router.GET("/", func(c *gin.Context) {
    var userinfo Userinfo
    if err := c.ShouldBind(&userinfo); err == nil {
        c.JSON(http.StatusOK, userinfo)
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
})
```

返回数据
{"user":"zhangsan","password":"123456"}

**Post 传值绑定到结构体**

```go
router.POST("/doLogin", func(c *gin.Context) {
    var userinfo Userinfo
    if err := c.ShouldBind(&userinfo); err == nil {
        c.JSON(http.StatusOK, userinfo)
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
})
```

返回数据
{"user":"zhangsan","password":"123456"}

### 7.1.5、获取 Post Xml 数据

在 API 的开发中，我们经常会用到 JSON 或 XML 来作为数据交互的格式，这个时候我们可以在 gin 中使用 c.GetRawData()获取数据。

```html
<?xml version="1.0" encoding="UTF-8"?>
<article>
    <content type="string">我是张三</content>
    <title type="string">张三</title>
</article>
```

![image2](assets/image2.png)

```go
type Article struct {
    Title string `xml:"title"` Content string `xml:"content"` }
router.POST("/xml", func(c *gin.Context) {
    b, _ := c.GetRawData() // 从 c.Request.Body 读取请求数据
    article := &Article{}
    if err := xml.Unmarshal(b, &article); err == nil {
        c.JSON(http.StatusOK, article)
    } else {
        c.JSON(http.StatusBadRequest, err.Error())
    }
})
```

## 7.2、简单的路由组

https://gin-gonic.com/zh-cn/docs/examples/grouping-routes/

```go
func main() {
    router := gin.Default()
    // 简单的路由组: v1
    v1 := router.Group("/v1")
    {
        v1.POST("/login", loginEndpoint)
        v1.POST("/submit", submitEndpoint)
        v1.POST("/read", readEndpoint)
    }
    // 简单的路由组: v2
    v2 := router.Group("/v2")
    {
        v2.POST("/login", loginEndpoint)
        v2.POST("/submit", submitEndpoint)
        v2.POST("/read", readEndpoint)
    }
    router.Run(":8080")
}
```

## 7.3、Gin 路由文件 分组

### 7.3.1、新建 routes 文件夹，routes 文件下面新建 adminRoutes.go、apiRoutes.go、defaultRoutes.go

#### 1、新建 adminRoutes.go

```go
package routes
import ( 
    "net/http"
    "github.com/gin-gonic/gin"
)
func AdminRoutesInit(router *gin.Engine) {
    adminRouter := router.Group("/admin")
    {
        adminRouter.GET("/user", func(c *gin.Context) {
            c.String(http.StatusOK, "用户")
        })
        adminRouter.GET("/news", func(c *gin.Context) {
            c.String(http.StatusOK, "news")
        })
    }
}
```

#### 2、新建 apiRoutes.go

```go
package routes
import ( "net/http"
        "github.com/gin-gonic/gin"
       )
func ApiRoutesInit(router *gin.Engine) {
    apiRoute := router.Group("/api")
    {
        apiRoute.GET("/user", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{ "username": "张三",
                                        "age": 20, })
        })
        apiRoute.GET("/news", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{ "title": "这是新闻", })
        })
    }
}
```

#### 3、新建 defaultRoutes.go

```go
package routes
import ( "github.com/gin-gonic/gin"
       )
func DefaultRoutesInit(router *gin.Engine) {
    defaultRoute := router.Group("/")
    {
        defaultRoute.GET("/", func(c *gin.Context) {
            c.String(200, "首页")
        })
    }
}
```

### 7.3.2 、配置 main.go

```go
package main
import ( "gin_demo/routes"
        "github.com/gin-gonic/gin"
       )
//注意首字母大写
type Userinfo struct {
    Username string `form:"username" json:"user"` Password string `form:"password" json:"password"` }
func main() {
    r := gin.Default()
    routes.AdminRoutesInit(r)
    routes.ApiRoutesInit(r)
    routes.DefaultRoutesInit(r)
    r.Run(":8080")
}
```

访问 /api/user /admin/user 测试

# 八、Gin 中自定义控制器

## 8.1、控制器分组

当我们的项目比较大的时候有必要对我们的控制器进行分组
新建 controller/admin/NewsController.go

```go
package admin
import ( "net/http"
        "github.com/gin-gonic/gin"
       )
type NewsController struct {
}
func (c NewsController) Index(ctx *gin.Context) {
    ctx.String(http.StatusOK, "新闻首页")
}
```

新建 controller/admin/UserController.go

```go
package admin
import ( "net/http"
        "github.com/gin-gonic/gin"
       )
type UserController struct {
}
func (c UserController) Index(ctx *gin.Context) {
    ctx.String(http.StatusOK, "这是用户首页")
}
func (c UserController) Add(ctx *gin.Context) {
    ctx.String(http.StatusOK, "增加用户")
}
```

配置对应的路由 --adminRoutes.go

其他路由的配置方法类似

```go
package routes
import ( "gin_demo/controller/admin"
        "net/http"
        "github.com/gin-gonic/gin"
       )
func AdminRoutesInit(router *gin.Engine) {
    adminRouter := router.Group("/admin")
    {
        adminRouter.GET("/user", admin.UserController{}.Index)
        adminRouter.GET("/user/add", admin.UserController{}.Add)
        adminRouter.GET("/news", admin.NewsController{}.Add)
    }
}
```

## 8.2、控制器的继承

### 1、新建 controller/admin/BaseController.go

```go
package admin
import ( "net/http"
        "github.com/gin-gonic/gin"
       )
type BaseController struct {
}
func (c BaseController) Success(ctx *gin.Context) {
    ctx.String(http.StatusOK, "成功")
}
func (c BaseController) Error(ctx *gin.Context) {
    ctx.String(http.StatusOK, "失败")
}
```

### 2、NewsController 继承 BaseController

继承后就可以调用控制器里面的公共方法了

```go
package admin
import ( "github.com/gin-gonic/gin"
       )
type NewsController struct {
    BaseController
}
func (c NewsController) Index(ctx *gin.Context) {
    c.Success(ctx)
}
```

# 九、Gin 中间件

Gin 框架允许开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数。这个钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、**耗时统计**等。

通俗的讲：中间件就是匹配路由前和匹配路由完成后执行的一系列操作

## 9.1、路由中间件

### 9.1.1、初识中间件

Gin 中的中间件必须是一个 gin.HandlerFunc 类型，配置路由的时候可以传递多个 func 回调函数，最后一个 func 回调函数前面触发的方法都可以称为中间件。


```go
package main
import ( "fmt"
        "github.com/gin-gonic/gin"

       )
func initMiddleware(ctx *gin.Context) {
    fmt.Println("我是一个中间件")
}
func main() {
    r := gin.Default()
    r.GET("/", initMiddleware, func(ctx *gin.Context) {
        ctx.String(200, "首页--中间件演示")
    })
    r.GET("/news", initMiddleware, func(ctx *gin.Context) {
        ctx.String(200, "新闻页面--中间件演示")
    })
    r.Run(":8080")
}
```

### 9.1.2、ctx.Next()调用该请求的剩余处理程序

中间件里面加上 ctx.Next()可以让我们在路由匹配完成后执行一些操作。
比如我们统计一个请求的执行时间

```go
package main
import ( "fmt"
        "time"
        "github.com/gin-gonic/gin"
       )
func initMiddleware(ctx *gin.Context) {
    fmt.Println("1-执行中中间件")
    start := time.Now().UnixNano()
    // 调用该请求的剩余处理程序
    ctx.Next()
    fmt.Println("3-程序执行完成 计算时间")
    // 计算耗时 Go 语言中的 Since()函数保留时间值，并用于评估与实际时间的差异
    end := time.Now().UnixNano()
    fmt.Println(end - start)
}
func main() {
    r := gin.Default()
    r.GET("/", initMiddleware, func(ctx *gin.Context) {
        fmt.Println("2-执行首页返回数据")
        ctx.String(200, "首页--中间件演示")
    })
    r.GET("/news", initMiddleware, func(ctx *gin.Context) {
        ctx.String(200, "新闻页面--中间件演示")
    })
    r.Run(":8080")
}
```

### 9.1.3、一个路由配置多个中间件的执行顺序

```go
func initMiddlewareOne(ctx *gin.Context) {
    fmt.Println("initMiddlewareOne--1-执行中中间件")
    // 调用该请求的剩余处理程序
    ctx.Next()
    fmt.Println("initMiddlewareOne--2-执行中中间件")
}
func initMiddlewareTwo(ctx *gin.Context) {
    fmt.Println("initMiddlewareTwo--1-执行中中间件")
    // 调用该请求的剩余处理程序
    ctx.Next()
    fmt.Println("initMiddlewareTwo--2-执行中中间件")
}
func main() {
    r := gin.Default()
    r.GET("/", initMiddlewareOne, initMiddlewareTwo, func(ctx *gin.Context) {
        fmt.Println("执行路由里面的程序")
        ctx.String(200, "首页--中间件演示")
    })
    r.Run(":8080")
}
```

控制台内容：
initMiddlewareOne--1-执行中中间件
initMiddlewareTwo--1-执行中中间件
执行路由里面的程序
initMiddlewareTwo--2-执行中中间件
initMiddlewareOne--2-执行中中间件

### 9.1.4、 c.Abort()--（了解）

Abort 是终止的意思， c.Abort() 表示终止调用该请求的剩余处理程序

```go
package main
import ( "fmt"
        "github.com/gin-gonic/gin"
       )
func initMiddlewareOne(ctx *gin.Context) {
    fmt.Println("initMiddlewareOne--1-执行中中间件")
    // 调用该请求的剩余处理程序
    ctx.Next()
    fmt.Println("initMiddlewareOne--2-执行中中间件")
}
func initMiddlewareTwo(ctx *gin.Context) {
    fmt.Println("initMiddlewareTwo--1-执行中中间件")
    // 终止调用该请求的剩余处理程序
    ctx.Abort()
    fmt.Println("initMiddlewareTwo--2-执行中中间件")
}
func main() {
    r := gin.Default()
    r.GET("/", initMiddlewareOne, initMiddlewareTwo, func(ctx *gin.Context) {
        fmt.Println("执行路由里面的程序")
        ctx.String(200, "首页--中间件演示")
    })
    r.Run(":8080")
}
```

initMiddlewareOne--1-执行中间件
initMiddlewareTwo--1-执行中间件
initMiddlewareTwo--2-执行中间件
initMiddlewareOne--2-执行中间件

## 9.2、全局中间件

```go
package main
import ( "fmt"
        "github.com/gin-gonic/gin"
       )
func initMiddleware(ctx *gin.Context) {
    fmt.Println("全局中间件 通过 r.Use 配置")
    // 调用该请求的剩余处理程序
    ctx.Next()
}
func main() {
    r := gin.Default()
    r.Use(initMiddleware)
    r.GET("/", func(ctx *gin.Context) {
        ctx.String(200, "首页--中间件演示")
    })
    r.GET("/news", func(ctx *gin.Context) {
        ctx.String(200, "新闻页面--中间件演示")
    })
    r.Run(":8080")
}
```

## 9.3、在路由分组中配置中间件

### 1、为路由组注册中间件有以下两种写法

写法 1：

```go
shopGroup := r.Group("/shop", StatCost())
{
shopGroup.GET("/index", func(c *gin.Context) {...})
... }
```

写法 2：

```go
shopGroup := r.Group("/shop")
shopGroup.Use(StatCost())
{
shopGroup.GET("/index", func(c *gin.Context) {...})
... }
```

### 2、分组路由 AdminRoutes.go 中配置中间件

```go
package routes
import ( "fmt"
        "gin_demo/controller/admin"
        "net/http"
        "github.com/gin-gonic/gin"
       )
func initMiddleware(ctx *gin.Context) {
    fmt.Println("路由分组中间件")
    // 调用该请求的剩余处理程序
    ctx.Next()
}
func AdminRoutesInit(router *gin.Engine) {
    adminRouter := router.Group("/admin", initMiddleware)
    {
        adminRouter.GET("/user", admin.UserController{}.Index)
        adminRouter.GET("/user/add", admin.UserController{}.Add)
        adminRouter.GET("/news", func(c *gin.Context) {
            c.String(http.StatusOK, "news")
        })
    }
}
```

## 9.4、中间件和对应控制器之间共享数据

设置值

```go
ctx.Set("username", "张三")
```

获取值

```go
username, _ := ctx.Get("username")
```

中间件设置值

```go
func InitAdminMiddleware(ctx *gin.Context) {
    fmt.Println("路由分组中间件")
    // 可以通过 ctx.Set 在请求上下文中设置值，后续的处理函数能够取到该值
    ctx.Set("username", "张三")
    // 调用该请求的剩余处理程序
    ctx.Next()
}
```

控制器获取值

```go
func (c UserController) Index(ctx *gin.Context) {
    username, _ := ctx.Get("username")
    fmt.Println(username)
    ctx.String(http.StatusOK, "这是用户首页 111")
}
```

## 9.5、中间件注意事项

**gin 默认中间件**

gin.Default()默认使用了 Logger 和 Recovery 中间件，其中：

- Logger 中间件将日志写入 gin.DefaultWriter，即使配置了 GIN_MODE=release。
- Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500 响应码。

如果不想使用上面两个默认的中间件，可以使用 `gin.New()`新建一个没有任何默认中间件的路由。

**gin 中间件中使用 goroutine**
当在中间件或 handler 中启动新的 goroutine 时，**不能使用**原始的上下文（c *gin.Context），
必须使用其只读副本（c.Copy()）

```go
r.GET("/", func(c *gin.Context) {
    cCp := c.Copy()
    go func() {
        // simulate a long task with time.Sleep(). 5 seconds
        time.Sleep(5 * time.Second)
        // 这里使用你创建的副本
        fmt.Println("Done! in path " + cCp.Request.URL.Path)
    }()
    c.String(200, "首页")
})
```

# 十、Gin 中自定义 Model

## 10.1、关于 Model

如果我们的应用非常简单的话，我们可以在 Controller 里面处理常见的业务逻辑。但是如果我们有一个功能想在多个控制器、或者多个模板里面复用的话，那么我们就可以把公共的功能单独抽取出来作为一个模块（Model）。 Model 是逐步抽象的过程，一般我们会在 Model里面封装一些公共的方法让不同 Controller 使用，也可以在 Model 中实现和数据库打交道

## 10.2、Model 里面封装公共的方法

### 1、新建 models/ tools.go

```go
package models
import ( "crypto/md5"
        "fmt"
        "time"
        "github.com/astaxie/beego"
       )
//时间戳间戳转换成日期
func UnixToDate(timestamp int) string {
    t := time.Unix(int64(timestamp), 0)
    return t.Format("2006-01-02 15:04:05")
}
//日期转换成时间戳 2020-05-02 15:04:05
func DateToUnix(str string) int64 {
    template := "2006-01-02 15:04:05"
    t, err := time.ParseInLocation(template, str, time.Local)
    if err != nil {
        return 0
    }
    return t.Unix()
}
func GetUnix() int64 {
    return time.Now().Unix()
}
func GetDate() string {
    template := "2006-01-02 15:04:05"
    return time.Now().Format(template)
}
func GetDay() string {
    template := "20060102"
    return time.Now().Format(template)
}
func Md5(str string) string {
    data := []byte(str)
    return fmt.Sprintf("%x\n", md5.Sum(data))
}
```

## 10.3、控制器中调用 Model

```go
package controllers
import ( "gin_demo/models"
       )
day := models.GetDay()
```

## 10.4、调用 Model 注册全局模板函数

models/tools.go
//时间戳间戳转换成日期

```go
func UnixToDate(timestamp int64) string {

    t := time.Unix(timestamp, 0)
    return t.Format("2006-01-02 15:04:05")
}
```

main.go

//注册全局模板函数 注意顺序，注册模板函数需要在加载模板上面

```go
r := gin.Default()
r.SetFuncMap(template.FuncMap{ "unixToDate": models.UnixToDate, })
```

**控制器**

```go
func (c UserController) Add(ctx *gin.Context) {
    ctx.HTML(http.StatusOK, "admin/user/add.html", gin.H{ "now": models.GetUnix(), })
}
```

**模板**

```html
<h2>{{.now | unixToDate}}</h2>
```

## 10.5、Golang Md5 加密

打开 golang 包对应的网站：https://pkg.go.dev/，搜索 md5
方法一：

```go
data := []byte("123456")
has := md5.Sum(data)
md5str := fmt.Sprintf("%x", has)
fmt.Println(md5str)
```

方法二：

```go
h := md5.New()
io.WriteString(h, "123456")
fmt.Printf("%x\n", h.Sum(nil))
```

# 十一、Gin 文件上传

**注意：需要在上传文件的 form 表单上面需要加入 enctype="multipart/form-data"**

## 11.1、单文件上传

https://gin-gonic.com/zh-cn/docs/examples/upload-file/single-file/

官方示例：

```go
func main() {
    router := gin.Default()
    // 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
    router.MaxMultipartMemory = 8 << 20 // 8 MiB
    router.POST("/upload", func(c *gin.Context) {
        // 单文件
        file, _ := c.FormFile("file")
        log.Println(file.Filename)
        // 上传文件至指定目录
        c.SaveUploadedFile(file, dst)
        c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
    })
    router.Run(":8080")
}
```

项目中实现文件上传：

1、定义模板 需要在上传文件的 form 表单上面需要加入enctype="multipart/form-data

```html
<!-- 相当于给模板定义一个名字 define end 成对出现-->
{{ define "admin/user/add.html" }}
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <form action="/admin/user/doAdd" method="post" enctype="multipart/form-data">
            用户名： <input type="text" name="username" placeholder="用户名"> <br> <br>
            头 像：<input type="file" name="face"><br> <br>
            <input type="submit" value="提交">
        </form>
    </body>
</html>
{{ end }}
```

2、定义业务逻辑

```go
func (c UserController) DoAdd(ctx *gin.Context) {
    username := ctx.PostForm("username")
    file, err := ctx.FormFile("face")
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{ "message": err.Error(), })
        return
    }
    // 上传文件到指定的目录
    dst := path.Join("./static/upload", file.Filename)
    fmt.Println(dst)
    ctx.SaveUploadedFile(file, dst)
    ctx.JSON(http.StatusOK, gin.H{ "message": fmt.Sprintf("'%s' uploaded!", file.Filename), "username": username, })
}
```

## 11.2、多文件上传--不同名字的多个文件

1、定义模板 需要在上传文件的 form 表单上面需要加入 enctype="multipart/form-data"

```html
<!-- 相当于给模板定义一个名字 define end 成对出现-->
{{ define "admin/user/add.html" }}
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <form action="/admin/user/doAdd" method="post" enctype="multipart/form-data">
            用户名： <input type="text" name="username" placeholder="用户名"> <br> <br>
            头 像 1：<input type="file" name="face1"><br> <br>
            头 像 2：<input type="file" name="face2"><br> <br>
            <input type="submit" value="提交">
        </form>
    </body>
</html>
{{ end }}
```

2、定义业务逻辑

```go
func (c UserController) DoAdd(ctx *gin.Context) {
    username := ctx.PostForm("username")
    face1, err1 := ctx.FormFile("face1")
    face2, err2 := ctx.FormFile("face2")
    // 上传文件到指定的目录
    if err1 == nil {
        dst1 := path.Join("./static/upload", face1.Filename)
        ctx.SaveUploadedFile(face1, dst1)
    }
    if err2 == nil {
        dst2 := path.Join("./static/upload", face2.Filename)
        ctx.SaveUploadedFile(face2, dst2)
    }
    ctx.JSON(http.StatusOK, gin.H{
        "message": "文件上传成功", "username": username, })
    // ctx.String(200, username)
}
```

## 11.3、多文件上传--相同名字的多个文件

参考：https://gin-gonic.com/zh-cn/docs/examples/upload-file/multiple-file/

1、定义模板 需要在上传文件的 form 表单上面需要加入 enctype="multipart/form-data"

```html
<!-- 相当于给模板定义一个名字 define end 成对出现-->
{{ define "admin/user/add.html" }}
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <form action="/admin/user/doAdd" method="post" enctype="multipart/form-data">
            用户名： <input type="text" name="username" placeholder="用户名"> <br> <br>
            头 像 1：<input type="file" name="face[]"><br> <br>
            头 像 2：<input type="file" name="face[]"><br> <br>
            <input type="submit" value="提交">
        </form>
    </body>
</html>
{{ end }}
```

2、定义业务逻辑

```go
func (c UserController) DoAdd(ctx *gin.Context) {
    username := ctx.PostForm("username")
    // Multipart form
    form, _ := ctx.MultipartForm()
    files := form.File["face[]"]
    // var dst;
    for _, file := range files {
        // 上传文件至指定目录
        dst := path.Join("./static/upload", file.Filename)
        ctx.SaveUploadedFile(file, dst)
    }
    ctx.JSON(http.StatusOK, gin.H{ "message": "文件上传成功", "username": username, })
}
```

11.4、文件上传 按照日期存储

1、定义模板 需要在上传文件的 form 表单上面需要加入 enctype="multipart/form-data"

```html
<!-- 相当于给模板定义一个名字 define end 成对出现-->
{{ define "admin/user/add.html" }}
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <form action="/admin/user/doAdd" method="post" enctype="multipart/form-data">
            用户名： <input type="text" name="username" placeholder="用户名"> <br> <br>
            头 像： <input type="file" name="face"><br> <br>
            <input type="submit" value="提交">
        </form>
    </body>
</html>
{{ end }}
```

2、定义业务逻辑

```go
func (c UserController) DoAdd(ctx *gin.Context) {
    username := ctx.PostForm("username")
    //1、获取上传的文件
    file, err1 := ctx.FormFile("face")
    if err1 == nil {
        //2、获取后缀名 判断类型是否正确 .jpg .png .gif .jpeg
        extName := path.Ext(file.Filename)
        allowExtMap := map[string]bool{ ".jpg": true, ".png": true, ".gif": true, ".jpeg": true, }
        if _, ok := allowExtMap[extName]; !ok {
            ctx.String(200, "文件类型不合法")
            return
        }
        //3、创建图片保存目录 static/upload/20200623
        day := models.GetDay()
        dir := "./static/upload/" + day
        if err := os.MkdirAll(dir, 0666); err != nil {
            log.Error(err)
        }
        //4、生成文件名称 144325235235.png
        fileUnixName := strconv.FormatInt(models.GetUnix(), 10)
        //static/upload/20200623/144325235235.png
        saveDir := path.Join(dir, fileUnixName+extName)
        ctx.SaveUploadedFile(file, saveDir)
    }
    ctx.JSON(http.StatusOK, gin.H{ "message": "文件上传成功", "username": username, })
    // ctx.String(200, username)
}
```

3、models/tools.go

```go
package models
import ( "crypto/md5"
        "fmt"
        "time"
       )
//时间戳间戳转换成日期
func UnixToDate(timestamp int) string {
    t := time.Unix(int64(timestamp), 0)
    return t.Format("2006-01-02 15:04:05")
}
//日期转换成时间戳 2020-05-02 15:04:05
func DateToUnix(str string) int64 {
    template := "2006-01-02 15:04:05"
    t, err := time.ParseInLocation(template, str, time.Local)
    if err != nil {
        beego.Info(err)
        return 0
    }
    return t.Unix()
}
func GetUnix() int64 {
    return time.Now().Unix()
}
func GetDate() string {
    template := "2006-01-02 15:04:05"
    return time.Now().Format(template)
}
func GetDay() string {
    template := "20060102"
    return time.Now().Format(template)
}
func Md5(str string) string {
    data := []byte(str)
    return fmt.Sprintf("%x\n", md5.Sum(data))
}
func Hello(in string) (out string) {
    out = in + "world"
    return
}
```

# 十二、Gin 中的 Cookie

## 12.1、Cookie 介绍

- HTTP 是无状态协议。简单地说，当你浏览了一个页面，然后转到同一个网站的另一个页面，服务器无法认识到这是同一个浏览器在访问同一个网站。每一次的访问，都是没有任何关系的。如果我们要实现多个页面之间共享数据的话我们就可以使用 Cookie 或者 Session 实现
- cookie 是存储于访问者计算机的浏览器中。可以让我们用同一个浏览器访问同一个域名的时候共享数据。

## 12.2、Cookie 能实现的功能

1. 保持用户登录状态
2. 保存用户浏览的历史记录
3. 猜你喜欢，智能推荐
4. 电商网站的加入购物车

## 12.3、设置和获取 Cookie

https://gin-gonic.com/zh-cn/docs/examples/cookie/

**设置 Cookie**

```go
c.SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
```

第一个参数 key
第二个参数 value
第三个参数 过期时间.如果只想设置 Cookie 的保存路径而不想设置存活时间，可以在第三个参数中传递 nil
第四个参数 cookie 的路径
第五个参数 cookie 的路径 Domain 作用域 本地调试配置成 localhost , 正式上线配置成域名
第六个参数是 secure ，当 secure 值为 true 时，cookie 在 HTTP 中是无效，在 HTTPS 中才有效
第七个参数 httpOnly，是微软对 COOKIE 做的扩展。如果在 COOKIE 中设置了“httpOnly”属性，则通过程序（JS 脚本、applet 等）将无法读取到 COOKIE 信息，防止 XSS 攻击产生

**获取 Cookie**

```
cookie, err := c.Cookie("name")
```

**完整 demo**

```go
package main
import ( "gin_demo/models"
        "html/template"
        "github.com/gin-gonic/gin"
       )
func main() {
    r := gin.Default()
    r.SetFuncMap(template.FuncMap{ "unixToDate": models.UnixToDate, })
    r.GET("/", func(c *gin.Context) {
        c.SetCookie("usrename", "张三", 3600, "/", "localhost", false, true)
        c.String(200, "首页")
    })
    r.GET("/user", func(c *gin.Context) {
        username, _ := c.Cookie("usrename")
        c.String(200, "用户-"+username)
    })
    r.Run(":8080")
}
```

## 12.4 、多个二级域名共享 cookie

1、分别把 a.itying.com 和 b.itying.com 解析到我们的服务器
2、我们想的是用户在 a.itying.com 中设置 Cookie 信息后在 b.itying.com 中获取刚才设置的cookie，也就是实现多个二级域名共享 cookie

这时候的话我们就可以这样设置 cookie

```go
c.SetCookie("usrename", "张三", 3600, "/", ".itying.com", false, true)
```

# 十三、Gin 中的 Session

## 13.1、Session 简单介绍

session 是另一种记录客户状态的机制，不同的是 Cookie 保存在客户端浏览器中，而 session保存在服务器上。

## 13.2、Session 的工作流程

当客户端浏览器第一次访问服务器并发送请求时，服务器端会创建一个 session 对象，生成一个类似于 key,value 的键值对，然后将 value 保存到服务器 将 key(cookie)返回到浏览器(客户)端。浏览器下次访问时会携带 key(cookie)，找到对应的 session(value)。

## 13.3、Gin 中使用 Session

Gin 官方没有给我们提供 Session 相关的文档，这个时候我们可以使用第三方的 Session 中间件来实现

- https://github.com/gin-contrib/sessions
  gin-contrib/sessions 中间件支持的存储引擎：
- cookie
- memstore
- redis 
- memcached
- mongodb

## 13.4、基于 Cookie 存储 Session

### 1、安装 session 包

```
go get github.com/gin-contrib/sessions
```

### 2、基本的 session 用法

```go
package main
import ( "github.com/gin-contrib/sessions"
        "github.com/gin-contrib/sessions/cookie"
        "github.com/gin-gonic/gin"
       )
func main() {
    r := gin.Default()
    // 创建基于 cookie 的存储引擎，secret11111 参数是用于加密的密钥
    store := cookie.NewStore([]byte("secret11111"))
    // 设置 session 中间件，参数 mysession，指的是 session 的名字，也是 cookie 的名字
    // store 是前面创建的存储引擎，我们可以替换成其他存储引擎
    r.Use(sessions.Sessions("mysession", store))
    r.GET("/", func(c *gin.Context) {
        //初始化 session 对象
        session := sessions.Default(c)
        //设置过期时间
        session.Options(sessions.Options{
            MaxAge: 3600 * 6, // 6hrs
        })
        //设置 Session
        session.Set("username", "张三")
        session.Save()
        c.JSON(200, gin.H{"msg": session.Get("username")})
    })
    r.GET("/user", func(c *gin.Context) {
        // 初始化 session 对象
        session := sessions.Default(c)
        // 通过 session.Get 读取 session 值
        username := session.Get("username")
        c.JSON(200, gin.H{"username": username})
    })
    r.Run(":8000")
}
```

## 13.5、基于 Redis 存储 Session

如果我们想将 session 数据保存到 redis 中，只要将 session 的存储引擎改成 redis 即可。使用 redis 作为存储引擎的例子：
首先安装 redis 存储引擎的包

```
go get github.com/gin-contrib/sessions/redis
```

例子：

```go
package main
import ( "github.com/gin-contrib/sessions"
        "github.com/gin-contrib/sessions/redis"
        "github.com/gin-gonic/gin"
       )
func main() {
    r := gin.Default()
    // 初始化基于 redis 的存储引擎
    // 参数说明：
    // 第 1 个参数 - redis 最大的空闲连接数
    // 第 2 个参数 - 数通信协议 tcp 或者 udp
    // 第 3 个参数 - redis 地址, 格式，host:port
    // 第 4 个参数 - redis 密码
    // 第 5 个参数 - session 加密密钥
    store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
    r.Use(sessions.Sessions("mysession", store))
    r.GET("/", func(c *gin.Context) {
        session := sessions.Default(c)
        session.Set("username", "李四")
        session.Save()
        c.JSON(200, gin.H{"username": session.Get("username")})
    })
    r.GET("/user", func(c *gin.Context) {
        // 初始化 session 对象
        session := sessions.Default(c)
        // 通过 session.Get 读取 session 值
        username := session.Get("username")
        c.JSON(200, gin.H{"username": username})
    })
    r.Run(":8000")
}
```

