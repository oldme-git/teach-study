package a

import "base/pkg/b"

func init() {
	b.Reg(100)
}

func ChangeRegInt() {
	b.Reg(200)
}
