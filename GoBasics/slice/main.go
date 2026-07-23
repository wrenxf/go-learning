package main

func main() {
	/*切片的声明和初始化*/
	/*var arr1 []int
	fmt.Printf("%v %T %v\n", arr1, arr1, len(arr1))

	var arr2 = []int{1, 2, 3, 4}
	fmt.Printf("%v %T %v\n", arr2, arr2, len(arr2))

	var arr3 = []int{1: 2, 3: 4, 5: 6}
	fmt.Printf("%v %T %v", arr3, arr3, len(arr3))*/

	/*切片的循环遍历*/
	/*var strSlice = []string{"php", "java", "nodejs", "golang"}
	for i := 0; i < len(strSlice); i++ {
		fmt.Println(strSlice[i])
	}*/

	/*var strSlice = []string{"php", "java", "nodejs", "golang"}
	for k, v := range strSlice {
		fmt.Println(k, v)
	}*/

	/*基于数组定义切片*/
	/*a := [5]int{55, 56, 57, 58, 59}
	b := a[:] //获取数组里的所有值
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)

	c := a[1:4] //左闭右开
	fmt.Printf("%v %T\n", c, c)
	d := a[2:]
	fmt.Printf("%v-%T\n", d, d) //[57, 58, 59]
	e := a[:3]                  //表示获取第三个下标前面的数据
	fmt.Printf("%v-%T\n", e, e) //[55,56,57]*/

	/*基于切片再切片*/
}
