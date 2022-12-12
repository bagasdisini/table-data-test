package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	connect := "root:@tcp(127.0.0.1:3306)/databaseTest?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database Connected!")
}
