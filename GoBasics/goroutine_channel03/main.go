package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 写数据
func fn1(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("写入数据%v成功\n", i)
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
	wg.Done()
}

// 读数据
func fn2(ch chan int) {
	for v := range ch {
		fmt.Printf("读取数据%v成功\n", v)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done()
}
func main() {
	var ch = make(chan int, 10)
	wg.Add(1)
	go fn1(ch)
	wg.Add(1)
	go fn2(ch)
	wg.Wait()
	fmt.Println("成功退出")
}
