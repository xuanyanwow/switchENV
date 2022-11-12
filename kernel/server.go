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

	r := gin.Default()
	// 开启跨域
	r.Use(Cors())

	router.Init(r)


	//fmt.Println(address+":"+port)
	//os.Exit(1)

	err := r.Run(ServerAddress())
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080
}

func ServerAddress() string {
	address := config.Cfg.Section("gin").Key("address").String()
	port := config.Cfg.Section("gin").Key("port").String()
	if address == "" {
		address = "127.0.0.1"
	}
	if port == "" {
		port = "8080"
	}
	return address+":"+port
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