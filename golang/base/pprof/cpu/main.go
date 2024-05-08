package main

import (
	"math"
	"math/rand"
	"os"
	"runtime/pprof"
)

func main() {
	// 保存 CPU 采样数据
	file := "cpu.pprof"
	os.Remove(file)
	f, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 开始采样
	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	// 测试程序
	for i := 0; i < 1000; i++ {
		nums := genRandomNumbers(10000)
		for _, v := range nums {
			_ = math.Pow(float64(v), rand.Float64())
		}
	}
}

// 测试程序，生成一个随机数切片
func genRandomNumbers(n int) []int {
	nums := make([]int, n)
	for i := 1; i < n; i++ {
		nums[i] = rand.Int() * n
	}
	return nums
}
