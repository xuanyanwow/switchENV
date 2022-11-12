package config

import (
	"gopkg.in/ini.v1"
	"os"
)

var DbPath string
var DbName string

var BatPath string
var BatIdentify string

var Cfg, _ = ini.Load("set.ini") //初始化一个cfg

func IsDev() bool {
	return Cfg.Section("").Key("app_mode").String() == "dev"
}

func Init() {
	DbPath = Cfg.Section("").Key("db_path").String()
	if DbPath == "" {
		DbPath, _ = os.Getwd()
	}

	DbName = Cfg.Section("").Key("db_name").String()
	if DbName == "" {
		DbName = "switchENV.db"
	}

	BatPath = Cfg.Section("").Key("bat_path").String()
	if BatPath == "" {
		BatPath, _ = os.Getwd()
	}

	BatIdentify = Cfg.Section("").Key("bat_identify").String()
	if BatIdentify == "" {
		BatIdentify, _ = os.Getwd()
	}

	//fmt.Println(DbPath)
	//fmt.Println(DbName)
	//fmt.Println(BatPath)
	//fmt.Println(BatIdentify)
	//os.Exit(1)

}