package main

import "fmt"

func main() {
	// rune 实际上是 int32 的别名
	var r1 rune = 'A'          // ASCII 字符
	var r2 rune = '中'          // Unicode 字符
	var r3 rune = '\u4e2d'     // Unicode 转义
	var r4 rune = '\U0001f600' // emoji 😊

	fmt.Printf("字符: %c, Unicode: %U\n", r1, r1)
	fmt.Printf("字符: %c, Unicode: %U\n", r2, r2)
	fmt.Printf("字符: %c, Unicode: %U\n", r3, r3)
	fmt.Printf("字符: %c, Unicode: %U\n", r4, r4)

	// 字符串遍历
	text := "Hello世界"
	fmt.Println("按字符遍历:")
	for i, char := range text {
		fmt.Printf("位置: %d, 字符: %c, Unicode: %U\n", i, char, char)
	}
}
