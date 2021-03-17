package utils

import (
	"fmt"
	"time"
)

func GenSession(name string) string {
	return MD5(fmt.Sprintf("%v-%v", name, time.Now().UnixNano()))
}
