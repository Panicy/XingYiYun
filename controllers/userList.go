package controllers

import (
	"XingYiYun/models"
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func  (u *UserController) Get(){
	var err error
	u.Data["users"],err=models.UserDB()
	fmt.Println(models.UserDB())
	u.TplName="html/user.html"
	if err!=nil{
		fmt.Println(err)
	}

}
