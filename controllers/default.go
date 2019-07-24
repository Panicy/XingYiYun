package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	err:=CheckSion
	fmt.Println("session错误",err)
	fmt.Sprint(err)
	if err!=nil{
		c.Redirect("/login",302)
		return
	}
	//p:=c.GetSecureCookie("phone")
	//p:=c.GetSession("phone")
	//fmt.Println("session电话",p)
	//if p==nil{
	//	fmt.Println("sessionphone")
	//	c.Redirect("/login",302)
	//	return
	//}
	//pd:=c.GetSession("pwd")
	//fmt.Println("session密码",pd)
	//if pd==nil{
	//	c.Redirect("/login",302)
	//	fmt.Println("session密码")
	//	return
	//}
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
func CheckSion(c *MainController) error{
	p:=c.GetSession("phone")
	pwd:=c.GetSession("pwd")
	var err error
	fmt.Println(p,pwd)
	if p!=nil && pwd!=nil{
		return nil
	}else{
		return err
	}

}