package entity

import "time"

// 记录用户和文件关系的表

type UserFile struct {
	Id          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	FileId      int64     `json:"file_id"`
	DeleteFlag  int8      `json:"delete_flag"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}

func (userFile *UserFile) TableName() string {
	return "user_file"
}

func CreateUserFile(userFile *UserFile) error {
	return GetDB().Model(&UserFile{}).Create(userFile).Error
}

func GetUserFileByUserId(userId int64) ([]*UserFile, error) {
	userFiles := []*UserFile{}
	err := GetDB().Model(UserFile{}).Where("user_id = ?", userId).Find(&userFiles).Error
	return userFiles, err
}
