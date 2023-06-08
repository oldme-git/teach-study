// 获取File的方式，底层都调用了 OpenFile
package os

import (
	"os"
	"testing"
)

var fileErr error

func TestOpen(t *testing.T) {
	_, fileErr = os.Open("./demo.txt")
	if fileErr != nil {
		t.Fatal(fileErr)
	}
}

// 创建a.txt
func TestCreate(t *testing.T) {
	_, fileErr = os.Create("./demo.txt")
	if fileErr != nil {
		t.Fatal(fileErr)
	}
}

// 会创建txt+随机数字的文件
func TestCreateTemp(t *testing.T) {
	_, fileErr = os.CreateTemp("", "oldme")
	if fileErr != nil {
		t.Fatal(fileErr)
	}
}

// 创建一个虚拟的文件对象，不会有真正的文件创建
func TestNewFile(t *testing.T) {
	_ = os.NewFile(1, "./b.txt")
}
