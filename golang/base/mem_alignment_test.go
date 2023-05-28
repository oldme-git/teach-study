// 解释内存对齐

package test

import (
	"fmt"
	"testing"
	"unsafe"
)

type MemStruct1 struct {
	a int8
	b int16
	c complex128
}

type MemStruct2 struct {
	a int8
	c int32
	b int16
}

func TestMemBase(t *testing.T) {
	var m1 MemStruct1
	t.Log(unsafe.Alignof(m1))
	t.Log(unsafe.Sizeof(m1))
	t.Log(unsafe.Alignof(float32(1)))
	t.Log(unsafe.Sizeof(float32(1)))

	var m2 MemStruct2
	t.Log(unsafe.Alignof(m2))
	t.Log(unsafe.Sizeof(m2))
}

func TestMemBase2(t *testing.T) {
	var b int8
	t.Log(unsafe.Alignof(b))
	t.Log(unsafe.Sizeof(b))
}

func TestMemBase3(t *testing.T) {
	t.Log(213)
}

func TestAlignOf(t *testing.T) {
	var (
		b     bool
		s     string
		i     int
		i8    int8
		i16   int16
		i32   int32
		i64   int64
		ui    uint
		ui8   uint8
		ui16  uint16
		ui32  uint32
		ui64  uint64
		uip   uintptr
		f32   float32
		f64   float64
		c64   complex64
		c128  complex128
		in    interface{}
		a     [3]int8
		slice []string
	)

	fmt.Printf("bool sizeof:%v, alignof: %v\n", unsafe.Sizeof(b), unsafe.Alignof(b))
	fmt.Printf("string sizeof:%v, alignof: %v\n", unsafe.Sizeof(s), unsafe.Alignof(s))
	fmt.Printf("int sizeof:%v, alignof: %v\n", unsafe.Sizeof(i), unsafe.Alignof(i))
	fmt.Printf("int8 sizeof:%v, alignof: %v\n", unsafe.Sizeof(i8), unsafe.Alignof(i8))
	fmt.Printf("int16 sizeof:%v, alignof: %v\n", unsafe.Sizeof(i16), unsafe.Alignof(i16))
	fmt.Printf("int32 sizeof:%v, alignof: %v\n", unsafe.Sizeof(i32), unsafe.Alignof(i32))
	fmt.Printf("int64 sizeof:%v, alignof: %v\n", unsafe.Sizeof(i64), unsafe.Alignof(i64))
	fmt.Printf("uint sizeof:%v, alignof: %v\n", unsafe.Sizeof(ui), unsafe.Alignof(ui))
	fmt.Printf("uint8 sizeof:%v, alignof: %v\n", unsafe.Sizeof(ui8), unsafe.Alignof(ui8))
	fmt.Printf("uint16 sizeof:%v, alignof: %v\n", unsafe.Sizeof(ui16), unsafe.Alignof(ui16))
	fmt.Printf("uint32 sizeof:%v, alignof: %v\n", unsafe.Sizeof(ui32), unsafe.Alignof(ui32))
	fmt.Printf("uint64 sizeof:%v, alignof: %v\n", unsafe.Sizeof(ui64), unsafe.Alignof(ui64))
	fmt.Printf("uintptr sizeof:%v, alignof: %v\n", unsafe.Sizeof(uip), unsafe.Alignof(uip))
	fmt.Printf("float32 sizeof:%v, alignof: %v\n", unsafe.Sizeof(f32), unsafe.Alignof(f32))
	fmt.Printf("float64 sizeof:%v, alignof: %v\n", unsafe.Sizeof(f64), unsafe.Alignof(f64))
	fmt.Printf("complex64 sizeof:%v, alignof: %v\n", unsafe.Sizeof(c64), unsafe.Alignof(c64))
	fmt.Printf("complex128 sizeof:%v, alignof: %v\n", unsafe.Sizeof(c128), unsafe.Alignof(c128))
	fmt.Printf("interface sizeof:%v, alignof: %v\n", unsafe.Sizeof(in), unsafe.Alignof(in))
	fmt.Printf("array sizeof:%v, alignof: %v\n", unsafe.Sizeof(a), unsafe.Alignof(a))
	fmt.Printf("slice sizeof:%v, alignof: %v\n", unsafe.Sizeof(slice), unsafe.Alignof(slice))
}

func TestAbc(t *testing.T) {
	var s string
	fmt.Println(unsafe.Sizeof(s))
}
