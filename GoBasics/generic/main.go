package main

import "fmt"

func getData[T any](value T) T {
	return value
}
func main() {
	str1 := "hello"
	fmt.Println(getData(str1))
	str2 := "GoLang"
	fmt.Println(len(str2))
}
