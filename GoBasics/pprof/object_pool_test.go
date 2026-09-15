package main

import (
	"sync"
	"testing"
)

var (
	result = make([]int, 10000)
	pool   = sync.Pool{
		New: func() any {
			return make([]int, 10000)
		},
	}
)

func recall() []int {
	result := []int{}
	for i := 0; i < 10000; i++ {
		result = append(result, i)
	}
	result[len(result)-1] = -1
	return result
}
func recall2() []int {
	result := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		result[i] = i
	}
	result[len(result)-1] = -1
	return result
}
func recall3() []int {
	for i := 0; i < 10000; i++ {
		result[i] = i
	}
	result[len(result)-1] = -1
	return result
}
func recall4() []int {
	result := pool.Get().([]int)
	for i := 0; i < 10000; i++ {
		result[i] = i
	}
	result[len(result)-1] = -1
	return result
}
func Serach() []int {
	result := []int{}
	recallResult := recall4()
	for _, ele := range recallResult {
		if ele == -1 {
			break
		}
		if ele%1000 == 0 {
			result = append(result, ele)
		}
	}
	pool.Put(recallResult)
	return result
}
func BenchmarkSearch(b *testing.B) {
	for b.Loop() {
		Serach()
	}
}

//go test ./micro_service/pprof -bench=^BenchmarkSearch$ -cpuprofile=data/cpu1
// -memprofile=data/mem1 -run=^$ -count=1 -benchmem -benchtime=10s
