package models

import (
	"fmt"
	"log"
	"todo/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open(mysql.Open(config.Mysql), &gorm.Config{})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	fmt.Println("MYSQL: " + config.Mysql + " CONNECTED")
}
