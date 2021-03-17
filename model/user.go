package model

import (
	"encoding/json"
	"fmt"
	"netdisk/entity"
	"netdisk/utils"
	"time"
)

func CreateUser(name, passwd, phone, email string) error {
	user := &entity.User{
		Id:          utils.GenId(),
		Name:        name,
		Pwd:         passwd, // todo 加密
		Email:       email,
		Phone:       phone,
		UserStatus:  entity.UserStatus_Default,
		CreatedTime: time.Now(),
		LastActive:  time.Now(),
	}

	b, _ := json.Marshal(user)
	fmt.Println(string(b))

	return entity.CreateUser(user)
}
