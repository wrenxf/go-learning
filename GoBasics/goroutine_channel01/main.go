package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 10
	a := <-ch
	fmt.Println(a)

}
