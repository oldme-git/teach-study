package gdb

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
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
		Link:   link,
		Debug:  true,
		DryRun: true,
		Prefix: "pre",
		//QueryTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	return db
}

func TestBase(t *testing.T) {
	data, err := db.Model("saying").All()
	if err != nil {
		panic(err)
	}
	t.Log(data)
}

func TestBaseInsert(t *testing.T) {
	data, err := db.Model("saying").Data(g.Map{"saying": "ok"}).Insert()
	if err != nil {
		panic(err)
	}
	t.Log(data)
}
