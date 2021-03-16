package entity

type DeleteFlag int64

const (
	DeleteFlag_Default      = 0
	DeleteFlag_Logical_Del  = 1
	DeleteFlag_Physical_Del = 2
)
