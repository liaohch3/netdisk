package service

import (
	"netdisk/entity"
	"netdisk/utils"
	"time"
)

func CreateFileMetaAndBindUserFile(fileMeta *entity.FileMeta, userName string) error {
	user, err := entity.GetUserByName(userName)
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
