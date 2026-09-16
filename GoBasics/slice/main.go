package main

import "fmt"

func main() {
	// 切片声明和初始化
	var slice1 []int                  // nil 切片
	var slice2 = []int{1, 2, 3, 4, 5} // 直接初始化
	slice3 := make([]int, 5)          // 使用 make 创建
	slice4 := make([]int, 5, 10)      // 指定长度和容量

	fmt.Println("slice1:", slice1, "len:", len(slice1), "cap:", cap(slice1))
	fmt.Println("slice2:", slice2, "len:", len(slice2), "cap:", cap(slice2))
	fmt.Println("slice3:", slice3, "len:", len(slice3), "cap:", cap(slice3))
	fmt.Println("slice4:", slice4, "len:", len(slice4), "cap:", cap(slice4))

	// 切片操作
	slice2 = append(slice2, 6, 7) // 追加元素
	fmt.Println("追加后:", slice2)

	subSlice := slice2[1:4] // 切片操作
	fmt.Println("子切片:", subSlice)

	// 从数组创建切片
	arr := [5]int{1, 2, 3, 4, 5}
	slice5 := arr[1:4]
	fmt.Printf("从数组创建:%T %v", slice5, slice5)
}
