package mytest

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestQuery(t *testing.T) {
	dsn := "root:JCPHqknyy8ATR5ME@tcp(192.168.10.47:3306)/service?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	type Saying struct {
		gorm.Model
		Id     int
		Saying string
	}
	var s []Saying
	db.Table("saying").Where("id = ?", 51552).Find(&s)
	fmt.Println(s)
}
