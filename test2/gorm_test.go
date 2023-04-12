package test2

import (
	"context"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"testing"
)

var ctx = context.Background()
var link = "mysql:root:JCPHqknyy8ATR5ME@tcp(192.168.10.47:3306)/service?loc=Local"

func TestNew(t *testing.T) {
	// 这个db包含了driver和core
	db, err := gdb.New(gdb.ConfigNode{
		Link:  link,
		Debug: true,
	})
	if err != nil {
		panic(err)
	}
	db2 := db.Ctx(ctx).Model("student")
	db2 = db2.Where("stud_no", "20001")
	db2 = db2.Where("stud_no", "20002")
	db2.All()
	//fmt.Println(d)
}
