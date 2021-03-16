package utils

import "time"

// todo 雪花算法做id generator，暂时这里用时间戳
func GenId() int64 {
	return time.Now().Unix()
}
