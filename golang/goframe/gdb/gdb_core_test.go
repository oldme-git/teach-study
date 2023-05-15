package gdb

import "fmt"

// 统一不同数据库字段，统一使用golang的变量的类型
func ExampleCheckLocalTypeForField() {
	var (
		db   = getDb()
		core = db.GetCore()
	)
	field, err := core.CheckLocalTypeForField(ctx, "varbinary", "")
	if err != nil {
		return
	}
	fmt.Println(field)
	// Output:
	// []byte
}
