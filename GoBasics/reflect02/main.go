package main

import (
	"fmt"
	"reflect"
)

// 反射获取任意变量的类型
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	kind := v.Kind()

	switch kind {
	case reflect.Int64:
		fmt.Printf("int类型的原始值%v 计算后的值%v\n", v.Int(), v.Int()+10)
	case reflect.Float32:
		fmt.Printf("float32类型的原始值%v\n", v.Float())
	case reflect.Float64:
		fmt.Printf("float64类型的原始值%v\n", v.Float())
	case reflect.String:
		fmt.Printf("string类型的原始值%v\n", v.String())
	default:
		fmt.Printf("还没有判断该类型\n")

	}
}
func main() {
	var a float32 = 3.14
	var b int64 = 100
	var c string = "hello GoLang"
	reflectValue(a)
	reflectValue(b)
	reflectValue(c)
}
