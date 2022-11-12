package config

import (
	"os"
)

var DbPath string
var DbName = "switchENV.db"

var BatPath string
var BatIdentify = "bat"

func IsDev() bool {
	return false
}

func Init() {
	DbPath, _ = os.Getwd()
	if IsDev() {
		DbPath = "D:\\CodeSoft\\switchEnv"
	}


	BatPath, _ = os.Getwd()
	if IsDev() {
		BatPath = "D:\\CodeSoft\\switchEnv"
	}

}