package db

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"project/kernel/config"
	"project/kernel/model"
)

var Db *gorm.DB

func Init() {
	fmt.Println(config.DbPath + "/" + config.DbName)
	db, err := gorm.Open(sqlite.Open(config.DbPath+"/"+config.DbName), &gorm.Config{})

	if err != nil {
		panic("error --> failed to connect database " + err.Error())
	}

	Db = db

	migrateErr := db.AutoMigrate(&model.Version{})
	if migrateErr != nil {
		fmt.Println("error --> " + migrateErr.Error())
		return
	}
}
