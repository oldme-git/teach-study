package other

import (
	"encoding/json"
	"testing"
)

type User struct {
	Pwd  string `json:"pwd"`
	PwdB string `json:"pwdbc"`
	Age  int    `json:"age"`
}

// 常用掩盖字段
type Out struct {
	A *struct{} `json:"pwd,omitempty"`
	B *struct{} `json:"pwdbc,omitempty"`
}
type UserOut struct {
	Out
	User
}

func TestA(t *testing.T) {
	u := User{Pwd: "123", Age: 11}
	bb := UserOut{User: u}
	b, _ := json.MarshalIndent(bb, "", "")
	t.Log(string(b))
}
