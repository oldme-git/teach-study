package os

import (
	"fmt"
	"os"
	"testing"
)

var fw = getFileCreate()

func getFileCreate() *os.File {
	f, err := os.Create("./demo")
	if err != nil {
		panic(err)
	}
	return f
}

func TestWrite(t *testing.T) {
	w, err := fw.Write([]byte("012"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("第一次写入了%d个字节\n", w)

	w, err = fw.Write([]byte("34"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("第二次写入了%d个字节\n", w)
}

func TestWriteAt(t *testing.T) {
	w, err := fw.Write([]byte("01abc"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("第一次写入了%d个字节\n", w)

	w, err = fw.WriteAt([]byte("234"), 2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("第二次写入了%d个字节\n", w)
}

func TestWriteString(t *testing.T) {
	w, err := fw.WriteString("012")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("第一次写入了%d个字节\n", w)

	w, err = fw.WriteString("34")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("第二次写入了%d个字节\n", w)
}

// 追加写入
func TestWriteAppend(t *testing.T) {
	fw, _ := os.OpenFile("./demo.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0)
	_, err := fw.WriteString("abc")
	if err != nil {
		t.Fatal(err)
	}
}
