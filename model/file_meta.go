package model

import (
	"netdisk/entity"
	"netdisk/utils"
	"time"
)

func CreateFileMeta(sha1, name string, size int64, location string) error {
	meta := &entity.FileMeta{
		Id:          utils.GenId(),
		Sha1:        sha1,
		Name:        name,
		Size:        size,
		Location:    location,
		DeleteFlag:  entity.DeleteFlag_Default,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	return entity.CreateFileMeta(meta)
}

func NewFileMeta(sha1, name string, size int64, location string) *entity.FileMeta {
	return &entity.FileMeta{
		Id:          utils.GenId(),
		Sha1:        sha1,
		Name:        name,
		Size:        size,
		Location:    location,
		DeleteFlag:  entity.DeleteFlag_Default,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
}
