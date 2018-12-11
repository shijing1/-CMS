package controllers

import "github.com/astaxie/beego"
import (
	"ItcastCms/models"
	"github.com/astaxie/beego/orm"
)
type HomeController struct {
	beego.Controller
}

func(this *HomeController) Index()  {
  this.TplName="Home/Index.html"
}
func (this *HomeController)ShowIndex()  {
	this.TplName="Home/ShowIndex.html"
}
//菜单权限过滤.
func (this *HomeController)GetMenus()  {
	//首先根据的“用户”---“角色”--"权限"这条线进行过滤。
	//1:根据登录用户，获取对应的用户信息。
	userId:=this.GetSession("userId")//获取seesion中存储的登录用户的编号
	var userInfo models.UserInfo
	o:=orm.NewOrm()
	o.QueryTable("user_info").Filter("id",userId).One(&userInfo)
	//2:根据登录用户信息，查找对应的角色。
	//将查询出的角色，存储到roles切片中。
	var roles []*models.RoleInfo
	o.LoadRelated(&userInfo,"Roles")
	for _,role:=range userInfo.Roles {
		roles=append(roles,role)
	}
	//3:根据角色，查询对应的权限。
	var actions[]*models.ActionInfo
	for i:=0;i<len(roles);i++{
		o.LoadRelated(roles[i],"Actions")
		for _,action:= range roles[i].Actions  {
			actions=append(actions,action)
		}
	}
	//4；判断这些查询出的权限，哪些菜单权限。
	var menuActions []*models.ActionInfo
	for i:=0;i<len(actions) ;i++  {
		if actions[i].ActionTypeEnum==1 {
           menuActions=append(menuActions,actions[i])
		}
	}
	//按照“用户”--“权限”这条线，查询出登录的菜单权限。
	var subActions []models.UserAction
//根据登录用户查询中间表（user_action）
	o.QueryTable("user_action").Filter("users_id",userId).All(&subActions)
	//根据上面的查询，查询的权限编号不能判断是否为菜单权限。所以，还需要查询权限表
	var subMenuActions[]*models.ActionInfo

	for i:=0;i<len(subActions);i++ {
		var actionInfo models.ActionInfo
		o.QueryTable("action_info").Filter("id",subActions[i].Actions.Id).Filter("action_type_enum",1).One(&actionInfo)
		if actionInfo.Id>0{
			subMenuActions=append(subMenuActions,&actionInfo)
		}
	}

   //将两条线，查询出的菜单权限进行合并。
   menuActions=append(menuActions,subMenuActions...)
    //去重操作。
	temp:=RemoveRepeatedElement(menuActions)
	//接下来判断temp中存储的当前登录用户的权限是否有被禁止，如果有，则去掉。
	var userForActions []models.UserAction
	o.QueryTable("user_action").Filter("is_pass",0).Filter("users_id",userId).All(&userForActions)
	//判断是否找到了登录用户的禁用权限。
	if len(userForActions)>0{
       //找到了禁用权限后，清除。
       var newTmep[]*models.ActionInfo
		for i,action:=range temp {
			//判断权限的编号是否在禁用的集合中存在，如果返回的是false,表示该权限没有被禁用，存在newTemp切片中。
			if CheckUserForAction(userForActions,action.Id)==false{
				newTmep=append(newTmep,temp[i])
			}
		}
		this.Data["json"]=map[string]interface{}{"menus":newTmep}

	}else{
		//没有找到禁用的权限。
		this.Data["json"]=map[string]interface{}{"menus":temp}
	}
	this.ServeJSON()


}
//判断用户是否有禁用的权限
func CheckUserForAction(userForActions[]models.UserAction,actionId int )(b bool)  {
	b=false
	for i:=0;i<len(userForActions) ;i++  {
		if userForActions[i].Actions.Id==actionId{
			b=true
			break
		}

	}
	return
}
//对切片中的内容进行去重
func RemoveRepeatedElement(arr []*models.ActionInfo) (newArr []*models.ActionInfo) {
	newArr = make([]*models.ActionInfo, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i].Id == arr[j].Id {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
