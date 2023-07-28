package bench

import (
	"encoding/json"
	"testing"
)

// 探究内存对齐和JSON解码之间的性能影响
// 结论，尽可能使结构体大小
type User struct {
	Pwd int8 `json:"pwd"`
	Age int8 `json:"age"`
	Sex bool `json:"sex"`
}

type UserSafeStruct struct {
	User
	Pwd *struct{} `json:"pwd,omitempty"`
}

type UserSafeOther struct {
	User
	Pwd bool `json:"pwd,omitempty"`
}

var JsonCount = 5000000

func BenchmarkJsonStruct(t *testing.B) {
	for i := 0; i < JsonCount; i++ {
		user := UserSafeStruct{
			User: User{
				Pwd: 123,
				Age: 18,
			},
		}
		json.MarshalIndent(user, "", "\t")
	}
}

func BenchmarkJsonBool(t *testing.B) {
	for i := 0; i < JsonCount; i++ {
		user := UserSafeOther{
			User: User{
				Pwd: 123,
				Age: 18,
			},
		}
		json.MarshalIndent(user, "", "\t")
	}
}

//func TestA(t *testing.T) {
//	fmt.Println(unsafe.Sizeof(UserSafeStruct{}))
//	fmt.Println(unsafe.Sizeof(UserSafeOther{}))
//}
