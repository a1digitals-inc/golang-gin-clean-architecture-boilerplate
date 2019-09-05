package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "maman:123459@tcp(127.0.0.1:3306)/product?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return db, err
}
