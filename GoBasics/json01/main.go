package main

import (
	"fmt"
)

/*结构体要转为json,结构体中的字段必须是共有的(首字母大写)*/
/*type Student struct {
	ID     int
	Gender string
	Name   string //私有属性不能被json访问
	Son    string
}*/
/*结构体标签 Tag*/
type Student struct {
	ID     int    `json01:"id"`
	Gender string `json01:"gender"`
	Name   string `json01:"name"`
	Son    string `json01:"son"`
}

func main() {
	var s1 = Student{
		ID:     12,
		Gender: "男",
		Name:   "李四",
		Son:    "s0001",
	}
	fmt.Printf("%#v\n", s1) //%v打印数据，%#v打印详细数据
	/*结构体对象转化成 Json 字符串*/
	/*jsonByte, _ := json01.Marshal(s1)
	jsonStr := string(jsonByte)
	fmt.Printf("%v\n", jsonStr)*/

	/*Json字符串转换成结G构体对象*/
	/*var str = `{"ID":1,"Gender":"男","Name":"李四","Sno":"s0001"}`
	var s2 Student
	err := json01.Unmarshal([]byte(str), &s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v %v\n", s2, s2.Name)*/

}
