package main

import (
	"fmt"
	"hash/fnv"
)

type BloomFilter struct {
	bitArray []bool
	hashList []BloomFilterHash
}

// BloomFilterHash 布隆过滤器的hash
type BloomFilterHash func(s string) uint32

func NewBloomFilter(size uint32) *BloomFilter {
	return &BloomFilter{
		bitArray: make([]bool, size),
	}
}

// RegisterHash 注册hash函数
func (b *BloomFilter) RegisterHash(h ...BloomFilterHash) {
	b.hashList = append(b.hashList, h...)
}

// Add 添加一个元素
func (b *BloomFilter) Add(s string) {
	bitLen := uint32(len(b.bitArray))
	for _, v := range b.hashList {
		index := v(s) % bitLen
		b.bitArray[index] = true
	}
}

// Exist 判断一个元素是否存在
func (b *BloomFilter) Exist(s string) bool {
	var (
		bitLen         = uint32(len(b.bitArray))
		hashLen        = uint32(len(b.hashList))
		count   uint32 = 0
	)

	for _, v := range b.hashList {
		index := v(s) % bitLen
		if b.bitArray[index] {
			count++
		}
	}

	if hashLen == count {
		return true
	}
	return false
}

// 使用New32a实现第一个hash
func hash1(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// 使用New32实现第一个hash
func hash2(s string) uint32 {
	h := fnv.New32()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {
	testStr := "tyyn1022@gmail.com"
	bloom := NewBloomFilter(1 << 16)
	bloom.RegisterHash(hash1, hash2)
	// 添加元素
	bloom.Add(testStr)
	// 检测元素是否存在
	fmt.Println(bloom.Exist(testStr))
	fmt.Println(bloom.Exist(testStr + "not exist"))
}
