package entity

import "time"

type FileMeta struct {
	Id          int64     `json:"id"`
	Sha1        string    `json:"sha1"`
	Name        string    `json:"name"`
	Size        int64     `json:"size"`
	Location    string    `json:"location"`
	DeleteFlag  int8      `json:"delete_flag"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}

func (meta FileMeta) TableName() string {
	return "file_meta"
}

func CreateFileMeta(meta *FileMeta) error {
	return GetDB().Model(FileMeta{}).Create(meta).Error
}

func LogicalDelFileMeta(sha1 string) error {
	return GetDB().Model(FileMeta{}).Where("sha1 = ?", sha1).UpdateColumns(map[string]interface{}{
		"delete_flag": DeleteFlag_Logical_Del,
		// todo 逻辑删除应该不更新更新时间吧。。
	}).Error
}

func PhysicalDelFileMeta(sha1 string) error {
	return GetDB().Model(FileMeta{}).Where("sha1 = ?", sha1).UpdateColumns(map[string]interface{}{
		"delete_flag":  DeleteFlag_Physical_Del,
		"updated_time": time.Now(),
	}).Error
}

// todo 定义errcode
func GetFileMetaBySha1(sha1 string) (*FileMeta, error) {
	meta := &FileMeta{}
	err := GetDB().Model(FileMeta{}).Where("sha1 = ?", sha1).First(meta).Error
	return meta, err
}
