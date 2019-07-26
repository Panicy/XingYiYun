package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	b:=CheckSion(c)
	fmt.Print("%t\n",b)
	if b{
		fmt.Println("不存在")
		c.Redirect("/login",302)
		return
	}
	fmt.Println("存在")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
func CheckSion(c *MainController) bool{
	p:=c.GetSession("phone")
	pwd:=c.GetSession("pwd")
	fmt.Println(p,pwd)
	if p!=nil && pwd!=nil{
		return false
	}else{
		return true
	}

}