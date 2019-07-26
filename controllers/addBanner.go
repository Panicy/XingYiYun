package controllers

import "github.com/astaxie/beego"

type AddBanner struct {
	beego.Controller
}

func (a *AddBanner)Get (){
	a.TplName="html/add_banner.html"
}
