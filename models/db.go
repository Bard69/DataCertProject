package models

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)
var Db *sql.DB

func ConnectDB()  {
	config := beego.AppConfig
	dbDriver := config.String("driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
	db, err := sql.Open(dbDriver, connUrl)
	if err != nil {
		panic("数据库连接错误")
	}
	Db = db
}