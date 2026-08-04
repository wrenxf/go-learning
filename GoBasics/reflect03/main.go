package main

import (
	"fmt"
	"reflect"
)

func PrintStructField(s interface{}) {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}
}
func main() {

}
