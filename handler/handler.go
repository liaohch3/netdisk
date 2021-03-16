package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"netdisk/entity"
	"netdisk/model"
	"netdisk/utils"
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

		location := fmt.Sprintf("./tmp/%v", header.Filename)
		newFile, err := os.Create(location)
		if err != nil {
			io.WriteString(w, fmt.Sprintf("create new file fail, err: %v", err))
			return
		}
		defer newFile.Close()

		size, err := io.Copy(newFile, file)
		if err != nil {
			io.WriteString(w, fmt.Sprintf("copy file fail, err: %v", err))
			return
		}

		// 更新文件meta信息
		newFile.Seek(0, 0)
		sha1 := utils.FileSha1(newFile)
		err = model.CreateFileMeta(sha1, header.Filename, size, location)
		if err != nil {
			io.WriteString(w, fmt.Sprintf("CreateFileMeta fail, err: %v", err))
			// todo 这里失败最好可以起一个协程删掉文件
			return
		}
		fmt.Println(sha1)

		// 重定向到上传成功页面
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func UploadSucPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "upload success...")
}

func GetFileMeta(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filePath := r.Form.Get("file_hash")
	fileMeta, err := entity.GetFileMetaBySha1(filePath)
	if err != nil {
		// todo 处理not found
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, fmt.Sprintf("GetFileMetaBySha1 fail, err: %v", err))
		return
	}
	data, err := json.Marshal(fileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)

	fmt.Println(string(data))
}

func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filePath := r.Form.Get("file_hash")
	fileMeta, err := entity.GetFileMetaBySha1(filePath)
	if err != nil {
		// todo 处理not found
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, fmt.Sprintf("GetFileMetaBySha1 fail, err: %v", err))
		return
	}
	// todo 确定权限常量
	file, err := os.Open(fileMeta.Location)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("open file fail, err: %v", err))
		return
	}
	defer file.Close()

	// todo 文件比较大的话，需要做分批读入
	data, err := ioutil.ReadAll(file)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("read file fail, err: %v", err))
		return
	}

	w.Header().Set("Content-Type", "application/x-msdownload;charset=utf-8")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%v", fileMeta.Name))
	w.Write(data)
}

// todo 校验post请求方法
func DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filePath := r.Form.Get("file_hash")
	fileMeta, err := entity.GetFileMetaBySha1(filePath)
	if err != nil {
		// todo 处理not found
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, fmt.Sprintf("GetFileMetaBySha1 fail, err: %v", err))
		return
	}

	err = entity.PhysicalDelFileMeta(fileMeta.Sha1)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("LogicalDelFileMeta fail, err: %v", err))
		return
	}
	// todo 目前都是物理删除，最好找时间处理一下逻辑删除和物理删除
	err = os.Remove(fileMeta.Location)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("delete file fail, err: %v", err))
		return
	}
	w.WriteHeader(http.StatusOK)
}
