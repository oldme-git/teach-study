// 解释内存对齐

package mem_alignment

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestSizeOfAndAlignOf(t *testing.T) {
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
		a     [3]string
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

// 对于数组类型而言，它的大小是元素数量 * 元素类型字节数
func TestArrayAndSlice(t *testing.T) {
	var (
		slice []int8
		array [3]int8
	)
	fmt.Printf("切片：%v\n", unsafe.Sizeof(slice))
	fmt.Printf("数组：%v\n", unsafe.Sizeof(array))
}

type MemStruct struct {
	b   bool   // alignof: 1
	i8  int8   // alignof: 1
	i32 int32  // alignof: 4
	s   string // alignof: 8
}

func TestStruct(t *testing.T) {
	fmt.Printf("MemStruct的对齐系数：{%v}, 等于string的对齐系数：{%v}", unsafe.Alignof(MemStruct{}), unsafe.Alignof(string("1")))
}

func TestArray(t *testing.T) {
	var (
		it  interface{}
		arr [3]interface{}
	)
	fmt.Printf("数组interface{}的对齐系数：{%v}, 等于interface的对齐系数：{%v}", unsafe.Alignof(arr), unsafe.Alignof(it))
}

func TestAlignOf(t *testing.T) {
	var (
		i8    int8
		s     string
		array [20]string
	)
	fmt.Printf("string sizeof:%v, alignof: %v\n", unsafe.Sizeof(s), unsafe.Alignof(s))        // min(8, 1)
	fmt.Printf("int8 sizeof:%v, alignof: %v\n", unsafe.Sizeof(i8), unsafe.Alignof(i8))        // min(8, 16)
	fmt.Printf("array sizeof:%v, alignof: %v\n", unsafe.Sizeof(array), unsafe.Alignof(array)) // min(8, 16)
}

type emptyStruct struct{}

type S1 struct {
	empty emptyStruct
	i8    int8
}

type S2 struct {
	i8    int8
	empty emptyStruct
}

type S3 struct {
	i16   int16
	empty emptyStruct
}

type S4 struct {
	i16   int16
	i8    int8
	empty emptyStruct
}

func TestSpaceStructMem(t *testing.T) {
	fmt.Printf("S1的占用: %v\n", unsafe.Sizeof(S1{}))
	fmt.Printf("S2的占用: %v\n", unsafe.Sizeof(S2{}))
	fmt.Printf("S3的占用: %v\n", unsafe.Sizeof(S3{}))
	fmt.Printf("S4的占用: %v\n", unsafe.Sizeof(S4{}))
	// S3 空结构从第二位开始，往后补充两个字节
	fmt.Printf("S3的空结构体偏移量: %v\n", unsafe.Offsetof(S3{}.empty))
	// S4 空结构从第三位开始，往后补充一个字节
	fmt.Printf("S4的空结构体偏移量: %v\n", unsafe.Offsetof(S4{}.empty))
}
