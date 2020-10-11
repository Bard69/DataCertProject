package controllers

import (
	"DataCertProject/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post() {
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("解析数据错误")
		return
	}

}