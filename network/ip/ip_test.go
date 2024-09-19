package ip

import (
	"fmt"
	"testing"
)

func Test使用掩码判断是否属于同一网络(t *testing.T) {
	var (
		// 11111111 11111111 11111111 11000000
		subnetMask = 255<<24 | 255<<16 | 255<<8 | 192

		// 11111111 11111111 11111111 10000000
		subnetMask2 = 255<<24 | 255<<16 | 255<<8 | 128

		// 11000000 10101000 00000001 00000001
		ip1 = 192<<24 | 168<<16 | 1<<8 | 1

		// 11000000 10101000 00000001 00000011
		ip2 = 192<<24 | 168<<16 | 1<<8 | 89
	)

	fmt.Printf("掩码二进制1: %b，十位表示: %d\n", subnetMask, 计算掩码的十位表示法(subnetMask))
	fmt.Printf("掩码二进制2: %b，十位表示: %d\n", subnetMask2, 计算掩码的十位表示法(subnetMask2))

	var (
		ip1Network = ip1 & subnetMask
		ip2Network = ip2 & subnetMask
	)
	fmt.Printf("ip1和计算掩码1结果：%b\n", ip1Network)
	fmt.Printf("ip2和计算掩码1结果：%b\n", ip2Network)
	if ip1Network == ip2Network {
		fmt.Println("属于同一网络,计算掩码1")
	} else {
		fmt.Println("不属于同一网络,计算掩码1")
	}

	var (
		ip1Network2 = ip1 & subnetMask2
		ip2Network2 = ip2 & subnetMask2
	)
	fmt.Printf("ip1和计算掩码2结果：%b\n", ip1Network2)
	fmt.Printf("ip2和计算掩码2结果：%b\n", ip2Network2)
	if ip1Network2 == ip2Network2 {
		fmt.Println("属于同一网络,计算掩码2")
	} else {
		fmt.Println("不属于同一网络,计算掩码2")
	}
}

// 例如：255.255.255.192 等于 26
func 计算掩码的十位表示法(subnetMask int) int {
	var mask int
	for i := 0; i < 32; i++ {
		if subnetMask&(1<<uint(i)) != 0 {
			mask++
		}
	}
	return mask
}
