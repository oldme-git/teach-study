package test2

import (
	"fmt"
	"testing"
	"time"
)

type Aint int

type A interface {
	Table(int) int
}

func (ai *Aint) Table(a int) int {
	return int(*ai) + 2
}

func TestA(t *testing.T) {
	go func() {
		<-time.After(3 * time.Millisecond)
		fmt.Println("处理任务")
	}()
	fmt.Println(time.Duration(1) * time.Second)
	time.Sleep(5 * time.Millisecond)
	fmt.Println("main")
	time.Sleep(5 * time.Second)
	fmt.Println("main end")
}
