// regexp 包测试
package other

import (
	"fmt"
	"regexp"
	"testing"
)

// 字符串是否能匹配到正则表达式
func TestMatch(t *testing.T) {
	mStr, _ := regexp.MatchString("me", "oldme")
	fmt.Println(mStr)
	mByte, _ := regexp.Match("me", []byte("oldme"))
	fmt.Println(mByte)
}

// 可以用 Compile 来准备一个正则对象
func TestCompile(t *testing.T) {
	c, _ := regexp.Compile("me")
	fmt.Println(c.MatchString("oldme"))
}

// 匹配到字符串
func TestFindString(t *testing.T) {
	c, _ := regexp.Compile("me")
	s := c.FindString("oldme")
	// 字符串第一次出现的位置
	i := c.FindStringIndex("oldme")
	fmt.Println(s)
	fmt.Println(i)
}

// 匹配到所有的子字符串，即括号中的正则表达式
func TestStringSubmatch(t *testing.T) {
	c, _ := regexp.Compile("([a-z]+)\\d+")
	s := c.FindStringSubmatch("oldme123")
	fmt.Println(s)
}
