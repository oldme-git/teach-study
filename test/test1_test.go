package test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

func TestT1(t *testing.T) {
	lock := &sync.Mutex{}
	for i := 0; i < 1000; i++ {
		go Count(lock)
	}
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 1000 {
			break
		}
	}
}
