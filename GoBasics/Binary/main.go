package main

import "fmt"

func main() {
	a := 0b1010
	b := 0b0101
	fmt.Printf("a=%04b (%d)\n", a, a)
	fmt.Printf("b=%04b (%d)\n", b, b)

	fmt.Printf("%04b & %04b = %04b (%d)\n", a, b, a&b, a&b) // 按位与
	fmt.Printf("%04b | %04b = %04b (%d)\n", a, b, a|b, a|b) // 按位或
	fmt.Printf("%04b ^ %04b = %04b (%d)\n", a, b, a^b, a^b) // 按位异或
	fmt.Printf("^%04b = %04b (%d)\n", a, ^a, ^a)            // 按位取反
	// 移位运算
	x := 8                                               // 0b1000
	fmt.Printf("%04b << 2 = %04b (%d)\n", x, x<<2, x<<2) // 左移
	fmt.Printf("%04b >> 1 = %04b (%d)\n", x, x>>1, x>>1) // 右移

	// 实际应用：权限管理
	const (
		READ    = 1 << 0 // 1  (0b0001)
		WRITE   = 1 << 1 // 2  (0b0010)
		EXECUTE = 1 << 2 // 4  (0b0100)
		DELETE  = 1 << 3 // 8  (0b1000)
	)

	// 用户权限
	var permissions uint8 = READ | WRITE // 3 (0b0011)
	fmt.Printf("初始权限: %04b (%d)\n", permissions, permissions)

	// 添加执行权限
	permissions |= EXECUTE
	fmt.Printf("添加执行权限: %04b (%d)\n", permissions, permissions)

	// 检查权限
	hasRead := (permissions & READ) != 0
	hasWrite := (permissions & WRITE) != 0
	hasExecute := (permissions & EXECUTE) != 0

	fmt.Printf("可读: %t, 可写: %t, 可执行: %t\n", hasRead, hasWrite, hasExecute)

	// 移除写权限
	permissions &= ^uint8(WRITE)
	fmt.Printf("移除写权限: %04b (%d)\n", permissions, permissions)
}
