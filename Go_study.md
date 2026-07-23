# Go语言关键字

- Go语言中一共有25个关键字

| 1           | 2          | 3              | 4          | 5             | 6           | 7            | 8          |
| ----------- | ---------- | -------------- | ---------- | ------------- | ----------- | ------------ | ---------- |
| ***if***    | ***else*** | ***switch***   | ***case*** | ***default*** | ***break*** | ***return*** | ***goto*** |
| fallthrough | ***for***  | ***continue*** | type       | ***struct***  | var         | ***const***  | map        |
| func        | interface  | range          | import     | package       | defer       | go           | select     |
| chan        |            |                |            |               |             |              |            |

# 1.Go语言定义变量

## 1.var定义变量

```go
var 变量名 类型=表达式
```

```go
var name string="zhangsan"
```

## 2.类型推导方式定义变量

a在函数内部，可以使用更简略的:=方式声明并初始化变量。 

**注意：短变量只能用于声明局部变量，不能用于全局变量的声明**

```go
变量名 := 表达式
```

```go
n := 10
```

# 切片

切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。 它非常灵活，支持自动扩容。 

切片是一个**引用类型**，它的内部结构包含**地址**、**长度**和**容量**。

 声明切片类型的基本语法如下： 

```go
var name []T 
```

其中： 

1. name:表示变量名 
2. T:表示切片中的元素类型

## 关于 nil 的认识

当你声明了一个变量 , 但却还并没有赋值时 , golang 中会自动给你的变量赋值一个默认零
值。这是每种类型对应的零值。

```go
bool -> false
numbers -> 0
string-> "" 
pointers -> nil
slices -> nil
maps -> nil
channels -> nil
functions -> nil
interfaces -> nil
```

## 切片的循环遍历

切片的循环遍历和数组的循环遍历是一样的

```go
var a = []string{"北京", "上海", "深圳"}
// 方法 1：for 循环遍历
for i := 0; i < len(a); i++ {
fmt.Println(a[i])
}
// 方法 2：for range 遍历
for index, value := range a {
fmt.Println(index, value)
}
```

## 基于数组定义切片

由于切片的底层就是一个数组，所以我们可以基于数组定义切片。

```go
func main() {
// 基于数组定义切片
a := [5]int{55, 56, 57, 58, 59}
b := a[1:4] //基于数组 a 创建切片，包括元素 a[1],a[2],a[3]
fmt.Println(b) //[56 57 58]
fmt.Printf("type of b:%T\n", b) //type of b:[]int
}
还支持如下方式：
c := a[1:] //[56 57 58 59]
d := a[:4] //[55 56 57 58]
e := a[:] //[55 56 57 58 59]
```

## 切片再切片

除了基于数组得到切片，我们还可以通过切片来得到切片。

```go
func main() {
//切片再切片
a := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
fmt.Printf("a:%v type:%T len:%d cap:%d\n", a, a, len(a), cap(a))
b := a[1:3]
fmt.Printf("b:%v type:%T len:%d cap:%d\n", b, b, len(b), cap(b))
c := b[1:5]
fmt.Printf("c:%v type:%T len:%d cap:%d\n", c, c, len(c), cap(c))
}

```

```go
输出：
a:[北京 上海 广州 深圳 成都 重庆] type:[6]string len:6 cap:6
b:[上海 广州] type:[]string len:2 cap:5
c:[广州 深圳 成都 重庆] type:[]string len:4 cap:4
```

**注意： 对切片进行再切片时，索引不能超过原数组的长度，否则会出现索引越界的错误。**

## 关于切片的长度和容量

切片拥有自己的长度和容量，我们可以通过使用内置的 **len()函数**求长度，使用内置的 **cap()函数**求切片的容量。
切片的长度就是它所包含的元素个数。
切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。

```go
s := []int{2, 3, 5, 7, 11, 13}
fmt.Println(s)
fmt.Printf("长度:%v 容量 %v\n", len(s), cap(s))
c := s[:2]
fmt.Println(c)
fmt.Printf("长度:%v 容量 %v\n", len(c), cap(c))
d := s[1:3]
fmt.Println(d)
fmt.Printf("长度:%v 容量 %v", len(d), cap(d))

```

```go
输出：
D:\golang\src\demo01>go run main.go
[2 3 5 7 11 13]
长度:6 容量 6
[2 3]
长度:2 容量 6
[3 5]
长度:2 容量 5
```

