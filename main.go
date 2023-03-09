package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const CREATE_TABLE_SQL = "CREATE TABLE `mytable` (`aaa` int); INSERT INTO `mytable` (`aaa`) VALUES ('1');"

func main() {
	dsn := "sandbox:sandbox@tcp(localhost)/sandbox?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Exec(CREATE_TABLE_SQL).Error; err != nil {
		panic(err)
	}

	if err := tx.Commit().Error; err != nil {
		panic(err)
	}

	fmt.Println("OK")
}
