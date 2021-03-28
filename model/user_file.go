package model

import (
	"netdisk/entity"
	"netdisk/utils"
	"time"
)

func CreateFileMetaAndBindUserFile(fileMeta *entity.FileMeta, userID int64) error {
	user, err := entity.GetUserByUserID(userID)
	if err != nil {
		return err
	}

	userFile := &entity.UserFile{
		Id:          utils.GenId(),
		UserId:      user.Id,
		FileId:      fileMeta.Id,
		DeleteFlag:  entity.DeleteFlag_Default,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	return entity.CreateUserFile(userFile)
}
