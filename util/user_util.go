package util

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5PWD(password string) (string, bool) {
	h := md5.New()
	if _, err := io.WriteString(h, password); err != nil {
		panic(err)
		return "", false
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	return passWord, true
}
