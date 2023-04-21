package test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

var m2 = sync.Mutex{}

func p2(i int) {
	//m2.Lock()
	//defer m2.Unlock()
	fmt.Println(i)
}

func TestT2(t *testing.T) {
	for i := 0; i < 10; i++ {
		go p2(i)
	}
	runtime.Gosched()
}
