package entity

import "time"

type User struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Pwd         string    `json:"pwd"`
	UserStatus  int8      `json:"status" gorm:"column:status"`
	CreatedTime time.Time `json:"created_time"`
	LastActive  time.Time `json:"last_active"`
}

func (user *User) TableName() string {
	return "user"
}

func CreateUser(user *User) error {
	return GetDB().Model(&User{}).Create(user).Error
}

func GetUserByName(name string) (*User, error) {
	user := &User{}
	err := GetDB().Model(User{}).Where("name = ?", name).First(user).Error
	return user, err
}
