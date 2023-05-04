package goframe

import (
	"context"
	"fmt"
	"time"

	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"testing"
)

var ctx = context.Background()
var link = "mysql:root:123456@tcp(192.168.10.47:3306)/oldme?loc=Local"

func TestNew(t *testing.T) {
	// 这个db包含了driver和core
	db, err := gdb.New(gdb.ConfigNode{
		Link:         link,
		Debug:        true,
		QueryTimeout: 1 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	ctx = context.WithValue(ctx, "k", "v")
	db2 := db.Ctx(ctx).Model("saying")
	db2 = db2.Where("id = ?", 18)
	//db2 = db2.Where("author=?", "half")
	//db2 = db2.LeftJoin("article_grp", "article.grp_id=article_grp.id")
	data, err := db2.All()
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
