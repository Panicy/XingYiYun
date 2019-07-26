package models

import (

	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	_DB_NAME="xyy"
	_DB_DRIVER="mysql"
	duser ="xyy"
	dhost ="47.96.102.156"
	dport ="3306"
	DPWD ="romantic9"
	db="xyy"
	mysqlCon=duser+":"+DPWD+"@tcp("+dhost+":"+dport+")/"+db+"?charset=utf8"
)

type User struct {
	Id int64
	Openid int64 `orm:"description(微信openid)"`
	Nickname string `orm:"description(昵称)"`
	Phone int64 `orm:"description(电话)"`
	Tid int64 `orm:"description(推荐id)"`
	Avatar string `orm:"description(头像)"`
	Pwd string `orm:"description(后台密码)"`
	Created time.Time
}
type Banner struct {
	Id int64
	BannerImg string `orm:"description(轮播图url)"`
	Status int64 `orm:"description(轮播图状态)"`
	Operation string `orm:"description(轮播图连接)"`
	State int64 `orm:"description(状态1位启动，0为未启用);default(1)"`
	Created time.Time
}
type Introduce struct {
	Id int64
	Title string `orm:"description(介绍标题)"`
	Content string `orm:"size(5000);description(介绍内容)"`
	Created time.Time
}
type Case struct {
	Id int64
	CaseName string `orm:"description(案例名)"`
	CaseImg string `orm:"description(案例图片)"`
	CaseUrl string `orm:"description(案例连接)"`
	Stete int64 `orm:"description(案例状态)"`
	Created time.Time
}
type Concat struct {
	Id int64
	ConcatDemo string `orm:"description(联系方式)"`
	Created time.Time `orm:"description(创建时间)"`
}
type About struct {
 	Id int64  `orm:"description(用户id)"`
 	Content string `orm:"size(5000);description(关于我们内容)"`
 }

func RegisterDB(){
	orm.RegisterModel(new(User),new(Banner),new(Introduce),new(Case),new(Concat),new(About))
	orm.RegisterDriver(_DB_DRIVER,orm.DRMySQL)
	orm.RegisterDataBase("default",_DB_DRIVER,mysqlCon)
}
func LoginDB(phone int,pwd string) (*User,error){
	o:=orm.NewOrm()
	us:=&User{}
	qs:=o.QueryTable("user")
	err:=qs.Filter("Phone",phone).Filter("Pwd",pwd).One(us)
	fmt.Println(us)
	if err!=nil{
		return nil,err
	}
	return  us,nil
}

func UserDB()( []*User,error){
	o:=orm.NewOrm()
	users:=make([]*User,0)
	qs:=o.QueryTable("user")
	_,err:=qs.All(&users)
	return  users,err
}

