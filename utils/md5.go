package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5Value(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5Value := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5Value
}
