package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 测试程序
		for i := 0; i < 1000; i++ {
			nums := genRandomNumbers(10000)
			for _, v := range nums {
				_ = math.Pow(float64(v), rand.Float64())
			}
		}
		fmt.Fprint(w, "Hello, world!")
	})
	http.ListenAndServe(":8080", nil)
}

// 测试程序，生成一个随机数切片
func genRandomNumbers(n int) []int {
	nums := make([]int, n)
	for i := 1; i < n; i++ {
		nums[i] = rand.Int() * n
	}
	return nums
}
