package main

import (
	"XingYiYun/models"
	_ "XingYiYun/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//打印DUB
	orm.Debug=true
	//自动建表
	//第一个参数默认数据库，第二个是否每次都自动建表，第三 是否打印相关信息
	orm.RunSyncdb("default",false,true)
	beego.Run()
}

func init(){
	models.RegisterDB()
}
