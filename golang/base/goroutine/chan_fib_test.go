// 使用通道协程完成斐波那契数列
package goroutine

import (
	"fmt"
	"testing"
	"time"
)

var quit = make(chan bool)

func fib(c chan int) {
	x, y := 1, 1

	for {
		// 协程阻塞
		fmt.Println("协程:等待")
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Println("协程:运算出斐波那契值", x)
		case <-quit:
			fmt.Println("Done calculating Fibonacci!")
			return
		}
	}
}

func TestEnter(t *testing.T) {
	start := time.Now()

	command := ""
	data := make(chan int)

	go fib(data)

	for {
		fmt.Println("主进程:接受协程通道中过来的斐波那契值")
		num := <-data
		fmt.Println("fib", num)
		// 主进程阻塞，等待用户输入
		fmt.Scanf("%s", &command)
		if command == "quit" {
			quit <- true
			break
		}
	}

	time.Sleep(1 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
