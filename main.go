package main

import (
	"fmt"
	"net/http"
	"netdisk/handler"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucPage)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(fmt.Sprint("server start err: %v", err))
	}
}
