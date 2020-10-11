package models

import (
	"BeegoWork/db_mysql"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	Id int  'form:"id"'
	Phone string  'form:"phone"'
	Password string 'form:"password"'
}

func (u User) SaveUser() {
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	bytes := hashMd5.Sum(nil)
	u.Pwd = hex.EncodeToString(passwordBytes)
	db_mysql.Db.Exec("")
}