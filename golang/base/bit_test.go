// 位运算操作
package test

import (
	"fmt"
	"testing"
)

func TestBase(t *testing.T) {
	// 与运算，两个位都为1时，结果才为1
	fmt.Printf("%b\n", 0b110&0b100) // 100
	// 或运算，两个位都为0时，结果才为0
	fmt.Printf("%b\n", 0b110|0b100) // 110
	// 异或运算，两个位相同为0，相异为1
	fmt.Printf("%b\n", 0b110^0b100) // 010
	// 左移运算，全部左移若干位，高位丢弃，低位补0
	fmt.Printf("%b\n", 0b110<<1) // 1100
	// 右移运算，全部右移若干位，高位丢弃，低位补0
	fmt.Printf("%b\n", 0b110>>1) // 11
}

// m左移n位相当于m乘以2的n次方，m右移n位相当于m除以2的n次方
func TestCompute(t *testing.T) {
	fmt.Printf("%d\n", 1<<16) // 1 * 2^16 65536
	fmt.Printf("%d\n", 2<<16) // 2 * 2^16 131072
	fmt.Println(65536 >> 8)   // 65536 / 2^8 256
}

// 判断奇偶
func TestIsEven(t *testing.T) {
	fmt.Println(10 & 1) // 最后一位为0 偶数
	fmt.Println(11 & 1) // 最后一位为1 奇数
}

// 异或特性
func TestEq(t *testing.T) {
	i := 10
	fmt.Println(i ^ i) // 0
	fmt.Println(i ^ 0) // i
}
