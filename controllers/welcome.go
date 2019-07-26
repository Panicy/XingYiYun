package controllers

import "github.com/astaxie/beego"

type WelController struct {
	beego.Controller
}
func (w *WelController)Get(){
	w.TplName="html/welcome.html"
}
