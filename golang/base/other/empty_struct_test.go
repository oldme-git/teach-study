package other

import (
	"encoding/json"
	"fmt"
	"testing"
)

type User struct {
	Pwd  string `json:"pwd"`
	Pwd2 string `json:"pwd2"`
	Age  int    `json:"age"`
}

type UserInner struct {
	Pwd2 struct{} `json:"pwd2,omitempty"`
}

type UserOut struct {
	User
	Pwd struct{} `json:"pwd,omitempty"`
	UserInner
}

// TestA和TestAPtr，这是一个很有趣的问题
func TestA(t *testing.T) {
	user := UserOut{User: User{
		Pwd:  "123",
		Pwd2: "abc",
		Age:  11,
	}}
	res, _ := json.MarshalIndent(user, "", "\t")
	fmt.Println(string(res))
}

type UserOutPtr struct {
	User
	// 用于掩盖的类型不一定要是空结构体指针，能被omitempty忽略的都可以
	Pwd *bool `json:"pwd,omitempty"`
	UserInner
}

func TestAPtr(t *testing.T) {
	user := UserOutPtr{User: User{
		Pwd:  "123",
		Pwd2: "abc",
		Age:  11,
	}}
	res, _ := json.MarshalIndent(user, "", "\t")
	fmt.Println(string(res))
}

// 解释这个问题1：结构体嵌套，相同字段的情况下，层级越浅，优先级越高
type User2 struct {
	Pwd string `json:"pwd"`
	UserInner2
}

type UserInner2 struct {
	Pwd string `json:"pwd"`
	UserInner3
}

type UserInner3 struct {
	Pwd string `json:"pwd"`
}

func TestB(t *testing.T) {
	user := User2{
		Pwd: "1",
		UserInner2: UserInner2{
			Pwd: "2",
			UserInner3: UserInner3{
				Pwd: "3",
			},
		},
	}
	fmt.Println(user.Pwd)
	fmt.Println(user.UserInner2.Pwd)
}

// 解释这个问题2：多个匿名结构体字段在同级别下冲突会不能正确引用，导致JSON无法正常转换
type UserInner3A struct {
	Name string `json:"name"`
}

type UserInner3B struct {
	Name string `json:"name"`
}

type User3 struct {
	UserInner3A
	UserInner3B
}

func TestC(t *testing.T) {
	user3 := User3{
		UserInner3A: UserInner3A{
			Name: "3a",
		},
		UserInner3B: UserInner3B{
			Name: "3b",
		},
	}
	res, _ := json.MarshalIndent(user3, "", "\t")
	fmt.Println(string(res))
}

func TestD(t *testing.T) {
	fmt.Printf("%p\n", &struct{}{})
	fmt.Printf("%p\n", &struct{}{})
	fmt.Printf("%p\n", &struct{}{})
}
