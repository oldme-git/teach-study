package gdb

import (
	"fmt"
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
