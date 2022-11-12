package kernel

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"project/kernel/config"
	"project/kernel/db"
	"project/kernel/router"
)


func Server() {

	db.Init()
	if !config.IsDev() {
		fmt.Println("配置 --> 正式模式")
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	router.Init(r)

	err := r.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080
}
