package main

import (
	"fmt"
	"net/http"
	"netdisk/cache"
	"netdisk/entity"
	"netdisk/handler"
	"netdisk/middleware"
)

// todo 改为gin框架
func main() {
	initial()

	// 静态资源处理
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(assets.AssetFS())))
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// todo 中间件配置化
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucPage)
	http.HandleFunc("/file/meta", handler.GetFileMeta)
	http.HandleFunc("/file/download", handler.DownloadFileHandler)
	http.HandleFunc("/file/delete", handler.DeleteFileHandler)
	http.HandleFunc("/file/query", handler.QueryFileHandler)
	http.HandleFunc("/file/fastupload", handler.TryFastUploadHandler)

	http.HandleFunc("/user/signup", handler.SignUpHandler)
	http.HandleFunc("/user/signin", handler.SignInHandler)
	http.HandleFunc("/user/info", middleware.AuthHandler(handler.UserInfoHandler))

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(fmt.Sprint("server start err: %v", err))
	}
}

func initial() {
	entity.InitOrm()
	cache.InitSessionMap()
}
