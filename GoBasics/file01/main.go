package main

import (
	"fmt"
	"os"
)

func main() {
	// 只读方式打开当前目录下的 main.go 文件
	file, err := os.Open("D:/桌面/go-learning/GoBasics/file01/main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	fmt.Println(file)  //&{0xc000078780}
	defer file.Close() // 关闭文件
}
