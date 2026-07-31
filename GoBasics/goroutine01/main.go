package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() hello GoLang-", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}
func main() {
	wg.Add(1)
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main() hello GoLang-", i)
		time.Sleep(time.Millisecond * 20)
	}

	//解决主线程结束导致协程还没运行完就被迫结束

	//第一种(不推荐),主线程延时一秒结束
	//time.Sleep(time.Second)

	//第二种
	wg.Wait() //等待协程执行完毕

	fmt.Println("主线程退出...")
}
