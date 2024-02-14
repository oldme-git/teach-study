// 结构体实现一个接口，并且此接口做为结构体的一个字段

package _struct

import (
	"fmt"
	"testing"
)

type Wrap struct {
	WInterface
}

type WInterface interface {
	Print(s string)
}

type AI interface {
	WInterface
	A()
}

type Wrap2 struct {
}

func (w2 Wrap2) Print(s string) {
	fmt.Println("w2" + s)
}

type Wrap3 struct {
}

func (w2 *Wrap3) Print(s string) {
	fmt.Println("w2" + s)
}

func (w2 *Wrap3) A() {
}

func TestA(t *testing.T) {
	var w2 = Wrap2{}
	var _ = Wrap{&w2}

	var a AI
	a = &Wrap3{}
	_, ok := a.(WInterface)
	fmt.Println(ok)
}
