// 一个结构体和接口互相嵌套的代码

package test

import (
	"fmt"
	"testing"
)

// 一个接口
type MyInterface interface {
	InternalStr() string
	SetInternalStruct(*InternalStruct)
	GetInternalStruct() *InternalStruct
}

type InternalStruct struct {
	MyInterface MyInterface
	str         string
}

func (i *InternalStruct) InternalStr() string {
	return i.str
}

func (i *InternalStruct) SetInternalStr(s string) {
	i.str = s
}

func (i *InternalStruct) GetInternalStruct() *InternalStruct {
	return i
}

// 嵌套InternalStruct以实现MyInterface
type ExternalStruct struct {
	*InternalStruct
}

func (e *ExternalStruct) SetInternalStruct(i *InternalStruct) {
	e.InternalStruct = i
}

func NewExternalStruct(i *InternalStruct) MyInterface {
	return &ExternalStruct{
		InternalStruct: i,
	}
}

func TestOk(t *testing.T) {
	// 创建InternalStruct
	internal := &InternalStruct{}
	// 创建ExternalStruct
	external := NewExternalStruct(internal)
	// InternalStruct赋值str
	internal.SetInternalStr("ok")
	// InternalStruct赋值MyInterface
	internal.MyInterface = external
	// 重新赋予InternalStruct到external
	external.SetInternalStruct(internal)
	// 打印ok
	ok := external.InternalStr()
	fmt.Println(ok)
	// 互相调用
	external.GetInternalStruct().MyInterface.GetInternalStruct().SetInternalStr("not")
	// 再次获取值，打印not
	ok = external.InternalStr()
	fmt.Println(ok)
}
