package main

import (
	"fmt"
	"sync"
	"time"
)

//需求：要统计 1-120000 的数字中那些是素数？
/*
1协程统计 1-30000
2 协程统计30001-60000
3协程统计60001-90000
4协程统计90001-120000
*/

var wg sync.WaitGroup

func test(n int) {
	defer wg.Done()
	for num := (n-1)*30000 + 1; num < n*30000; num++ {
		if num > 1 {
			var flag = true
			for i := 2; i < num; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				//fmt.Println(num, "是素数")
			}
		}

	}
}

/*func main() {
	start := time.Now()
	for num := 1; num <= 120000; num++ {
		flag := true //假设是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明该 num 不是素数
				flag = false
				break
			}
		}
		if flag {
			// fmt.Println(num)
		}
	}
	end := time.Now().Sub(start)
	fmt.Println("普通的方法耗时=", end)
}*/

func main() {
	start := time.Now()
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	end := time.Now().Sub(start)
	fmt.Println(end)
}
