package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"log"
)

type UpImageController struct {
	beego.Controller
}

func (this *UpImageController) Post(){
	fmt.Sprintln(this)
	f,h,err:=this.GetFile("file")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	err=this.SaveToFile("file", "static/upload/" + h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	fmt.Println(err)
	var datas interface{}
	if err!=nil{
		fmt.Println("进来的上传失败 ")
		datas= map[string]interface{}{
			"code":200,
			"msg":"上传失败",
		}
	}else{
		fmt.Println("进来的上传成功 ")
		datas= map[string]interface{}{
			"code":0,
			"msg":"上传成功",
			"data":map[string]interface{}{
				"src":"/static/upload/"+h.Filename,
				"title":h.Filename,
			},
		}
	}
	this.Data["json"]=datas
	this.ServeJSON()
}
