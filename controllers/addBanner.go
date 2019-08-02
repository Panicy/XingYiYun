package controllers

import (
	"XingYiYun/commons"
	"fmt"
	"github.com/astaxie/beego"
)

type AddBanner struct {
	beego.Controller
}

func (a *AddBanner)Get (){
	a.TplName="html/add_banner.html"
}

func (this *AddBanner)Post(){
	status:=this.GetString("status")
	desc:=this.GetString("desc")
	fmt.Println(status)
	fmt.Println(desc)
	f,h,err:=this.GetFile("file")
	var data commons.UploadImg
	if err!=nil{
		data.Code=201
		data.Msg="图片上传失败"
		this.Data["json"]=data
		this.ServeJSON()
		return
	}
	defer f.Close()
	err=this.SaveToFile("file","static/banner"+h.Filename)
	if err!=nil{
		data.Code=201
		data.Msg="图片上传失败"
		this.Data["json"]=data
		this.ServeJSON()
		return
	}
	data.Code=0
	data.Msg="图片上传成功"
	data.Data.Src="/static/banner"+h.Filename
	data.Data.Title=h.Filename
	this.Data["json"]=data
	this.ServeJSON()
	return
}
