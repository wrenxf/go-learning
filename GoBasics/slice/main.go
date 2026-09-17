package main

import "fmt"

func main() {
	// 切片声明和初始化
	/*var slice1 []int                  // nil 切片
	var slice2 = []int{1, 2, 3, 4, 5} // 直接初始化
	//slice3 := make([]int, 5)          // 使用 make 创建
	slice4 := make([]int, 5, 10) // 指定长度和容量*/

	/*fmt.Println("slice1:", slice1, "len:", len(slice1), "cap:", cap(slice1))
	fmt.Println("slice2:", slice2, "len:", len(slice2), "cap:", cap(slice2))
	fmt.Println("slice3:", slice3, "len:", len(slice3), "cap:", cap(slice3))
	fmt.Println("slice4:", slice4, "len:", len(slice4), "cap:", cap(slice4))*/

	// 切片操作
	/*slice2 = append(slice2, 6, 7) // 追加元素
	fmt.Println("追加后:", slice2)

	subSlice := slice2[1:4] // 切片操作
	fmt.Println("子切片:", subSlice)*/

	// 从数组创建切片
	//arr := [5]int{1, 2, 3, 4, 5}
	//slice5 := arr[1:4]
	//fmt.Printf("从数组创建:%T %v", slice5, slice5)

	/*arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice3 := arr[2:7:8] // [low:high:max]，容量为 max-low
	fmt.Printf("arr[2:7:8] = %v (len=%d, cap=%d)\n",
		slice3, len(slice3), cap(slice3))*/

	slice := []int{1, 2, 3, 4, 5}

	// 截取切片
	slice1 := slice[1:4] // [2,3,4]
	slice2 := slice[2:]  // [3,4,5]
	slice3 := slice[:3]  // [1,2,3]

	fmt.Printf("原切片: %v\n", slice)
	fmt.Printf("slice[1:4]: %v\n", slice1)
	fmt.Printf("slice[2:]: %v\n", slice2)
	fmt.Printf("slice[:3]: %v\n", slice3)

	// 注意：截取的切片共享底层数组
	slice1[0] = 99
	fmt.Printf("修改 slice1[0] 后:\n")
	fmt.Printf("原切片: %v\n", slice) // 也被修改了

}
