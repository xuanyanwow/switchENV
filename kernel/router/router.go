package router

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"project/kernel/controller"
	"project/kernel/public"
	"project/kernel/public/static"
)

func Init(r *gin.Engine){
	templateSetting := template.Must(template.New("").ParseFS(public.PEmbed, "*.html"))
	r.SetHTMLTemplate(templateSetting)

	// 比如我的映射规则是public/*   那么访问需要public/public/....
	// 所以解决路径的各种重叠问题，直接在目标路径里新建go文件 声明embed
	r.StaticFS("/static", http.FS(static.SEmbed))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{

		})
	})


	r.POST("/get_lists", func(c *gin.Context) {
		controller.GetSoftVersionLists(c)
	})
	r.POST("/create", func(c *gin.Context) {
		controller.CreateSoftVersion(c)
	})

	r.POST("/delete", func(c *gin.Context) {
		controller.DeleteSoftVersion(c)
	})

	r.POST("/switch", func(c *gin.Context) {
		controller.SwitchSoftVersion(c)
	})

	r.GET("/stop", func(c *gin.Context) {
		c.String(200, "程序结束运行")
		controller.OpenChannel <- 0
	})
}