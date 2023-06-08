package os

import (
	"fmt"
	"os"
	"testing"
)

var f = getFile()
var d = getDir()

func getFile() *os.File {
	f, err := os.Open("./demo.txt")
	if err != nil {
		panic(err)
	}
	return f
}

func getDir() *os.File {
	f, err := os.Open("./")
	if err != nil {
		panic(err)
	}
	return f
}

func TestFileIsDir(t *testing.T) {
	d, _ := os.Open("./")
	var b = make([]byte, 2)
	_, err := d.Read(b)
	if err != nil {
		t.Fatal(err)
	}
}

// Read 会改变文件的偏移量
func TestRead(t *testing.T) {
	var b = make([]byte, 2)
	i, err := f.Read(b)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("第一次读取字节数：%d\n", i)
	t.Logf("读取的值：%s\n", string(b))
	i, err = f.Read(b)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("第二次读取字节数：%d\n", i)
	t.Logf("读取的值：%s\n", string(b))
}

// ReadAt 从指定位置开始读，不会改变文件的偏移量
func TestReadAt(t *testing.T) {
	f, _ := os.Open("./demo.txt")
	var b = make([]byte, 2)
	i, err := f.ReadAt(b, 2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("第一次读取字节数：%d\n", i)
	t.Logf("读取的值：%s\n", string(b))
	i, err = f.Read(b)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("第二次读取字节数：%d\n", i)
	t.Logf("读取的值：%s\n", string(b))
}

// 如果文件是目录
func TestReadDir(t *testing.T) {
	dir, err := d.ReadDir(0)
	if err != nil {
		return
	}
	for _, v := range dir {
		fmt.Println(v.Type())
	}
}

// 如果文件是目录
func TestReaddir(t *testing.T) {
	d, err := os.Open("./")
	dir, err := d.Readdir(0)
	if err != nil {
		return
	}
	for _, v := range dir {
		fmt.Printf("文件名: %s, isDir: %t\n", v.Name(), v.IsDir())
	}
}

func TestDirnames(t *testing.T) {
	name, err := d.Readdirnames(0)
	if err != nil {
		return
	}
	for _, v := range name {
		fmt.Println(v)
	}
}
