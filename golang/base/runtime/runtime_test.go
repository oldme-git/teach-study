// runtime包相关测试
package runtime

import (
	"fmt"
	"runtime"
	"testing"
)

func TestNumCPU(t *testing.T) {
	fmt.Println(runtime.NumCPU())
}
