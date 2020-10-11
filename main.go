package main

import (
	"DataCertProject/models"
	_ "DataCertProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.ConnectDB()
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}

