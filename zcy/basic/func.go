package main

import "fmt"

func main() {
	c, d := 3, 5
	arg1(c, d)
}
func arg1(a int, b int) {
	a = a + b
	return
	fmt.Println("我不会被输出")
}

func return1(a, b int) int {
	a = a + b
	c := a
	return c
}
func return2(a, b int) (c int) {
	a = a + b
	c = a
	return
}
