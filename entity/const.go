package entity

type DeleteFlag int8

const (
	DeleteFlag_Default      = 0
	DeleteFlag_Logical_Del  = 1
	DeleteFlag_Physical_Del = 2
)

type UserStatus int8

const (
	UserStatus_Default = 0
	UserStatus_Ban     = 1
)
