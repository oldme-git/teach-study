package _struct

import (
	"fmt"
	"strconv"
	"testing"
)

// 定义一个平台接口，包含一个支付方法
type Platform interface {
	Pay(amount int) error
	User
}

type User interface {
	Login()
	Logout()
}

type UserS struct {
}

func (u *UserS) Login()  {}
func (u *UserS) Logout() {}

// 微信平台
type Wechat struct {
	UserS
}

func (w *Wechat) Pay(amount int) error {
	fmt.Printf("wechat amount: %d\n", amount)
	return nil
}

func ExamplePlatform() {
	var (
		p Platform
		w = Wechat{}
	)
	p = &w
	p.Pay(2)

	// 类型断言
	_, ok := p.(User)
	fmt.Println(ok)

	// Output:
	// wechat amount: 2
	// true
}

func TestAb(t *testing.T) {
	a := ToInt(1)
	b := ToInt(2.1)
	c := ToInt("333")
	d := ToInt(true)
	e := ToInt(false)
	f := ToInt(map[string]int{"a": 1})
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

func ToInt(i any) int {
	switch v := i.(type) {
	case int:
		return v
	case float64:
		return int(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		vint, _ := strconv.Atoi(v)
		return vint
	}

	return 0
}
