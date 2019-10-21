package utils

import (
	"crypto/md5"
	"fmt"
)

// md5加密处理
func MD5(pwd string) string {
	return  fmt.Sprintf("%x", md5.Sum([]byte(pwd)))
}


