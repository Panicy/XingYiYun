package controllers

import "github.com/astaxie/beego"

type BannerController struct {
	beego.Controller
}

func (b *BannerController) Get(){
	b.TplName="html/banner.html"
}