package meta

import "time"

type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt time.Time
}

// key: 文件的sha1值, value: 文件meta信息
var fileMetas map[string]FileMeta

func InitFileMetas() {
	fileMetas = make(map[string]FileMeta)
}

// 新增或更新filemeta
func UpdateFileMetas(meta FileMeta) {
	fileMetas[meta.FileSha1] = meta
}

// 获取文件元信息
func GetFileMeta(fileSha1 string) (FileMeta, bool) {
	fileMeta, ok := fileMetas[fileSha1]
	return fileMeta, ok
}
