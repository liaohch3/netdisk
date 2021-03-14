package main

import (
	"net/http"
	"netdisk/handler"
)

func initHandler() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucPage)
	http.HandleFunc("/file/meta", handler.GetFileMeta)
	http.HandleFunc("/file/download", handler.DownloadFileHandler)
}
