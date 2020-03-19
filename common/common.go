package common

import (
	"crypto/md5"
	"fmt"
	"io"
)

// 用户登录表结构体
type UserLogin struct {
	Id       int    `db:"id"`
	Mobile   string `db:"mobile"`
	Password string `db:"password"`
}

func Md5V(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}
