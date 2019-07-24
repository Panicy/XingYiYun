package controllers

import (
	"XingYiYun/commons"
	"XingYiYun/models"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get(){
	l.TplName="login.html"
}

func (l *LoginController) Post(){
	phone:=l.Input().Get("user")
	pwd:=l.Input().Get("pwd")
	u,_:=strconv.Atoi(phone)
	fmt.Println(phone,pwd)
	us,err:=models.LoginDB(u,pwd)
	if err!=nil{
		l.Data["json"]=commons.Jsonstruct{Msg:err,Code:201}
		l.ServeJSON()
		return
	}
	l.SetSession("phone",u)
	l.SetSession("pwd",pwd)
	l.Data["json"]=commons.Jsonstruct{Msg:us,Code:200}
	l.ServeJSON()
	return
}