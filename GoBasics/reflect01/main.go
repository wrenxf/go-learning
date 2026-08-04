package main

import (
	"fmt"
	"reflect"
)

type myInt int
type Person struct {
	Name string
	Age  int
}

// 反射获取任意变量的类型
func reflectFn(x interface{}) {
	v := reflect.TypeOf(x)
	//fmt.Println(v)
	fmt.Printf("类型:%v 类型名称:%v 类型种类:%v\n", v, v.Name(), v.Kind())
}
func main() {
	a := 10
	b := 23.1
	c := true
	d := "你好Golang"
	var e myInt = 10
	var f Person = Person{
		Name: "张三",
		Age:  20,
	}
	var g = 25
	var h = [3]int{1, 2, 3}
	var i = []int{11, 22, 33}
	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(d)
	reflectFn(e)
	reflectFn(f)
	reflectFn(&g)
	reflectFn(h)
	reflectFn(i)
}
