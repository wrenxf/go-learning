package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 数值类型转换
	var i int = 42
	var f float64 = float64(i)
	var b byte = byte(i)

	fmt.Printf("int: %d, float64: %f, byte: %d\n", i, f, b)

	// 字符串与数值转换
	str := strconv.Itoa(i)       // int 转 string
	newI, _ := strconv.Atoi(str) // string 转 int

	fmt.Printf("int->string: %s, string->int: %d\n", str, newI)

	// 字符串与浮点数转换
	f = 3.14159
	strFloat := strconv.FormatFloat(f, 'f', 2, 64)
	newF, _ := strconv.ParseFloat(strFloat, 64)

	fmt.Printf("float64->string: %s, string->float64: %f\n", strFloat, newF)

	// 注意：不同类型之间不能直接运算
	// var result = i + f  // 编译错误！
	var result = float64(i) + f // 正确做法
	fmt.Printf("42 + 3.14159 = %f\n", result)
}
