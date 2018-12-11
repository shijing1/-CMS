package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ItcastCms/models"
)

type LoginController struct {
	beego.Controller
}

func(this *LoginController) Index()  {
  this.TplName="Login/Index.html"
}
func (this *LoginController)UserLogin()  {
	userName:=this.GetString("LoginCode")
	LoginPwd:=this.GetString("LoginPwd")
	o:=orm.NewOrm()
	var userInfo models.UserInfo
	o.QueryTable("user_info").Filter("user_name",userName).Filter("user_pwd",LoginPwd).One(&userInfo)
	if userInfo.Id>0{
		this.SetSession("userId",userInfo.Id)
		this.SetSession("userName",userInfo.UserName)
		this.Data["json"]=map[string]interface{}{"flag":"ok"}
	}else{
		this.Data["json"]=map[string]interface{}{"flag":"no"}
	}
	this.ServeJSON()
}
