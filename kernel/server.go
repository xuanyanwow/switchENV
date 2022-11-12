package kernel

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
	// 开启跨域
	Cors()

	r := gin.Default()

	router.Init(r)

	err := r.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method

		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}