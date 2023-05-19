package gdb

import (
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"testing"
)

// 不同数据库数据类型，统一使用golang的变量的类型
func ExampleCheckLocalTypeForField() {
	field, err := core.CheckLocalTypeForField(ctx, "varbinary", "")
	if err != nil {
		return
	}
	fmt.Println(field)
	// Output:
	// []byte
}

func ExampleConvertValueForLocal() {
	var (
		fType  = "float(5,2)"
		fValue = "123.1"
	)

	local, err := core.ConvertValueForLocal(ctx, fType, fValue)
	if err != nil {
		return
	}
	fmt.Printf("%T, %v", local, local)

	// Output:
	// float64, 123.1
}

func TestDoCommit(t *testing.T) {
	in := &gdb.DoCommitInput{
		Db:            nil,
		Tx:            nil,
		Stmt:          nil,
		Link:          nil,
		Sql:           "",
		Args:          nil,
		Type:          "",
		IsTransaction: false,
	}
	out, err := core.DoCommit(ctx, *in)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
