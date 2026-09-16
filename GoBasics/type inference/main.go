package main

import "fmt"

func main() {
	// 自动类型推断
	i := 42                 // int
	f := 3.14               // float64
	s := "Hello"            // string
	b := true               // bool
	arr := [3]int{1, 2, 3}  // [3]int
	slice := []int{1, 2, 3} // []int

	fmt.Printf("类型: %T, 值: %v\n", i, i)
	fmt.Printf("类型: %T, 值: %v\n", f, f)
	fmt.Printf("类型: %T, 值: %v\n", s, s)
	fmt.Printf("类型: %T, 值: %v\n", b, b)
	fmt.Printf("类型: %T, 值: %v\n", arr, arr)
	fmt.Printf("类型: %T, 值: %v\n", slice, slice)
}
