package db

import (
	"database/sql"
	"fmt"
	_ "github.com/logoove/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"project/kernel/config"
	"project/kernel/model"
)

var Db *gorm.DB

func Init() {
	fmt.Println(config.DbPath + "/" + config.DbName)
	open, err := sql.Open("sqlite3", config.DbPath+config.DbName)
	if err != nil {
		return
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: open,
	}), &gorm.Config{})

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
