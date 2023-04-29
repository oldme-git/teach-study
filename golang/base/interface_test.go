// 接口实现
package test

import (
	"fmt"
	"testing"
)

type A interface {
	Query() string
}

type B struct {
}

func (b *B) Query() string {
	return "Query"
}

func (b *B) Query2() string {
	return "Query2"
}

func TestInterface(t *testing.T) {
	var a A
	b := new(B)
	a = b
	fmt.Println(a.Query())
	fmt.Println(b.Query2())
}
