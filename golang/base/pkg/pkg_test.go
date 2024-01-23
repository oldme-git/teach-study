package pkg

import (
	"base/pkg/a"
	"base/pkg/b"
	"testing"
)

// 包全局变量的测试
func TestMy(t *testing.T) {
	b.GetInt()
	a.ChangeRegInt()
	b.GetInt()
}
