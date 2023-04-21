package goframe

import (
	"context"
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"testing"
)

var ctx = context.Background()
var link = "mysql:root:JCPHqknyy8ATR5ME@tcp(192.168.10.47:3306)/oldme?loc=Local"

func TestNew(t *testing.T) {
	fmt.Println(213)
	// 这个db包含了driver和core
	db, err := gdb.New(gdb.ConfigNode{
		Link:  link,
		Debug: true,
	})
	if err != nil {
		panic(err)
	}
	db2 := db.Ctx(ctx).Model("article")
	db2 = db2.Where("id =?", "2")
	db2 = db2.Where("author=?", "half")
	//db2 = db2.LeftJoin("article_grp", "article.grp_id=article_grp.id")
	data, _ := db2.All()
	fmt.Println(data)
}
