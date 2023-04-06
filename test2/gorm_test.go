package test2

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/v2/database/gdb"
	"testing"
)

var link = "mysql:root:JCPHqknyy8ATR5ME@tcp(192.168.10.47:3306)/oldme?loc=Local"

func TestNew(t *testing.T) {
	db, err := gdb.New(gdb.ConfigNode{
		Link: link,
	})
	if err != nil {
		panic(err)
	}
}
