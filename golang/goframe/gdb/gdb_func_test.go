package gdb

import (
	"github.com/gogf/gf/v2/database/gdb"
	"log"
	"testing"
)

func TestFormatSqlWithArgs(t *testing.T) {
	sql := "update oldme set key = ?, key1 = ? where id = ? "
	s := gdb.FormatSqlWithArgs(sql, []interface{}{"ok", nil, 1})
	log.Println(s)
}

func TestIsConfigured(t *testing.T) {
	log.Println(gdb.IsConfigured())
}

func TestListItemValues(t *testing.T) {
	var list = []map[string]map[string]string{
		{
			"key": {
				"key1": "value",
			},
		},
		{
			"key": {
				"key1": "value2",
			},
		},
		{
			"key": {
				"key1": "value",
			},
		},
	}

	r := gdb.ListItemValues(list, "key", "key1")
	log.Println(r)
}
