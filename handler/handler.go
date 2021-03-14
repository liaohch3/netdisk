package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// GET 请求时返回index页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, fmt.Sprintf("open file fail: %v", err))
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		file, header, err := r.FormFile("file")
		if err != nil {
			io.WriteString(w, fmt.Sprintf("got form file fail, err: %v", err))
			return
		}
		fmt.Printf("fileName: %v, fileSize: %v\n", header.Filename, header.Size)
		defer file.Close()
		// time.Now().Format("2006-01-02-15:04:05")

		err = os.MkdirAll("tmp", os.ModeDir)
		if err != nil {
			io.WriteString(w, fmt.Sprintf("create new dir fail, err: %v", err))
			return
		}

		newFile, err := os.Create(fmt.Sprintf("./tmp/%v", header.Filename))
		if err != nil {
			io.WriteString(w, fmt.Sprintf("create new file fail, err: %v", err))
			return
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, file)
		if err != nil {
			io.WriteString(w, fmt.Sprintf("copy file fail, err: %v", err))
			return
		}

		// 重定向到上传成功页面
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func UploadSucPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "upload success...")
}
