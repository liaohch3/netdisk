package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"netdisk/consts"
	"netdisk/entity"
	"netdisk/model"
	"netdisk/service"
	"netdisk/utils"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	// 请求时返回index页面
	c.HTML(http.StatusOK, "index.html", nil)
}

func DoUploadHandler(c *gin.Context) {
	userID := c.GetInt64(consts.USER_ID)
	header, err := c.FormFile("file")
	if err != nil {
		c.String(-1, "got form file fail, err: %v", err)
		return
	}
	fmt.Printf("fileName: %v, fileSize: %v\n", header.Filename, header.Size)

	err = os.MkdirAll("tmp", os.ModeDir)
	if err != nil {
		c.String(-1, "create new dir fail, err: %v", err)
		return
	}

	location := fmt.Sprintf("./tmp/%v", header.Filename)
	newFile, err := os.Create(location)
	if err != nil {
		c.String(-1, "create new file fail, err: %v", err)
		return
	}
	defer newFile.Close()

	err = c.SaveUploadedFile(header, location)
	if err != nil {
		c.String(-1, "copy file fail, err: %v", err)
		return
	}

	newFile.Seek(0, 0)
	sha1 := utils.FileSha1(newFile)
	fileMeta := model.NewFileMeta(sha1, header.Filename, header.Size, location)
	fmt.Println(sha1)

	err = service.CreateFileMetaAndBindUserFile(fileMeta, userID)
	if err != nil {
		c.String(-1, "CreateFileMetaAndBindUserFile fail, err: %v", err)
		return
	}

	err = fileMeta.Save()
	if err != nil {
		c.String(-1, "fileMeta.Save() fail, err: %v", err)
		return
	}

	// 重定向到上传成功页面
	c.Redirect(http.StatusFound, "/static/view/home.html")
}

func UploadSucPage(c *gin.Context) {
	c.String(0, "upload success...")
}

func GetFileMeta(c *gin.Context) {
	filePath := c.Query("file_hash")
	fileMeta, err := entity.GetFileMetaBySha1(filePath)
	if err != nil {
		// todo 处理not found
		//w.WriteHeader(http.StatusInternalServerError)
		c.String(http.StatusInternalServerError, "GetFileMetaBySha1 fail, err: %v", err)
		return
	}
	data, err := json.Marshal(fileMeta)
	if err != nil {
		c.String(http.StatusInternalServerError, "Marshal fail")
		return
	}
	c.String(0, string(data))

	fmt.Println(string(data))
}

func DownloadFileHandler(c *gin.Context) {
	filePath := c.Query("file_hash")
	fileMeta, err := entity.GetFileMetaBySha1(filePath)
	if err != nil {
		// todo 处理not found
		c.String(http.StatusInternalServerError, "GetFileMetaBySha1 fail, err: %v", err)
		return
	}
	// todo 确定权限常量
	file, err := os.Open(fileMeta.Location)
	if err != nil {
		c.String(-1, "open file fail, err: %v", err)
		return
	}
	defer file.Close()

	// todo 文件比较大的话，需要做分批读入
	data, err := ioutil.ReadAll(file)
	if err != nil {
		c.String(-1, "read file fail, err: %v", err)
		return
	}

	c.Header("Content-Type", "application/x-msdownload;charset=utf-8")
	c.Header("Content-Disposition", fmt.Sprintf("attachment;filename=%v", fileMeta.Name))
	c.String(0, string(data))
}

// todo 校验post请求方法
func DeleteFileHandler(c *gin.Context) {
	filePath := c.Query("file_hash")
	fileMeta, err := entity.GetFileMetaBySha1(filePath)
	if err != nil {
		// todo 处理not found
		c.String(http.StatusInternalServerError, "GetFileMetaBySha1 fail, err: %v", err)
		return
	}

	err = entity.PhysicalDelFileMeta(fileMeta.Sha1)
	if err != nil {
		c.String(-1, "LogicalDelFileMeta fail, err: %v", err)
		return
	}
	// todo 目前都是物理删除，最好找时间处理一下逻辑删除和物理删除
	err = os.Remove(fileMeta.Location)
	if err != nil {
		c.String(-1, "delete file fail, err: %v", err)
		return
	}
	c.String(http.StatusOK, "")
}

//todo 规范所有的c.String 换成c.JSON好一点
func QueryFileHandler(c *gin.Context) {
	userID := c.GetInt64(consts.USER_ID)
	user, err := entity.GetUserByUserID(userID)
	if err != nil {
		c.String(-1, "GetUserByName fail, err: %v", err)
		return
	}
	userFiles, err := entity.GetUserFileByUserId(user.Id)
	if err != nil {
		c.String(-1, "GetUserFileByUserId fail, err: %v", err)
		return
	}

	files := []*entity.FileMeta{}
	for _, userFile := range userFiles {
		file, err := entity.GetFileMetaByFileId(userFile.Id)
		if err != nil {
			c.String(-1, "GetFileMetaByFileId fail, err: %v", err)
			return
		}
		files = append(files, file)
	}

	c.JSON(0, utils.NewSuccessMsg(files))
}

// todo 处理所有返回值
func TryFastUploadHandler(c *gin.Context) {
	userID := c.GetInt64(consts.USER_ID)
	hash := c.Query("file_hash")

	fileMeta, err := entity.GetFileMetaBySha1(hash)
	// 还需要更准确地判断成notFound才行
	if err != nil {
		c.String(-1, "Not Uploaded")
		return
	}

	user, err := entity.GetUserByUserID(userID)
	if err != nil {
		c.String(-1, "GetUserByName fail, err: %v", err)
		return
	}

	userFile := &entity.UserFile{
		Id:          utils.GenId(),
		UserId:      user.Id,
		FileId:      fileMeta.Id,
		DeleteFlag:  entity.DeleteFlag_Default,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	err = entity.CreateUserFile(userFile)
	if err != nil {
		c.String(-1, "CreateUserFile fail, err: %v", err)
		return
	}

	c.String(http.StatusOK, "")
}
