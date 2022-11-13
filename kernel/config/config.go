package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var DbPath string
var DbName string

var BatPath string
var BatIdentify string

var Cfg *ini.File

func IsDev() bool {
	return Cfg.Section("").Key("app_mode").String() == "dev"
}

func Init() {
	iniFileExists, _ := PathExists("set.ini")
	if !iniFileExists {
		createIni()
	}

	Cfg, _ = ini.Load("set.ini") //初始化一个cfg

	DbPath = Cfg.Section("").Key("db_path").String()
	fmt.Println(DbPath)
	if DbPath == "" {
		DbPath, _ = os.Getwd()
	}
	fmt.Println(DbPath)

	DbName = Cfg.Section("").Key("db_name").String()
	if DbName == "" {
		DbName = "switchENV.db"
	}
	fmt.Println(DbName)

	BatPath = Cfg.Section("").Key("bat_path").String()
	if BatPath == "" {
		BatPath, _ = os.Getwd()
	}

	BatIdentify = Cfg.Section("").Key("bat_identify").String()
	if BatIdentify == "" {
		BatIdentify = "bat"
	}

	//fmt.Println(DbPath)
	//fmt.Println(DbName)
	//fmt.Println(BatPath)
	//fmt.Println(BatIdentify)
	//os.Exit(1)

}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func createIni() {
	file6, err := os.Create("set.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	data := "app_mode = dev\n\ndb_path = \ndb_name = \n\nbat_path = \n\nbat_identify = \n\n[gin]\naddress = 127.0.0.1\nport = 8899"

	_, err = file6.WriteString(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = file6.Close()
}
