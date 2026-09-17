package main

import "fmt"

func makeAdder(increment int) func(int) int {
	return func(x int) int {
		return x + increment
	}
}
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
func main() {
	add5 := makeAdder(5)
	add10 := makeAdder(10)

	fmt.Printf("add5(20)=%d\n", add5(20))
	fmt.Printf("add10(20)=%d\n", add10(20))

	c1 := counter()
	c2 := counter()

	fmt.Printf("\n计数器1: %d\n", c1()) // 1
	fmt.Printf("计数器1: %d\n", c1())   // 2
	fmt.Printf("计数器2: %d\n", c2())   // 1
	fmt.Printf("计数器1: %d\n", c1())   // 3
	fmt.Printf("计数器2: %d\n", c2())   // 2
	// 闭包在实际应用中
	message := "Hello"

	// 延迟执行闭包
	delayedFunc := func() {
		fmt.Println("延迟执行:", message)
	}

	message = "World"
	delayedFunc() // 输出: 延迟执行: World

}
