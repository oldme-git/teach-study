// runtime包相关测试
package test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestNumCPU(t *testing.T) {
	fmt.Println(runtime.NumCPU())
}
