package routers

import (
	"ItcastCms/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"ItcastCms/models"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/Admin/UserInfo/Index",&controllers.UserInfoController{},"get:Index")
    beego.Router("/Admin/UserInfo/AddUser",&controllers.UserInfoController{},"post:AddUser")
	beego.Router("/Admin/UserInfo/GetUserInfo",&controllers.UserInfoController{},"post:GetUserInfo")
beego.Router("/Admin/UserInfo/DeleteUser",&controllers.UserInfoController{},"post:DeleteUser")
beego.Router("/Admin/UserInfo/ShowEditUser",&controllers.UserInfoController{},"post:ShowEditUser")
	beego.Router("/Admin/UserInfo/ShowSetUserRole",&controllers.UserInfoController{},"get:ShowSetUserRole")
beego.Router("/Admin/UserInfo/SetUserRole",&controllers.UserInfoController{},"post:SetUserRole")
   beego.Router("/Admin/UserInfo/ShowSetUserAction",&controllers.UserInfoController{},"get:ShowSetUserAction")
    beego.Router("/Admin/UserInfo/SetUserAction",&controllers.UserInfoController{},"post:SetUserAction")
beego.Router("/Admin/UserInfo/DeleteUserAction",&controllers.UserInfoController{},"post:DeleteUserAction")
    //-------------------------------角色管理--------------------------------------->
beego.Router("/Admin/RoleInfo/Index",&controllers.RoleInfoController{},"get:Index")
beego.Router("/Admin/RoleInfo/ShowAddRole",&controllers.RoleInfoController{},"get:ShowAddRole")
beego.Router("/Admin/RoleInfo/AddRole",&controllers.RoleInfoController{},"post:AddRole")
beego.Router("/Admin/RoleInfo/GetRoleInfo",&controllers.RoleInfoController{},"post:GetRoleInfo")
 beego.Router("/Admin/RoleInfo/ShowRoleAction",&controllers.RoleInfoController{},"get:ShowRoleAction")
beego.Router("/Admin/RoleInfo/SetRoleAction",&controllers.RoleInfoController{},"post:SetRoleAction")
//------------------------------------权限管理---------------------------------------------------------
beego.Router("/Admin/ActionInfo/Index",&controllers.ActionInfoCtroller{},"get:Index")
beego.Router("/Admin/ActionInfo/FileUp",&controllers.ActionInfoCtroller{},"post:FileUp")
beego.Router("/Admin/ActionInfo/AddAction",&controllers.ActionInfoCtroller{},"post:AddAction")
beego.Router("/Admin/ActionInfo/GetActionInfo",&controllers.ActionInfoCtroller{},"post:GetActionInfo")

//------------------------------------后台页面------------------------------------------------------
beego.Router("/Admin/Home/ShowIndex",&controllers.HomeController{},"get:ShowIndex")
beego.Router("/Admin/Home/Index",&controllers.HomeController{},"get:Index")
beego.Router("/Admin/Home/GetMenus",&controllers.HomeController{},"post:GetMenus")

 //----------------------------------新闻类别---------------------------------------
 beego.Router("/Admin/ArticelClass/Index",&controllers.ArticelClassController{},"get:Index")
beego.Router("/Admin/ArticelClass/ShowAddParent",&controllers.ArticelClassController{},"get:ShowAddParent")
    beego.Router("/Admin/ArticelClass/AddParentClass",&controllers.ArticelClassController{},"post:AddParentClass")
    beego.Router("/Admin/ArticelClass/ShowArticelClass",&controllers.ArticelClassController{},"post:ShowArticelClass")
 beego.Router("/Admin/ArticelClass/ShowAddChildClass",&controllers.ArticelClassController{},"get:ShowAddChildClass")
beego.Router("/Admin/ArticelClass/AddChildClass",&controllers.ArticelClassController{},"post:AddChildClass")
beego.Router("/Admin/ArticelClass/ShowChildClass",&controllers.ArticelClassController{},"post:ShowChildClass")

beego.Router("/Login/Index",&controllers.LoginController{},"get:Index")
beego.Router("/Login/UserLogin",&controllers.LoginController{},"post:UserLogin")


beego.InsertFilter("/Admin/*",beego.BeforeExec,FilterUserAction)
}

func FilterUserAction(ctx *context.Context)  {
	//contxt；上下文。
	//1:获取登录用户请求的Url地址，与请求的方式。
	  //根据用户请求的URL地址与请求的方式，从权限表中查询出对应的权限编号，看一下登录用户是否具有该url地址的访问权限。
	  //判断用户是否登录。
	  userName:=ctx.Input.Session("userName")//获取登录用户名。
	  if userName!=""{
	  	  if userName=="laowang"{//留后门。
	  	  	return
		  }
		  //根据用户请求的URL地址与请求的方式，从权限表中查询出对应的权限编号，看一下登录用户是否具有该url地址的访问权限。
		  //获取请求的URL地址(路由)
		 path:=ctx.Request.URL.Path
		 //获取请的方式
		 method:=ctx.Request.Method
		 //根据获取的请求的地址与请求的方式，查询权限表。
		 o:=orm.NewOrm()
		 var actionInfo models.ActionInfo
		 o.QueryTable("action_info").Filter("url",path).Filter("http_method",method).One(&actionInfo)
		 if actionInfo.Id>0{
			//判断用户是否具有找到的权限。
			//首先，先获取用户的信息。
			var userInfo models.UserInfo
			o.QueryTable("user_info").Filter("user_name",userName).One(&userInfo)
			//按照用户--权限  这条线进行权限的过滤。
			//查询用户与权限的中间表，
			var userAction models.UserAction
			o.QueryTable("user_action").Filter("users_id",userInfo.Id).Filter("actions_id",actionInfo.Id).One(&userAction)
			//判断是否具有该权限。
			if userAction.Id>0{
				//判断IsPass的取值。
				if userAction.IsPass==1{
					return
				}else{
					//找到权限了，但是发现被禁止了。
					ctx.Redirect(302,"/Login/Index")
				}
			}else{
				//没有，再按照“用户”---“角色”---“权限”过滤
				//首先查询用户角色。
				var roles []*models.RoleInfo
				o.LoadRelated(&userInfo,"Roles")
				for _,role:=range userInfo.Roles{
					roles=append(roles,role)
				}
				//查询角色对应的权限。
				var actions []*models.ActionInfo
				for i:=0;i<len(roles) ;i++  {
					o.LoadRelated(roles[i],"Actions")
					for _,action:=range roles[i].Actions {
						if action.Id==actionInfo.Id {//判断权限，是否与请求的地址对应的权限相等、
							actions=append(actions,action)
						}
					}
				}
				//判断集合actions的长度。
				if len(actions)<1{
					ctx.Redirect(302,"/Login/Index")
				}

			}
		 }else{
		 	//根据URL地址与请求的方式，没有从权限表中查询到记录。
			 ctx.Redirect(302,"/Login/Index")
		 }

	  }else{
	  	ctx.Redirect(302,"/Login/Index")
	  }



}
