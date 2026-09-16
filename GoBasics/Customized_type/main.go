package main

import "fmt"

// 定义新类型
type UserID int
type UserName string
type Score float64

// 类型别名（Go 1.9+）
type MyInt = int // 类型别名
type MyString = string

// 为自定义类型添加方法
func (u UserID) String() string {
	return fmt.Sprintf("用户#%d", u)
}

func (n UserName) Greet() string {
	return fmt.Sprintf("你好，%s！", n)
}

func main() {
	var id UserID = 1001
	var name UserName = "张三"
	var score Score = 95.5

	fmt.Println(id.String())
	fmt.Println(name.Greet())
	fmt.Printf("分数: %.1f\n", score)

	// 类型转换（即使底层类型相同，也是不同类型）
	var normalInt1 int = 42
	var customID UserID = UserID(normalInt1) // 需要显式转换
	fmt.Printf("普通int: %d, UserID: %s\n", normalInt1, customID)

	var i MyInt = 42
	var s MyString = "Hello"

	fmt.Printf("MyInt: %d, MyString: %s\n", i, s)

	// 类型别名与原类型完全相同，可以直接赋值
	var normalInt2 int = i // 不需要显式转换
	var normalStr string = s

	fmt.Printf("转换后的: %d, %s\n", normalInt2, normalStr)
}
