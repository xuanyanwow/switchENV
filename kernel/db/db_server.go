package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"project/kernel/config"
	"project/kernel/model"
)

var Db *gorm.DB





func Init(){
	db, err := gorm.Open(sqlite.Open(config.DbPath+ config.DbName), &gorm.Config{})

	if err != nil {
		panic("error --> failed to connect database")
	}

	Db = db

	migrateErr := db.AutoMigrate(&model.Version{})
	if migrateErr != nil {
		fmt.Println("error --> " + migrateErr.Error())
		return 
	}
}