package main

func main() {
	var ch1 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch1 <- i
	}
	close(ch1)
	for v := range ch1 {
		println(v)
	}
}
