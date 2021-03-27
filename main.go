package main

import (
	"fmt"
	"net/http"
	"netdisk/cache"
	"netdisk/entity"
	"netdisk/handler"

	"github.com/gin-gonic/gin"
)

// todo 改为gin框架
func main() {
	initial()

	router := gin.Default()

	// 静态资源处理
	router.StaticFS("/static/", http.Dir("./static"))
	router.LoadHTMLGlob("./static/view/*")

	file := router.Group("/file")
	{
		// todo 中间件配置化
		file.GET("/upload", handler.UploadHandler)
		file.POST("/upload", handler.DoUploadHandler)
		file.GET("/upload/suc", handler.UploadSucPage)
		file.GET("/meta", handler.GetFileMeta)
		file.POST("/download", handler.DownloadFileHandler)
		file.POST("/delete", handler.DeleteFileHandler)
		file.POST("/query", handler.QueryFileHandler)
		file.POST("/fastupload", handler.TryFastUploadHandler)
	}

	user := router.Group("/user")
	{
		user.GET("/signup", handler.SignUpHandler)
		user.POST("/signup", handler.DoSignUpHandler)
		user.GET("/signin", handler.SignInHandler)
		user.POST("/signin", handler.DoSignInHandler)
		user.POST("/info", handler.UserInfoHandler)
	}

	err := router.Run(":13081")
	if err != nil {
		panic(fmt.Sprintf("server start err: %v", err))
	}
}

func initial() {
	entity.InitOrm()
	cache.InitSessionMap()
	cache.InitCache()
}
