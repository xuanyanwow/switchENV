package service

import (
	"fmt"
	"os"
	"project/kernel/config"
	"project/kernel/model"
)

// CreateSoftBat 添加版本的时候 检测文件是否存在 不存在则创建
func CreateSoftBat(data *model.Version) {
	filePath := buildPath(data)
	needDo, _ := PathExists(filePath)
	needDo = !needDo
	if needDo == false {
		return
	}

	writeBat(filePath, data.SoftName, data.SoftPath)
}

// SwitchSoftBat 切换bat映射
func SwitchSoftBat(data *model.Version) {
	filePath := buildPath(data)
	writeBat(filePath, data.SoftName, data.SoftPath)
}

// DeleteSoftBat 数据库中所有的软件版本被删除后，需要把映射bat删除
func DeleteSoftBat(data *model.Version) {
	filePath := buildPath(data)
	os.Remove(filePath)
	if config.IsDev() {
		fmt.Println("删除文件-->")
		fmt.Println(filePath)
	}
}

func writeBat(filePath string, softName string, softPath string) {
	if config.IsDev() {
		fmt.Println("写入文件-->")
		fmt.Println(filePath)
		fmt.Println(softName)
		fmt.Println(softPath)
	}
	file6, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := "@echo OFF\n:: in case DelayedExpansion is on and a path contains ! \nsetlocal DISABLEDELAYEDEXPANSION\n" + softPath + softName + " %*"

	_, err = file6.WriteString(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = file6.Close()
}

func buildPath(data *model.Version) string {

	BasePath := config.BatPath
	BaseBatIdentify := config.BatIdentify

	needInitDir, _ := PathExists(BasePath + "/" + BaseBatIdentify)
	needInitDir = !needInitDir

	if config.IsDev() {
		fmt.Println("需要创建目录 --> ")
		fmt.Println(needInitDir)
	}

	if needInitDir {
		err := os.Mkdir(BasePath+"/"+BaseBatIdentify, 0777)
		if err != nil {
			panic("软件初始化目录错误,请联系管理员")
		}
	}

	if config.IsDev() {
		fmt.Println(BasePath + "/" + BaseBatIdentify + "/" + data.SoftName + ".bat")
	}

	return BasePath + "/" + BaseBatIdentify + "/" + data.SoftName + ".bat"
}
