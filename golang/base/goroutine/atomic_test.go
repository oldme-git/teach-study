package goroutine

import (
	"fmt"
	"sync/atomic"
	"testing"
)

// atomic value
func TestAtomicValue(t *testing.T) {
	var v atomic.Value
	v.Store(1)
	fmt.Println(v.Load()) // 1
	v.Swap(2)
	fmt.Println(v.Load()) // 2
	b := v.CompareAndSwap(2, 3)
	fmt.Println(v.Load()) // 3
	fmt.Println(b)
}
