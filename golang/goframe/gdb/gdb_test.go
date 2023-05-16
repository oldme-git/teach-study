package gdb

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"testing"
)

var (
	ctx  = context.Background()
	link = "mysql:root:123456@tcp(192.168.10.47:3306)/oldme?loc=Local"
	db   = getDb()
	core = db.GetCore().GetDB().GetCore()
)

func getDb() gdb.DB {
	// 这个db包含了driver和core
	db, err := gdb.New(gdb.ConfigNode{
		Link:  link,
		Debug: true,
		//QueryTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	return db
}

func TestCheckLocalTypeForField(t *testing.T) {
	field, err := core.CheckLocalTypeForField(ctx, "varbinary", "")
	if err != nil {
		return
	}
	t.Log(field)
}
