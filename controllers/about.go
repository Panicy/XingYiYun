package controllers

import (
	"XingYiYun/models"
	"github.com/astaxie/beego"
	"strconv"
)

type AboutController struct {
	beego.Controller
}

func (a *AboutController)Get(){
	a.Data["about"]=models.GetAboutDb()
	a.TplName="html/about.html"
}

func (a *AboutController)Post(){
	t:=a.GetString("texts")
	id:=a.GetString("id")
	is,_:=strconv.ParseInt(id,10,64)
	datas:=make(map[string]interface{})
	i,err:=models.AboutDb(t,is)
	if err!=nil{
		datas["code"]=201
		datas["msg"]="更新失败"
	}else{
		datas["code"]=0
		datas["msg"]="更新成功"
		datas["data"]=i
	}
	a.Data["json"]=datas
	a.ServeJSON()
	return
	//a.Redirect("/about",301)
}