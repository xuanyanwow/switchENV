package controller

import (
	"github.com/gin-gonic/gin"
	"project/kernel/db"
	"project/kernel/model"
	"project/kernel/service"
)


func GetSoftVersionLists(c *gin.Context) {

	var versionList []model.Version
	db.Db.Find(&versionList)

	//fmt.Println(reflect.TypeOf(versionList))
	//p, _ := json.Marshal(versionList)

	c.String(200, Json(200, map[string]interface{}{
		"list": versionList,
	}, "获取成功"))
}

func CreateSoftVersion(c *gin.Context){

	name:= c.PostForm("name")
	version:= c.PostForm("version")
	path:= c.PostForm("path")

	if name == "" {
		End(c, FailParam(map[string]interface{}{}, "软件名不可为空"))
		return
	}
	if version == "" {
		End(c, FailParam(map[string]interface{}{}, "软件版本不可为空"))
		return
	}
	if path == "" {
		End(c, FailParam(map[string]interface{}{}, "软件路径不可为空"))
		return
	}

	var hasVersionModel model.Version
	result := db.Db.Where(&model.Version{
		SoftName: name,
		SoftVersion: version,
	}).Find(&hasVersionModel)

	if result.RowsAffected > 0  {
		End(c, Fail(map[string]interface{}{}, "当前软件+版本已经存在，请检查或者更换软件标识"))
		return
	}

	versionModel := model.Version{
		SoftName:name,
		SoftVersion:version,
		SoftPath: path,
		IsChoose:0,
	}

	_ = db.Db.Create(&versionModel)


	service.CreateSoftBat(&versionModel)

	c.String(200, Json(200, map[string]interface{}{
		"id" : versionModel.ID,
	}, "创建成功"))
}



func SwitchSoftVersion(c *gin.Context){
	name:= c.PostForm("name")
	version:= c.PostForm("version")

	if name == "" {
		End(c, FailParam(map[string]interface{}{}, "软件名不可为空"))
		return
	}
	if version == "" {
		End(c, FailParam(map[string]interface{}{}, "软件版本不可为空"))
		return
	}

	var versionModel model.Version

	result := db.Db.Where(&model.Version{
		SoftName: name,
		SoftVersion: version,
	}).Find(&versionModel)

	if result.RowsAffected == 0 {
		End(c, Fail(map[string]interface{}{}, "当前软件+版本不存在，请检查或者新建"))
		return
	}


	db.Db.Model(&model.Version{}).Where(&model.Version{
		SoftName: name,
	}).Update("is_choose", 0)


	versionModel.IsChoose = 1
	db.Db.Save(&versionModel)

	service.SwitchSoftBat(&versionModel)

	End(c,Success(map[string]interface{}{}, "切换成功"))
}


func DeleteSoftVersion(c *gin.Context){
	name:= c.PostForm("name")
	version:= c.PostForm("version")

	if name == "" {
		End(c, FailParam(map[string]interface{}{}, "软件名不可为空"))
		return
	}
	if version == "" {
		End(c, FailParam(map[string]interface{}{}, "软件版本不可为空"))
		return
	}


	var versionModel model.Version
	result := db.Db.Where(&model.Version{
		SoftName: name,
		SoftVersion: version,
	}).Find(&versionModel)

	if result.RowsAffected == 0 {
		End(c, Fail(map[string]interface{}{}, "当前软件+版本不存在，请检查或者新建"))
		return
	}

	db.Db.Delete(&versionModel)

	service.DeleteSoftBat(&versionModel)

	End(c,Success(map[string]interface{}{}, "删除成功"))
}