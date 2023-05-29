package gdb

import (
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"log"
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

// 不同数据库数据类型的数据，统一使用golang的变量的类型
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

// DoCommit 将sql提交到link中执行
func TestDoCommit(t *testing.T) {
	link, err := db.GetCore().GetLink(ctx, false, "oldme")
	if err != nil {
		panic(err)
	}

	in := &gdb.DoCommitInput{
		Db:            nil,
		Tx:            nil,
		Stmt:          nil,
		Link:          link,
		Sql:           "select * from saying where id = ?",
		Args:          []interface{}{1},
		Type:          gdb.SqlTypeQueryContext,
		IsTransaction: false,
	}

	out, err := db.DoCommit(ctx, *in)
	if err != nil {
		panic(err)
	}
	log.Println(out)
}

func TestFilteredLink(t *testing.T) {
	// 获取可用的link链接
	f := core.FilteredLink()
	// root@tcp(192.168.10.47:3306)/oldme
	fmt.Println(f)
}

func TestQuotePrefixTableName(t *testing.T) {
	s := core.QuotePrefixTableName("saying")
	fmt.Println(s)
}

func TestQuoteString(t *testing.T) {
	s := core.QuoteString("saying")
	fmt.Println(s)
}

func TestQuoteWord(t *testing.T) {
	s := core.QuoteWord("admin")
	fmt.Println(s)
}

// 获取库的表，当前方法是空的
func TestTables(t *testing.T) {
	tables, _ := core.Tables(ctx, "oldme")
	fmt.Println(tables)
}

// 获取表的字段，当前方法是空的
func TestTableFields(t *testing.T) {
	fields, _ := core.TableFields(ctx, "saying", "oldme")
	fmt.Println(fields)
}
