package b

import "fmt"

var (
	globalInt int
)

func Reg(i int) {
	globalInt = i
}

func GetInt() {
	fmt.Println(globalInt)
}
