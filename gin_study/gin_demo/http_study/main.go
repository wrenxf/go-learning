package main

// 标准库 HTTP 服务器 - 用于对比学习 Gin 框架的区别

import (
	"fmt"
	"net/http"
)

// sayhello 是处理函数（Handler）
// 参数说明：
//   - w http.ResponseWriter: 用于向客户端写入响应数据
//   - r *http.Request: 包含客户端请求的所有信息（URL、Header、Body等）
func sayhello(w http.ResponseWriter, r *http.Request) {
	// Fprintln 将 "hello golang" 写入响应体
	// 返回值被忽略（_），因为这里不需要检查写入是否成功
	_, _ = fmt.Fprintln(w, "hello golang")
}

func main() {
	// 注册路由：将 /hello 路径绑定到 sayhello 处理函数
	// 当访问 http://localhost:9090/hello 时，会调用 sayhello
	http.HandleFunc("/hello", sayhello)

	// 启动 HTTP 服务器，监听 9090 端口
	// 第二个参数 nil 表示使用 DefaultServeMux（默认路由器）
	// 注意：ListenAndServe 是阻塞的，会一直运行直到出错
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		// 如果服务器启动失败，打印错误信息并退出
		// 注意：这里用的是 println 而不是 fmt.Println
		fmt.Printf("http serve failed, err:%v\n", err)
		return
	}
}
