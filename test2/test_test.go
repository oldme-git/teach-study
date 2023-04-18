package test2

import (
	"bytes"
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	buf := bytes.NewBufferString("hello")
	buf.WriteString(" oldme")
	fmt.Println(buf.String())
	b, _ := buf.ReadByte()
	fmt.Println(b)
	fmt.Println(buf.String())
}
