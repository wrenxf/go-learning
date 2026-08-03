package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 向 intChan 放入 1-120000 个数
func putNum(intChan chan int) {
	for i := 1; i <= 1000; i++ {
		intChan <- i
	}
	//关闭 intChan
	close(intChan)
	wg.Done()
}

// 从 intChan 取出数据，并判断是否为素数,如果是，就放入到 primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for num := range intChan {
		var flag bool = true
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明该 num 不是素数
				flag = false
				break
			}
		}
		if flag {
			//将这个数就放入到 primeChan
			primeChan <- num
		}
	}
	//判断关闭
	exitChan <- true
	wg.Done()
}

// 打印素数的方法
func printPrime(primeChan chan int) {
	for v := range primeChan {
		fmt.Println(v)
	}
	wg.Done()
}
func main() {
	start := time.Now()
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000) //放入结果
	//标识退出的管道
	exitChan := make(chan bool, 8) // 8 个
	//开启一个协程，向 intChan 放入 1-8000 个数
	wg.Add(1)
	go putNum(intChan)
	//开启 4 个协程，从 intChan 取出数据，并判断是否为素数,如果是，就放入到 primeChan
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}
	//打印素数
	wg.Add(1)
	go printPrime(primeChan)
	//判断什么时候退出
	wg.Add(1)
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		//当我们从 exitChan 取出了 8 个结果，就可以放心的关闭 prprimeChan
		close(primeChan)
		wg.Done()
	}()
	wg.Wait()
	end := time.Now().Sub(start)
	fmt.Println(end)
	fmt.Println("main 线程退出")
}
