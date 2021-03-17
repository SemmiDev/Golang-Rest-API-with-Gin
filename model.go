package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Gagal Conect Ke Database")
	}
	db.AutoMigrate(&student{})
}


type (
	student struct {
		gorm.Model
		Name        	string `json:"name"`
		Address      	string `json:"address"`
		PhoneNumber     string `json:"phone_number"`
		Class       	string `json:"class"`
		ActiveStatus 	int    `json:"active_status"`
	}
	transformedStudent struct {
		ID          	uint   `json:"id"`
		Name        	string `json:"name"`
		Address      	string `json:"address"`
		PhoneNumber     string `json:"phone_number"`
		Class       	string `json:"class"`
		ActiveStatus 	bool   `json:"active_status"`
	}
)