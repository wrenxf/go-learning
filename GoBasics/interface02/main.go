package main

import "fmt"

type A interface{}

func main() {
	var a A
	var str = "你好"
	a = str
	fmt.Printf("%v %T\n", a, a)

	var num = 20
	a = num
	fmt.Printf("%v %T", a, a)
}
