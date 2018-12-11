# -CMS




CMS:（Content Management System）内容管理系统 

视图函数。

# 一:创建Beego项目

在windows系统下创建Beego项目的过程如下：

使用cmd定位到项目存放的目录(gopath指定的目录)，然后执行 “bee new 项目名称” 完成项目的创建。
接下来就可以使用GOLand打开项目。

# 二：创建数据库

我们知道ORM不会帮我们创建数据库，所以我们需要手动创建数据库。

create database 数据库名称

# 三:查看数据库

创建完数据库后，可以查看数据库。

show databases

# 四：beego项目创建

  1：项目中关于数据库配置

在这里我们将关于数据库的配置全部写到配置文件中。配置文件所在的位置为：conf/app.conf

```go
dbhost="127.0.0.1"
dbport="3306"
dbuser="root"
dbpassword="123456"
db="数据库名称"
```

2：数据库连接，在models下创建db.go,并且在该文件中定义init方法，读取配置文件中关于数据库的配置信息，完成数据库的连接，具体代码如下：

```go
func init(){
   var dbhost string
   var dbport string
   var dbuser string
   var dbpassword string
   var db string
    //获取配置文件中对应的配置信息
   dbhost = beego.AppConfig.String("dbhost")
   dbport = beego.AppConfig.String("dbport")
   dbuser = beego.AppConfig.String("dbuser")
   dbpassword = beego.AppConfig.String("dbpassword")
   db = beego.AppConfig.String("db")
   orm.RegisterDriver("mysql", orm.DRMySQL) //注册mysql Driver
   //构造conn连接
   conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db + "?charset=utf8"
   //注册数据库连接
   orm.RegisterDataBase("default", "mysql", conn)

   orm.RegisterModel(new(Userinfo))//注册模型
   orm.RunSyncdb("default", false, true)

}
```

3:注意事项



<1>在db.go文件中导入mysql数据库驱动包

```go
_ "github.com/go-sql-driver/mysql"
```

<2>在main.go文件中导入models 包

```go
_"TestProject/models"
```

4：定义UserInfo模型

```go
type Userinfo struct {
   Id int	//用户编号
   UserName string	//用户名
   UserPwd string	//用户密码
   Remark string	//备注
   AddDate time.Time	//添加日期
   ModifDate time.Time	//修改日期
   DelFlag int	//删除标记
   }
```

在创建模型的时候，大家可以根据自己的需要设置对应属性的长度等。在这里，我们先简单定义，后面我们在根据具体的业务场景，来对属性进行设置。

执行bee  run命令执行项目

# 五：用户管理

## 1:JqueryEasyui概念

**JqueryEasyui:**jQuery EasyUI是一组基于jQuery的UI插件集合体，而jQuery EasyUI的目标就是帮助web开发者更轻松的打造出功能丰富并且美观的UI界面。开发者不需要编写复杂的javascript，也不需要对css样式有深入的了解，开发者需要了解的只有一些简单的html标签 

## 2:用户列表展示

- datagrid：向用户展示列表数据。

  datagrid是JqueryEasyUi提供的一个表格，我们可以通过该表格来展示数据，并且也提供了比较强大的功能，例如，分页等。

   现在我们一起来看一下，怎样使用datagrid来展示对应的用户数据。

   第一步：我们在UserInfoController中创建一个Index方法，在该方法中指定要展示的视图。

```go
  func(this *UserInfoController) Index()  {
     this.TplName="UserInfo/Index.html"
  }
```

第二步：在UserInfo/Index.html视图中引入对应的jquery文件和jquery.easyui.min.js文件，以及对应的css文件。

```javascript
 <script type="text/javascript" src="/static/js/jquery.js"></script>
    <script type="text/javascript" src="/static/js/jquery.easyui.min.js"></script>
    <script type="text/javascript" src="/static/js/easyui-lang-zh_CN.js"></script>
    <link href="/static/css/themes/default/easyui.css" rel="stylesheet" />
    <link href="/static/css/themes/icon.css" rel="stylesheet" />
    
```

说明：1：一定要先引用jquery文件，然后才能引用对应的jquery.easyui.min.js文件。因为jquery.easyui.min.js文件用到了jquery中提供的方法。

   2：jquery.easyui.min.js  是压缩版的，为了减少文件的大小。

  3：easyui-lang-zh_CN.js 表示使用的jqueryEasyui是简体中文版的，因为jqueryEasyui是老外开发的。

 4：easyui.css: 是jqueryEasyui核心的CSS样式文件。包括颜色的设置，边框的设置等等。

 5：icon.css： 该样式用来设置图表（图片），大家可以看到datagrid中有很多的小图表。

第三步：开始使用datagrid

   1:定义一个loadData方法（该方法的名字，随便起），在该方法中首先找到一个叫“tt”的元素（该元素其实就是一个table表格，大家可以看下面一步），给“tt”这个表格元素添加了一个datagrid方法并且设置了对应的属性。

在这里大家要注意的是，这些内容不需要记忆，用的时候可以直接拷贝。



```javascript
<script type="text/javascript">
    $(function () {
        $("#addDiv").css("display","none");
        loadData()
    })
    function loadData() {
        $('#tt').datagrid({
            url: '/Admin/UserInfo/GetUserInfo',
            title: '用户数据表格',
            width: 700,
            height: 400,
            fitColumns: true, //列自适应
            nowrap: false,//设置为true，当数据长度超出列宽时将会自动截取
            idField: 'Id',//主键列的列明
            loadMsg: '正在加载用户的信息...',
            pagination: true,//是否有分页
            singleSelect: false,//是否单行选择
            pageSize:2,//页大小，一页多少条数据
            pageNumber: 1,//当前页，默认的
            pageList: [2, 5, 10],
            queryParams: {},//往后台传递参数
            columns: [[
                { field: 'ck', checkbox: true, align: 'left', width: 50 },
                { field: 'Id', title: '编号', width: 80 },
                { field: 'UserName', title: '姓名', width: 120 },
                { field: 'UserPwd', title: '密码', width: 120 },
                { field: 'Remark', title: '备注', width: 120 },
                { field: 'AddDate', title: '时间', width: 80, align: 'right',
                    formatter: function (value, row, index) {
                        return value.split('T')[0]//对日期时间的处理
                    }
                }
            ]],
            toolbar: [{
                id: 'btnDelete',
                text: '删除',//显示的文本
                iconCls: 'icon-remove', //采用的样式
                handler: function () {	//当单击按钮时执行该方法
                    //获取在表格中选中的行（getSelections：表示获取选中的行）
                    var rows = $('#tt').datagrid('getSelections');
                    if (!rows || rows.length == 0) {//判断是否选择了，如果没有选择长度为0
                        //alert("请选择要修改的商品！");
                        $.messager.alert("提醒", "请选择要删除的记录!", "error");
                        return;
                    }

                }
            }],
        });
    }
```

注意：

​     1：field中指定的名字必须与后端返回的Json中的key保持一致。

​     2:datagrid常见的属性，可以查看注释，或者文档。

​    3：field中的formatter方法用来设置对应的数据格式。value:表示单元格展示的值，row:表示一行数据。

​    	index：表示对应的行索引。

  4：toolbar表示datagrid显示的按钮

第三步：在页面中添加一个表格

这一步很关键,如果缺少这一步，页面上就不会显示任何的内容。

```html
<div>
    <table id="tt" style="width: 700px;" title="标题，可以使用代码进行初始化，也可以使用这种属性的方式" iconcls="icon-edit">
    </table>
  </div>  
```

## 3:获取用户数据

/Admin/UserInfo/GetUserInfo

<1>接收前端页面传递过来的数据

```go
//1:接收传递过来的当前页码
pageIndex,_:=strconv.Atoi(this.GetString("page"))
//2；接收传递过来的每页记录数
pageSize,_:=strconv.Atoi(this.GetString("rows"))
```

<2>实现分页查找数据

```go
//3：确定出取值的范围
start:=(pageIndex-1)*pageSize
//4；查询出指定范围的数据
o:=orm.NewOrm()
var temp []models.Userinfo
o.QueryTable("userinfo").Filter("del_flag",0).OrderBy("id").Limit(pageSize,start).All(&temp)
   //5：求出满足条件的总的记录数
count,_:=o.QueryTable("userinfo").Filter("del_flag",0).Count()
//6；返回数据
this.Data["json"]=map[string]interface{}{"rows":temp,"total":count}
this.ServeJSON()
```

注意：返回的数据中，除了包含分页的数据，还包含了总页数。同时还要注意，返回数据的格式为JSON格式，数据对应的key必须为"rows",总页数对应的key必须是"total".

在上面的查询中是按照id进行升序排序，如果要降序排序：OrderBy("-id")

## 4:添加用户数据

<1>添加表单页面

```html
<div id="addDiv">
    <form id="addForm">
        <table>
            <tr><td>用户名</td><td><input type="text" name="Uname"></td></tr>
            <tr><td>密码</td><td><input type="password" name="Upwd"></td></tr>
            <tr><td>备注</td><td><input type="text" name="Uremark"> </td></tr>
        </table>
    </form>
</div>
```

<2>将上面的创建的表单页面隐藏，当单击“添加”按钮时，弹出该表单。

```css
$("#addDiv").css("display","none");
```

注意：我们在($(function){

})

中调用上面的代码。

<3>单击添加“按钮”时，弹出添加表单对话框。

在整个表格中添加一个“添加”按钮

```javascript
,{
    id:'btnAdd',
    text:'添加',
    iconCls:'icon-add',
    handler:function () {
        ShowAddUser()
    }
}
```

先调用了display将div显示出来，然后使用了jqueryEasyui中的dialog方法弹出一个窗口。

```javascript
//展示要添加用户信息的窗口表单
function ShowAddUser() {
    $("#addDiv").css("display", "block");
    $('#addDiv').dialog({
        title: '添加用户信息',
        width: 300,
        height: 300,
        collapsible: true, //可折叠
        maximizable: true, //最大化
        resizable: true,//可缩放
        modal: true,//模态，表示只有将该窗口关闭才能修改页面中其它内容
        buttons: [{ //按钮组
            text: 'Ok',//按钮上的文字
            iconCls: 'icon-ok',
            handler: function () {
                //提交表单。
                AddUserData();
            }
        }, {
            text: 'Cancel',
            handler: function () {
                $('#addDiv').dialog('close');
            }
        }]
    });

}
```



<3>完成表单的提交

serializeArray() 方法通过序列化表单值来创建对象数组（名称和值）,返回 JSON 数据结构数据。此方法返回的是 JSON 对象而非 JSON 字符串 

 

```javascript
//完成用户信息的添加
function AddUserData() {
    var pars=$("#addForm").serializeArray();
    $.post("/Admin/UserInfo/AddUser",pars,function (data) {
        if(data.flag=="yes"){
            $('#addDiv').dialog('close');
            $.messager.alert("提示","添加成功!","info")
            $("#addForm input").val("")//清空所有的文本框。
            $("#tt").datagrid('reload')//重新加载表格
        }else{
            $.messager.alert("提示","添加失败!","error")
        }
    })
}
```

注意：通过AJAX将数据发送到服务端，服务端处理完成后，会返回处理的结果，这时会调用回调函数。

在该回调函数中，将窗口关闭，将表单中的文本框清空，并且重新加载整个表格。

<4>服务端处理（完成用户信息的保存）

接收表单中的数据，并且进行校验，然后创建一个用户对象，完成对应的属性赋值，通过orm中的Insert方法完成数据的保存。

```go
func (this *UserInfoController)AddUser()  {
   //1:接收表单传递过来的数据
   userName:=this.GetString("Uname")
   userPwd:=this.GetString("Upwd")
   userRemark:=this.GetString("Uremark")
  //2:进行表单的校验
  //3:完成数据的保存
  o:=orm.NewOrm()
  var userInfo=models.Userinfo{}
  userInfo.AddDate=time.Now()
  userInfo.DelFlag=0
  userInfo.ModifDate=time.Now()
  userInfo.Remark=userRemark
  userInfo.UserName=userName
  userInfo.UserPwd=userPwd
 _,err:=o.Insert(&userInfo)
   if err==nil{
      this.Data["json"]=map[string]interface{}{"flag":"yes"}
   }else{
      this.Data["json"]=map[string]interface{}{"flag":"no"}
   }
   this.ServeJSON()
}
```

注意：不要忘记设置路由。

## 5:查询部分数据

前面我们已经完成了数据的添加，与数据的展示，但是在数据的展示上还是有问题的，就是我们将用户表中所有列（字段）都查询出来了。但是在前端表格中展示的时候，只是展示了部分列中的内容，我们应该是展示哪些数据，就查询出哪些数据，而不是将全部列（字段）中的内容查询出来，否则造成的问题就是性能比较低。

<1>使用Values查询

```go
var pars[]orm.Params
o.QueryTable("userinfo").Filter("del_flag",0).OrderBy("-id").Limit(pageSize,start).Values(&pars)
```

<2>定义一个前端要展示的数据Model

```go
type UserModel struct {
   Id int
   UserName string
   UserPwd string
   Remark string
   AddDate time.Time
}
```

<3>循环查询出的数据，并且进行相应的类型转换（注意interface{}类型的转换方式）

转换后的数据首先赋值给UserModel中的属性，然后在追加到temp集合中。

```go
var temp []UserModel
var userInfo=UserModel{}
for _,m:=range pars{
		value,ok:=m["Id"].(int64)
		if ok{

			userInfo.Id=int(value)

		}
		value1,ok1:=m["UserName"].(string)
		if ok1{
			userInfo.UserName=value1
		}
		value2,ok2:=m["UserPwd"].(string)
		if ok2 {
			userInfo.UserPwd=value2
		}
		value3,ok3:=m["Remark"].(string)
		if ok3{
			userInfo.Remark=value3
		}
		value4,ok4:=m["AddDate"].(time.Time)
		if ok4{
			userInfo.AddDate=value4
		}

		temp=append(temp,userInfo)
	}
this.Data["json"]=map[string]interface{}{"rows":temp,"total":count}
	this.ServeJSON()
```

## 6:编辑用户数据

6.1 查询出要修改的数据

 <1>创建修改数据的表单

由于Id和AddDate数据不需要修改，所以将其放在隐藏域中。

```html
<div id="editDiv">
    <form id="editForm">
        <input type="hidden" id="txtEditId" name="Eidtid">
        <input type="hidden" id="txtEditAddDate" name="EditaddDate">
        <table>
            <tr><td>用户名:</td><td><input type="text" name="Editname" id="txtEditName"></td></tr>
            <tr><td>密码:</td><td><input type="text" name="Editpwd" id="txtEidtPwd"></td></tr>
            <tr><td>备注:</td><td><input type="text" name="Editremark" id="txtEidtRemark"></td></tr>
        </table>
    </form>
</div>
```

<2>隐藏表单

```javascript
$("#editDiv").css("display","none");
```

在表格中添加“编辑”按钮

```javascript
{
    id:'btnEdit',
    text:'编辑',
    iconCls:'icon-edit',
    handler:function () {
        ShowEditUser()//展示要编辑的数据
    }
}
```

<3>查询出要编辑的数据，并将其填充到表单中。

```javascript
//展示要编辑的数据
function ShowEditUser() {
    var rows = $('#tt').datagrid('getSelections');
    if (rows.length!=1){
        $.messager.alert("提示","编辑只能选择一条记录！","error")
        return
    }
    var id=rows[0].Id 
    $.post("/Admin/UserInfo/ShowEdit",{"id":id},function (data) {
        if (data.flag=="yes"){
            $("#editDiv").css("display","block")
            $("#txtEidtRemark").val(data.serverData.Remark)
            $("#txtEditName").val(data.serverData.UserName)
            $("#txtEidtPwd").val(data.serverData.UserPwd)
            $("#txtEditId").val(data.serverData.Id)
            $("#txtEditAddDate").val(data.serverData.AddDate)
            $('#editDiv').dialog({
                title: '编辑用户信息',
                width: 300,
                height: 300,
                collapsible: true,
                maximizable: true,
                resizable: true,
                modal: true,
                buttons: [{
                    text: 'Ok',
                    iconCls: 'icon-ok',
                    handler: function () {
                        //提交表单。
                        EditUserData();
                    }
                }, {
                    text: 'Cancel',
                    handler: function () {
                        $('#editDiv').dialog('close');
                    }
                }]
            });

        }else{
            $.messager.alert("提示",data.msg,"error")
        }
    })
}
```

注意：在这里一定要判断用户是否选择了要进行编辑的数据,而且只能选择一条。

获取用户要编辑的数据的Id,然后通过AJAX方式发送到服务端，服务端根据该Id查询出对应的要修改的数据，生成JSON格式返回，这时在AJAX回调函数中，将返回的JSON数据填充到表单中。

同时，调用JavascriptEasyui中的dialog函数，弹出一个对话框进行展示。

<4>服务端查询要编辑的数据

根据客户端传递过来的Id,查询出要进行编辑的数据。

```go
//展示要编辑的用户信息
func (this *UserInfoController)ShowEdit()  {
   id,err:=strconv.Atoi(this.GetString("id"))
   if err!=nil{
      this.Data["json"]=map[string]interface{}{"flag":"no","msg":"类型转换错误"}
   }else{
      //查询出要修改的数据
      o:=orm.NewOrm()
      var userInfo models.Userinfo
      err:=o.QueryTable("userinfo").Filter("id",id).One(&userInfo)
      if err!=nil{
         this.Data["json"]=map[string]interface{}{"flag":"no","msg":"查询失败"}
      }else{
         this.Data["json"]=map[string]interface{}{"flag":"yes","serverData":userInfo}
      }

   }
   this.ServeJSON()
}
```

<5>完成用户信息的修改。

```javascript
//完成用户的更新
function EditUserData() {
    var pars=$("#editForm").serializeArray()
    $.post("/Admin/UserInfo/EditUser",pars,function (data) {
        if(data.flag=="yes"){
            $("#editDiv").dialog('close');
            $("#tt").datagrid('reload');
            $.messager.alert("提示",data.msg,"info")

        }else{
            $.messager.alert("提示",data.msg,"error");
        }
    })
}
```

服务端代码如下：

```go
//完成用户的更新操作
func (this *UserInfoController) EditUser() {
    //接收要更新的数据，赋值给userInfo对象中的属性
   var userInfo=models.Userinfo{}
   userInfo.UserName=this.GetString("Editname")
   userInfo.UserPwd=this.GetString("Editpwd")
   userInfo.Remark=this.GetString("Editremark")
   id,_:=strconv.Atoi(this.GetString("Eidtid"))
   userInfo.Id=id
   userInfo.ModifDate=time.Now()
   addDate:=this.GetString("EditaddDate")
   t,_:=time.Parse("2006-01-02T15:04:05+08:00",addDate)//注意格式
     userInfo.AddDate=t
   o:=orm.NewOrm()
    //完成数据的更新
  _,err:=o.Update(&userInfo)
  if err!=nil{
   this.Data["json"]=map[string]interface{}{"flag":"no","msg":err}
  }else{
   this.Data["json"]=map[string]interface{}{"flag":"yes","msg":"数据更新成功!"}
  }
this.ServeJSON()//生成JSON返回
}
```

注意：时间格式的转换。

## 7:删除用户数据

<1>在表格中添加删除按钮

```javascript
{
    id: 'btnDelete',
    text: '删除',
    iconCls: 'icon-remove',
    handler: function () {
       deleteUser();

    }
}
```

<2>在deleteUser( )方法中判断是否选择了要删除的数据（注意，这里可以选择多条数据。）, 如果选择了，给出“确定要删除数据”的提示，同时获取选中数据的id,然后进行拼接（具体数据之间用逗号进行分隔）。一定要将最后的逗号截取掉。

接下来通过AJAX将获取到的数据的Id值发送到服务端。

```javascript
//删除用户数据
function deleteUser() {
    var rows = $('#tt').datagrid('getSelections');
    if (!rows || rows.length == 0) {
        $.messager.alert("提醒", "请选择要删除的记录!", "error");
        return;
    };
    $.messager.confirm("提示","确定要删除吗?",function (r) {//这里要给出，删除的确认提示
        if (r){
            var strId="";//注意:这里是定义一个字符串变量，在定义的时候一定要赋值空字符串，否则substr不起作用。
            //对所有选中的行进行遍历，获取对应的id值，在每个id值后面跟一个逗号进行分隔。
            for(var i=0;i<rows.length;i++){
                strId=strId+rows[i].Id+","
            }
            //将最后一个逗号截取掉。
         strId=strId.substr(0,strId.length-1);
            //通过ajax方式，将获取到的编号数据发送到服务端。
            $.post("/Admin/UserInfo/DeleteUser",{"id":strId},function (data) {
                //服务端处理完成后，会将结果返回到客户端，在这里进行判断，如果服务端处理成功
                //给出相应的提示，并且重新加载表格。
                if (data.flag=="ok"){
                    $.messager.alert("提示","数据删除成功!!","info");
                    $("#tt").datagrid('reload');
                    $('#tt').datagrid('clearSelections');//清除原来以前选中的数据
                }else{
                    $.messager.alert("提示","数据删除出现错误!!","error")
                }
            })

        }

    })

}
```

<3>服务端处理

服务端处理的基本思路是：接收前端发送的用户的编号数据，并且按照逗号进行分隔（因为前端的数据是使用逗号的形式链接起来的），将分隔好的用户编号转换成整型，然后根据用户编号找到对应的用户信息，进行删除。

```go
//完成用户信息的删除操作
func (this *UserInfoController)DeleteUser() {
   var ids=this.GetString("id")//接收要删除的记录的id
   strIds:=strings.Split(ids,",")//按照逗号进行分隔，返回的是字符串切片
   var list[]int
    //遍历strIds切片，将其中的每个元素取出来，转成整型，存储到list中
   for i:=0;i<len(strIds) ;i++  {
      d,_:=strconv.Atoi(strIds[i])
      list=append(list,d)
   }
   o:=orm.NewOrm()
   var userInfo models.Userinfo
   var count int64
    //遍历切片，获取每一个用户的id,然后找到该用户进行删除。
   for i:=0;i<len(list) ;i++  {
      o.QueryTable("userinfo").Filter("id",list[i]).One(&userInfo)
      c,_:= o.Delete(&userInfo)
      count+=c
   }
    if int(count)==len(list){
       this.Data["json"]=map[string]interface{}{"flag":"ok"}
   }else{
      this.Data["json"]=map[string]interface{}{"flag":"no"}
   }
   this.ServeJSON()

}
```

## 8:搜索用户数据

<1>根据用户名和备注进行搜索

在页面中添加相应的表单元素

```html
用户名:<input type="text" id="txtSearchUName" />&nbsp;&nbsp;
备注:<input type="text" id="txtSearchRemark" />
<a href="#" class="easyui-linkbutton" data-options="iconCls:'icon-search'" style="width:80px" id="btnSearch">Search</a>
```

<2>将搜索按钮绑定单击事件，当触发该事件后，获取搜索条件，然后发送到服务端。

在$(function(){

})

中添加如下的代码：

```javascript
//给搜索按钮添加单击事件
$("#btnSearch").click(function () {
    var pars={
        username:$("#txtSearchUName").val(),
        remark:$("#txtSearchRemark").val()
    }
    loadData(pars)//调用加载数据的方法，传送搜索的参数。
})
```

<3>loadData方法的修改

第一：数的修改，需要给loadData方法添加参数 

第二：给datagrid中的queryParams属性赋值：  queryParams: pars,//往后台传递参数

<4>服务端处理

接收前端传递过来的数据，在服务端构建搜索条件。

在这里我们创建了一个结构体对象，专门用来组件搜索条件。

```go
//搜索的数据
type UserSearchData struct {
   UserName string
   Remark string
   PageIndex int
   PageSize int
   TotalCount int64
}
```

```go
func (this *UserInfoController)GetUserInfo()  {
   pageIndex,_:=strconv.Atoi(this.GetString("page"))
   pageSize,_:=strconv.Atoi(this.GetString("rows"))
   username:=this.GetString("username")
   remark:=this.GetString("remark")
    var userSearch=UserSearchData{}//通过该对象构建相应的搜索数据。
    userSearch.UserName=username
    userSearch.Remark=remark
    userSearch.PageIndex=pageIndex
    userSearch.PageSize=pageSize
   temp:=userSearch.GetSearchData(userSearch)//负责具体数据的搜索,这里传递的是对象，注意调用的方式
   this.Data["json"]=map[string]interface{}{"rows":temp,"total":userSearch.TotalCount}
   this.ServeJSON()
}
```

<5>GetSearchData方法的实现

```go
//构建搜索数据的方法
func (this *UserSearchData)GetSearchData(userSearchData UserSearchData)([]models.Userinfo)  {
   o:=orm.NewOrm()
    //构建搜索的条件
   var temp=o.QueryTable("userinfo")
   if userSearchData.UserName!=""{
      temp=temp.Filter("user_name__icontains",userSearchData.UserName)
   }
   if userSearchData.Remark!="" {
      temp=temp.Filter("user_pwd__icontains",userSearchData.Remark)
   }
   start:=(userSearchData.PageIndex-1)*userSearchData.PageSize
   count,_:=temp.Count()
   this.TotalCount=count//注意，总的记录数的计算，这里是this,表示的是*UserSearchData,也就是引用传递。
   var users []models.Userinfo//构建UserInfo对象
   temp.OrderBy("Id").Limit(userSearchData.PageSize,start).All(&users)
   return users//返回搜索查询的数据
}
```

# 六：角色管理

## 1:模型创建

在这里大家首先要思考的是用户与角色是什么关系？



```go
//用户信息
type Userinfo struct {
   Id int
   UserName string
   UserPwd string
   Remark string
   AddDate time.Time
   ModifDate time.Time
   DelFlag int
   Roles []*RoleInfo`orm:"rel(m2m)"`
  

}
```





```go
//角色信息
type RoleInfo struct {
   Id int
   RoleName string `orm:"size(32)"`
   Remark string
   DelFlag int
   AddDate time.Time
   ModifDate time.Time
   Users []*Userinfo`orm:"reverse(many)"`
}
```



## 2:角色信息展示

基本的流程与用户信息展示流程一致。

```javascript
$(function () {
    $("#addDiv").css("display","none")
    loadData()
})
function loadData() {
    $('#tt').datagrid({
        url: '/Admin/RoleInfo/GetRoleInfo',
        title: '角色数据表格',
        width: 700,
        height: 400,
        fitColumns: true, //列自适应
        nowrap: false,
        idField: 'Id',//主键列的列明
        loadMsg: '正在加载角色的信息...',
        pagination: true,//是否有分页
        singleSelect: false,//是否单行选择
        pageSize: 5,//页大小，一页多少条数据
        pageNumber: 1,//当前页，默认的
        pageList: [5, 10, 15],
        queryParams: {},//往后台传递参数
        columns: [[//c.UserName, c.UserPass, c.Email, c.RegTime
            { field: 'ck', checkbox: true, align: 'left', width: 50 },
            { field: 'Id', title: '编号', width: 80 },
            { field: 'RoleName', title: '角色', width: 120 },

            { field: 'Remark', title: '备注', width: 120 },
            {
                field: 'AddDate', title: '时间', width: 80, align: 'right',
                formatter: function (value, row, index) {
                    return  value.split("T")[0]
                }
            }
        ]],
        toolbar: [{
            id: 'btnDelete',
            text: '删除',
            iconCls: 'icon-remove',
            handler: function () {


            }
        }, {
            id: 'btnAdd',
            text: '添加',
            iconCls: 'icon-add',
            handler: function () {

                ShowAddRole();
            }
        }, {
            id: 'btnEdit',
            text: '编辑',
            iconCls: 'icon-edit',
            handler: function () {


            }
        }, {
            id: 'btnSetRoleAction',
            text: '为角色分配权限',
            iconCls: 'icon-edit',
            handler: function () {
               // setRoleAction();

            }
        }],
    });
}

```



```html
<table id="tt" style="width: 700px;" title="标题，可以使用代码进行初始化，也可以使用这种属性的方式" iconcls="icon-edit">
</table>
```



服务端代码的实现如下：

```go
//获取角色信息
func (this *RoleInfoController)GetRoleInfo()  {
   pageIndex,_:=strconv.Atoi(this.GetString("page"))
   pageSize,_:=strconv.Atoi(this.GetString("rows"))
   start:=(pageIndex-1)*pageSize
   o:=orm.NewOrm()
   var roles []models.RoleInfo
   o.QueryTable("role_info").Filter("del_flag",0).OrderBy("Id").Limit(pageSize,start).All(&roles)
   count,_:=o.QueryTable("role_info").Filter("del_flag",0).Count()
   this.Data["json"]=map[string]interface{}{"rows":roles,"total":count}
   this.ServeJSON()
}
```

## 3:角色添加

在前面，我们在实现用户的添加，编辑等操作时，直接将表单放在一个页面中，大家可以思考一下，是否感觉页面内容比较多，比较乱。所以这里，我们将角色的添加页面，单独的放在一个页面中。

但是问题时，单击角色页面表格上面的添加按钮时，怎样显示角色添加的表单页面呢？

这里我们需要用到iframe标签。

<1>将添加的表单，放在另外一个页面中进行展示,所以这里我们创建一个iframe.

```javascript
<div id="addDiv">
    <iframe id="addFrame" frameborder="0" width="100%" height="100%"></iframe>
</div>
```

<2>为iframe指定对应的url地址，并且调用了dialog( )方法，弹出一个窗口，显示对应的iframe指定的表单添加页面。

这里大家要注意，我们是在角色管理页面中（这里可以将该页面称之为主页面），通过dialog( )方法弹出一个窗口显示了另外一个添加页面（这个页面称之为子页面），但是问题是，当我们单击角色管理页面中的“ok”按钮时，怎样完成对子页面表单中填写的数据进行提交呢？

```javascript
//展示出添加表单
function ShowAddRole() {
    $("#addFrame").attr("src","/Admin/RoleInfo/ShowAddRole");//为iframe指定对应的url地址
    $("#addDiv").css("display","block");
    $('#addDiv').dialog({
        title: '添加角色信息',
        width: 300,
        height: 300,
        collapsible: true,
        maximizable: true,
        resizable: true,
        modal: true,
        buttons: [{
            text: 'Ok',
            iconCls: 'icon-ok',
            handler: function () {
                //提交表单。
               var childWindows=$("#addFrame")[0].contentWindow;//将jquery对象转换成javascript对象，获取子窗体的windows对象。

                childWindows.SubForm()//提交表单,调用子页面中的SubForm()方法

            }
        }, {
            text: 'Cancel',
            handler: function () {
                $('#addDiv').dialog('close');
            }
        }]
    });

}
```

<3>添加子页面的创建

子页面中是一个表单，怎样完成对表单中所有数据的提交呢？这里，我们使用了另外一个jquery的插件jquery.unobtrusive-ajax。

但是需要首先给form表单添加如下的属性

data-ajax="true":表示对表单进行ajax提交

data-ajax-method="post"：表示以POST方式进行提交。

data-ajax-success="afterAdd"：处理成功后执行afterAdd方法。其实就是回调函数。

data-ajax-url:指定提交的地址。

```html
<form   data-ajax="true" data-ajax-method="post" data-ajax-success="afterAdd" data-ajax-url="/RoleInfo/AddRole" id="addForm">
<table>
    <tr><td>角色名称</td><td><input type="text" name="roleName"></td></tr>
<tr><td>备注</td><td><input type="text" name="remark"></td></tr>
</table>
</form>
```

注意：这里需要引入：

```javascript
<script src="/static/js/jquery.unobtrusive-ajax.min.js" type="text/javascript"></script>
```

提交表单的方法：

```javascript
<script type="text/javascript">
        function SubForm() {
            $("#addForm").submit()//提交表单
        }
        function afterAdd(data) {//数据添加成功后调用该方法
            window.parent.afterAdd(data)//调用主窗口中的afterAdd方法。
        }
</script>
```

<4>服务端保存角色信息。

```go
//完成角色信息的保存
func (this *RoleInfoController)AddRole()  {
   roleName:=this.GetString("roleName")
   remark:=this.GetString("remark")
   var roleInfo=models.RoleInfo{}
   roleInfo.Remark=remark
   roleInfo.AddDate=time.Now()
   roleInfo.ModifDate=time.Now()
   roleInfo.DelFlag=0
   roleInfo.RoleName=roleName
   o:=orm.NewOrm()
   _,err:=o.Insert(&roleInfo)
   if err==nil {
      this.Data["json"]=map[string]interface{}{"flag":"ok"}
   }else{
      this.Data["josn"]=map[string]interface{}{"flag":"no"}
   }
   this.ServeJSON()
}
```

<5>主窗体中定义的afterAdd方法

```javascript
//完成添加后调用的方法
function afterAdd(data) {
    if(data.flag=="ok"){
        $('#addDiv').dialog('close');
        $.messager.alert("提示","角色添加完成!!","info")
        $('#tt').datagrid('reload')
    }else{
        $.messager.alert("提示","角色添加失败!!","error")
    }

}
```

# 七：为用户分配角色



在前面我们已经分析了用户与角色的关系，为多对多的关系，下面我们看一下怎样给用户分配角色。



## 1:前端处理

在这里我们也是通过一个iframe,来指定一个单独的页面。

<1>在/Admin/UserInfo/Index.html视图中，添加一个iframe,指定角色分配页面。

```html
<div id="setRoleDiv">
    <iframe id="setUserRoleFrame" frameborder="0" width="100%" height="100%"></iframe>
</div>
```

在指定地址时，判断一下是否选中某个用户，然后将选中的用户的id传递到服务端。

```javascript
//为用户分配角色
function SetUserRole() {
    //判断一下是否选择了要分配角色的用户
    var rows = $('#tt').datagrid('getSelections');
    if(rows.length!=1){
        $.messager.alert("提示","请选择要分配角色的用户!!","error");
        return;
    }
    $("#setUserRoleFrame").attr("src","/Admin/UserInfo/ShowUserRole?userId="+rows[0].Id);
    $("#setRoleDiv").css("display","block");
    $('#setRoleDiv').dialog({
        title: '为用户分配角色',
        width: 300,
        height: 300,
        collapsible: true,
        maximizable: true,
        resizable: true,
        modal: true,
        buttons: [{
            text: 'Ok',
            iconCls: 'icon-ok',
            handler: function () {
                //提交表单。
                AddUserData();
            }
        }, {
            text: 'Cancel',
            handler: function () {
                $('#setRoleDiv').dialog('close');
            }
        }]
    });
}
```

## 2:展示用户已有角色信息

在完成用户角色分配前，先将所有的角色展示出来，并且在每个角色名称前面加上一个复选框，如果用户有这个角色，将该复选框选中，如果没有则不选中。

在/Admin/UserInfo/ShowUserRole控制器中，接收用户的id,根据接收到的用户id,查询出对应的用户，并且根据查询出的用户查询出该用户具有的角色信息。同时，查询出所有的角色信息。

```go
func (this *UserInfoController)ShowUserRole()  {
//1:查询出对应的用户
   userId,_:=strconv.Atoi(this.GetString("userId"))
   o:=orm.NewOrm()
   var userInfo=models.Userinfo{}
   o.QueryTable("userinfo").Filter("Id",userId).One(&userInfo)
//2:根据查询出的用户，找到对应已有的角色信息。
   o.LoadRelated(&userInfo,"Roles")
   var roles []*models.RoleInfo//注意这里定义的时候，需要加*
   for _,role:=range userInfo.Roles{
      roles=append(roles,role)
   }
//3:查询出所有的角色信息。
   var allRoleList []models.RoleInfo
   o.QueryTable("role_info").Filter("del_flag",0).All(&allRoleList)
//4：展示出对应的信息
   this.Data["userInfo"]=userInfo
   this.Data["userExtRoles"]=roles
   this.Data["allRoles"]=allRoleList
   this.TplName="UserInfo/ShowSetUserRole.html"
}
```

## 3:模板处理

在UserInfo/ShowSetUserRole.html模板

在该模板文件中，输出所有的角色，并且将用户已经有的角色选中。

判断的方式是：对所有的角色进行遍历，每取出一个角色的编号，看一下用户是否具有该角色编号。（这个判断的过程，我们封装到了视图函数checkId中。）

```javascript
为用户<span style="font-size: 14px; color: red; font-weight: bold">{{.userInfo.UserName}}</span>分配角色
<form  data-ajax="true" data-ajax-method="post" data-ajax-success="afterSet" data-ajax-url="/Admin/UserInfo/SetUserRole" id="setRoleForm">
    <input type="hidden" name="userId" value="{{.userInfo.Id}}" />

    {{range .allRoles}}
         {{if checkId $.userExtRoles .Id}}
             <input type="checkbox" name="cba_{{.Id}}"  checked="checked" value="{{.Id}}">{{.RoleName}}
         {{else}}
             <input type="checkbox" name="cba_{{.Id}}" value="{{.Id}}">{{.RoleName}}
         {{end}}
    {{end}}
</form>
```

在该模板中，通过range循环输出所有的角色。

checkId为模板函数，将$.userExtRoles 集合和 .Id参数传递到该模板函数中。

```
注意$符号的作用：在一个循环中，我们循环一个切片
只能展示出该切片中的数据，但是想展示另外的数据，需要在对应的前面加上$符号

```

如果用户具有了某个角色，将其前面的复选框选中。

<2>模板函数处理

在main.go文件中定义如下的模板函数，完成用户角色的判断。

```go
func CheckId(userExitRoleList []*models.RoleInfo,roleId int)(b bool)  {
   b=false
   for  i:=0;i<len(userExitRoleList);i++  {
      if userExitRoleList[i].Id==roleId {
            b=true
            break
      }
   }
   return
}
```

注意：该函数的参数需要加上*

该模板函数创建出了以后，还需要注册一下。

```go
func main() {
   beego.AddFuncMap("checkId",CheckId)
   beego.Run()
}
```

AddFuncMap( )方法指定的是模板函数的别名，在模板视图中调用的时候是通过该别名调用的。

第二个参数就是具体的函数名称，但是这里不能加括号。（因为：这里不是调用函数）

注意：AddFuncMap( )函数一定要定义在Run( )方法前面

## 4:完成用户角色分配

<1>在在UserInfo/ShowSetUserRole.html模板中，添加一个提交表单的方法如下：

```javascript
function subForm() {
    $("#setRoleForm").submit()
}
```

注意:该方法的调用是在主窗口中完成的。

<2>服务端处理

实现为用户分配角色的基本思路：

接收所有被选中的角色的id值（所有选中的复选框），

但是这里有一个问题是大家需要注意的，就是获取的值，有可能是该用户以前就有的角色编号，所以我们这里采用的方法时先删除用户所有的角色，然后再重新添加。

```go
//完成用户角色的分配
func (this *UserInfoController)SetUserRole()  {
   userId,_:=this.GetInt("userId")//接收用户的编号
   var roleIdList[]int
   allKeys:=this.Ctx.Request.PostForm//获取所有post请求发送过来的表单数据
    //但是，这里我们只需要以cba_开头的数据
   for key,_:=range allKeys {
      if strings.Contains(key,"cba_"){
         id:=strings.Replace(key,"cba_","",-1)
         roleId,_:=strconv.Atoi(id)
         roleIdList=append(roleIdList,roleId)
      }
   }

   //查询用户信息
   o:=orm.NewOrm()
   var userInfo models.Userinfo
   o.QueryTable("userinfo").Filter("Id",userId).One(&userInfo)
   //查询用户具有的角色信息
   o.LoadRelated(&userInfo,"Roles")
   m2m:=o.QueryM2M(&userInfo,"Roles")
   //删除用户已经有的角色信息
   o.Begin()//开启事务
   var err1 error
   var err2 error
   for _,role:=range userInfo.Roles{
      _,err1=m2m.Remove(role)
   }
   //重新给用户分配角色信息
   var roleInfo models.RoleInfo
   for i:=0;i<len(roleIdList) ;i++  {
      o.QueryTable("role_info").Filter("Id",roleIdList[i]).One(&roleInfo)
      _,err2=m2m.Add(roleInfo)
   }
   if err2!=nil||err1!=nil{
      o.Rollback()
      this.Data["json"]=map[string]interface{}{"flag":"no"}
   }else{

      o.Commit()
      this.Data["json"]=map[string]interface{}{"flag":"yes"}
   }
   this.ServeJSON()

}
```

服务端处理完后，前端调用完成后的方法：

```javascript
function afterSet(data) {
    window.parent.afterSet(data)//调用主窗口中的afterSet方法
}
```

主窗口中afterSet方法的定义如下：

```javascript
//分配完角色后调用该方法
function afterSet(data) {
    if (data.flag=="yes"){
        $('#setRoleDiv').dialog('close');
        $.messager.alert("提示","用户角色分配成功!!","info");

    }else{
        $.messager.alert("提示","角色分配失败!!","error");
    }
}
```

# 八:权限管理

数据模型的创建：

思考：角色与权限之间的关系？

```go
type RoleInfo struct {
   Id int
   RoleName string `orm:"size(32)"`
   Remark string
   DelFlag int
   AddDate time.Time
   ModifDate time.Time
   Users []*UserInfo`orm:"reverse(many)"`
   Actions []*ActionInfo`orm:"rel(m2m)"`

}
```



```go
//权限信息
type ActionInfo struct {
   Id int
   Remark string
   DelFlag int
   AddDate time.Time
   ModifDate time.Time
   Url string
   HttpMethod string
   ActionInfoName string
   ActionTypeEnum int //权限类型。
   MenuIcon string //图片地址
   IconWidth int
   IconHeight int
   Roles[]*RoleInfo `orm:"reverse(many)"`
   
}
```

## 1:展示权限的信息

```javascript
$(function () {
    loadData()
})
```

```javascript
function loadData() {
    $('#tt').datagrid({
        url: '/Admin/ActionInfo/GetActionInfo',
        title: '权限数据表格',
        width: 700,
        height: 400,
        fitColumns: true, //列自适应
        nowrap: false,
        idField: 'Id',//主键列的列明
        loadMsg: '正在加载权限的信息...',
        pagination: true,//是否有分页
        singleSelect: false,//是否单行选择
        pageSize:5,//页大小，一页多少条数据
        pageNumber: 1,//当前页，默认的
        pageList: [2, 5, 10],
        queryParams: {},//往后台传递参数
        columns: [[//c.UserName, c.UserPass, c.Email, c.RegTime
            { field: 'ck', checkbox: true, align: 'left', width: 50 },
            { field: 'Id', title: '编号', width: 80 },
            { field: 'ActionInfoName', title: '权限名称', width: 120 },
            { field: 'HttpMethod', title: '请求方式', width: 120 },
            { field: 'Url', title: '请求地址', width: 120 },
            { field: 'Remark', title: '备注', width: 120 },
            { field: 'ActionTypeEnum', title: '权限类型', width: 120,
                formatter:function (value,row,index) {
                    return value=="1"?"菜单权限":"普通权限"
                }
            },
            { field: 'AddDate', title: '时间', width: 120, align: 'right',
                formatter: function (value, row, index) {
                   return value.split('T')[0]
                }
            }
        ]],
        toolbar: [{
            id: 'btnDelete',
            text: '删除',
            iconCls: 'icon-remove',
            handler: function () {


            }
        },{
            id:'btnAdd',
            text:'添加',
            iconCls:'icon-add',
            handler:function () {
                showAddAction();
            }
        },{
            id:'btnEdit',
            text:'编辑',
            iconCls:'icon-edit',
            handler:function () {

            }
        }],
    });
}

<body>
<table id="tt" style="width: 700px;" title="标题，可以使用代码进行初始化，也可以使用这种属性的方式" iconcls="icon-edit">
</table>
</body>
```

## 2:服务端处理

/Admin/ActionInfo/GetActionInfo

处理的过程与角色处理基本一致，查询出对应的数据，并且进行分页。

```go
//获取权限信息
func (this *ActionInfoController)GetActionInfo()  {
   pageIndex,_:=strconv.Atoi(this.GetString("page"))
   pageSize,_:=strconv.Atoi(this.GetString("rows"))
   start:=(pageIndex-1)*pageSize
   o:=orm.NewOrm()
   var actions []models.ActionInfo
   o.QueryTable("action_info").Filter("del_flag",0).OrderBy("id").Limit(pageSize,start).All(&actions)
   count,_:=o.QueryTable("action_info").Filter("del_flag",0).Count()
   this.Data["json"]=map[string]interface{}{"rows":actions,"total":count}
   this.ServeJSON()
}
```

## 3:上传图片文件

权限分为菜单权限和普通权限，只有添加“菜单权限”时，才允许上传图片。

<1>前端表单展示

注意：这里需要添加如下的文件引用

```javascript
<script src="/static/js/jquery.unobtrusive-ajax.min.js" type="text/javascript"></script>
```

```javascript
<div id="addActionDiv">
    <form  data-ajax="true" data-ajax-method="post" data-ajax-success="afterAdd" data-ajax-url="/Admin/ActionInfo/AddAction" id="form1">
        <input type="hidden" name="MenuIcon" id="hiddenMenuIcon" />
        <table>
            <tr>
                <td>权限名称</td>
                <td>
                    <input type="text" name="ActionInfoName" /></td>
            </tr>
            <tr>
                <td>Url</td>
                <td>
                    <input type="text" name="Url" /></td>
            </tr>
            <tr>
                <td>请求方式</td>
                <td>
                    <select name="HttpMethod">
                        <option value="GET">GET</option>
                        <option value="POST">POST</option>
                    </select>

                </td>
            </tr>

            <tr>
                <td>权限类型</td>
                <td>
                    <select name="ActionTypeEnum" id="changeActionTypeEnum">
                        <option value="0">普通权限</option>
                        <option value="1">菜单权限</option>
                    </select>

                </td>
            </tr>

            <tr style="display:none" id="iconTr">
                <td>上传图标</td>
                <td>
                    <input type="file" name="fileUp" />
                    <input type="button" value="上传图片" id="btnFileUp" />
                    <div id="showImage"></div>

                </td>
            </tr>


            <tr>
                <td>备注</td>
                <td>
                    <input type="text" name="Remark" /></td>
            </tr>

        </table>


    </form>
</div>
```

将该表单进行隐藏

```javascript
$(function () {
    $("#addActionDiv").css("display","none");
    loadData()
})
```

在单击”添加“按钮时，展示出”添加表单“

```javascript
//展示权限添加的表单
function showAddAction() {
    $("#addActionDiv").css("display","block");
    $('#addActionDiv').dialog({
        title: '添加权限信息',
        width: 600,
        height: 600,
        collapsible: true,
        maximizable: true,
        resizable: true,
        modal: true,
        buttons: [{
            text: 'Ok',
            iconCls: 'icon-ok',
            handler: function () {
                //提交表单。
                $("#form1").submit();

            }
        }, {
            text: 'Cancel',
            handler: function () {
                $('#addActionDiv').dialog('close');
            }
        }]
    });
}
```

<2>给权限类别绑定下拉框绑定改变事件，当用户选中”菜单权限“时，展示出上传表单。

```javascript
//绑定权限的类别
function bindChangeActionTypeEnum() {
    $("#changeActionTypeEnum").change(function () {
        if($(this).val()=="1"){//注意$(this):需要的是jquery对象
            $("#iconTr").fadeIn();
        }else{
            $("#iconTr").fadeOut();
        }
    })
}
```

注意在$(function{

})中需要调用一下

```javascript
$(function () {
    $("#addActionDiv").css("display","none");
    bindChangeActionTypeEnum();//绑定权限类别对话框
    loadData()
})
```

<3>当用户选择了“菜单权限”后，会出现“上传文件”按钮

需要给该按钮，绑定“单击”事件。

```javascript
$(function () {
    $("#addActionDiv").css("display","none");
    bindChangeActionTypeEnum();//绑定权限类别对话框
    bindBtnFileUp();//文件上传
    loadData()
})
```

bindBtnFileUp()方法的实现如下：

上传成功后，会执行回调函数sucess,在该函数中，接收从服务端返回的图片路径，并且创建一个img标签，添加到div中，将其显示出来。（这里需要注意的是，从服务端返回的图片路径是带有“.”,所以这里需要截取一下）。

同时将图片路径存储在隐藏域中，当单击“添加”按钮时，提交到服务端，在服务端将图片路径存储在数据库中。

```javascript
//绑定上传按钮的单击事件
function bindBtnFileUp() {
    $("#btnFileUp").click(function () {
        $("#form1").ajaxSubmit({
            success: function (str) {
                if (str.flag=="ok"){
                   $("#showImage").html("<img src='"+str.msg.substr(1)+"' width='50px' height='50px'>")
                    $("#hiddenMenuIcon").val(str.msg.substr(1))//注意服务端返回的内容是带“.”的

                }else{
                    alert(str.msg)
                }
            },
            error: function (error) { alert(error); },
            url: '/Admin/ActionInfo/FileUp', /*设置post提交到的页面*/
            type: "post", /*设置表单以post方法提交*/
            dataType: "json" /*设置返回值类型为文本*/
        });
    })
}
```

注意：这里需要引入，如下文件

```javascript
<script type="text/javascript" src="/static/js/MyAjaxForm.js"></script>
```

<4>服务端处理

/Admin/ActionInfo/FileUp

完成文件上传。

文件上传注意的问题如下：

1：对文件类型的判断，这里我们只允许上传图片文件

2：上传文件大小的限制

3：将上传的文件存放在不同的目录下，如果将所有的文件都存储在同一个目录下，大家可以想象一下，会出现什么问题？打开目录，查询文件都会非常慢。我们这里是根据日期，建立不同的文件夹。

在创建文件夹时，要判断一下，如果有了就不用建。如果没有，就创建

4：一定要对上传的文件进行重名。



```go
//文件上传
func(this *ActionInfoController)FileUp(){
   f,h,err:=this.GetFile("fileUp")
    if err!=nil{
      this.Data["json"]=map[string]interface{}{"flag":"no","msg":"上传文件错误"}
    }else{
      //获取上传的文件名称
      fileName:=h.Filename
      //获取上传文件的类型
      fileExt:=path.Ext(fileName)
      if fileExt==".jpg"||fileExt==".png"{
         if h.Size<50000000 {//判断上传文件的大小
         //创建上传图片文件存放的路径。
            dirPath:="./static/fileUp/"+strconv.Itoa(time.Now().Year())+"/"+time.Now().Month().String()+"/"+strconv.Itoa(time.Now().Day())+"/"
            _,err:=os.Stat(dirPath)
            if err!=nil{//表示没有目录信息
               dirError:=os.MkdirAll(dirPath,os.ModePerm)//创建目录
               if dirError!=nil{
                  this.Data["json"]=map[string]interface{}{"flag":"no","msg":"目录创建失败!!"}
                  return
               }
            }
            //按照日期时间对文件进行重命名。
            fileNewName:=strconv.Itoa(time.Now().Year())+time.Now().Month().String()+strconv.Itoa(time.Now().Day())+strconv.Itoa(time.Now().Hour())+strconv.Itoa(time.Now().Minute())+strconv.Itoa(time.Now().Nanosecond())//获取毫秒数
            fullDir:=dirPath+fileNewName+fileExt//构建完整的路径
            fileErr:=this.SaveToFile("fileUp",fullDir)//进行文件的保存
            if fileErr==nil{
               this.Data["json"]=map[string]interface{}{"flag":"ok","msg":fullDir}


            }else{
               this.Data["json"]=map[string]interface{}{"flag":"no","msg":"文件上传失败！！"}
            }

         }else{
            this.Data["json"]=map[string]interface{}{"flag":"no","msg":"文件太大"}
         }

      }else{
         this.Data["json"]=map[string]interface{}{"flag":"no","msg":"文件类型错误!!"}
      }
    }
   defer f.Close()
   this.ServeJSON()

}
```

## 4:完成权限信息保存

```go
/Admin/ActionInfo/AddAction
//代码如下：
//完成权限的保存
func (this *ActionInfoController)AddAction()  {
	var actionInfo=models.ActionInfo{}
	actionInfo.DelFlag=0
	actionInfo.ModifDate=time.Now()
	actionInfo.AddDate=time.Now()
	actionInfo.Remark=this.GetString("Remark")
	actionInfo.MenuIcon=this.GetString("MenuIcon")
	actionInfo.Url=this.GetString("Url")
	actionInfo.ActionInfoName=this.GetString("ActionInfoName")
	actionInfo.ActionTypeEnum,_=strconv.Atoi(this.GetString("ActionTypeEnum"))
	actionInfo.IconWidth=0
	actionInfo.IconHeight=0
	actionInfo.HttpMethod=this.GetString("HttpMethod")
	o:=orm.NewOrm()
	_,err:=o.Insert(&actionInfo)
	if err==nil {
		this.Data["json"]=map[string]interface{}{"flag":"ok"}
	}else{
		this.Data["json"]=map[string]interface{}{"flag":"no"}
	}
	this.ServeJSON()
}
```

完成后执行前端的方法：

```javascript
//添加权限完成后调用该方法
function afterAdd(data) {
    if(data.flag=="ok"){
        $.messager.alert("提示","添加成功","info");
        $('#addActionDiv').dialog('close');
        $('#tt').datagrid('reload');//重新加载表格中的数据
    }else{
        $.messager.alert("提示","添加失败","error");
    }
}
```

# 九：为角色分配权限

这里处理的过程与为用户分配角色一样。

## 1:前端处理

```javascript
{
    id: 'btnSetRoleAction',
    text: '为角色分配权限',
    iconCls: 'icon-edit',
    handler: function () {
       setRoleAction();

    }
}
```

setRoleAction( )方法实现如下：

```javascript
//为角色分配权限
function  setRoleAction() {
    var rows = $('#tt').datagrid('getSelections');
    if(rows.length!=1){
        $.messager.alert("提示","只能选择1个角色进行权限分配!!","error");
        return;
    }
    var roleId=rows[0].Id;
    $("#setFrame").attr("src","/Admin/RoleInfo/ShowSetRoleAction?roleId="+roleId);
    $("#setRoleActionDiv").css("display","block");
    $('#setRoleActionDiv').dialog({
        title: '为角色分配权限信息',
        width: 300,
        height: 300,
        collapsible: true,
        maximizable: true,
        resizable: true,
        modal: true,
        buttons: [{
            text: 'Ok',
            iconCls: 'icon-ok',
            handler: function () {
                //提交表单。
               var childWindows=$("#setFrame")[0].contentWindow;//获取子窗体的windows对象。
                childWindows.subForm()//提交表单

            }
        }, {
            text: 'Cancel',
            handler: function () {
                $('#setRoleActionDiv').dialog('close');
            }
        }]
    });
}
```

完成后调用如下方法：

```javascript
//给角色分配完成权限后调用该方法
function afterSet(data) {
    if(data.flag=="yes"){
        $.messager.alert("提示","为角色分配权限成功!!","info")
        $('#setRoleActionDiv').dialog('close');
    }else{
        $.messager.alert("提示","为角色分配权限失败!!","error")
    }
}
```

在这里需要不要忘记，添加一个iframe标签

```html
<div id="setRoleActionDiv">
<iframe id="setFrame" frameborder="0" width="100%" height="100%"></iframe>
</div>
```

## 2:服务端处理

```
/Admin/RoleInfo/ShowSetRoleAction?roleId="+roleId

```

```go
//为角色分配权限信息
func (this *RoleInfoController)ShowSetRoleAction()  {
   //1:获取角色编号
   roleId,_:=strconv.Atoi(this.GetString("roleId"))
   //2:根据该角色编号查询出对应的角色信息
   o:=orm.NewOrm()
   var roleInfo models.RoleInfo
   o.QueryTable("role_info").Filter("Id",roleId).One(&roleInfo)
   //3:查询该角色已经有的权限信息
   var actions []*models.ActionInfo
   o.LoadRelated(&roleInfo,"Actions")
   for _,action:=range roleInfo.Actions {
      actions=append(actions,action)
   }  
   //4:查询出所有的权限信息
   var allActionList []models.ActionInfo
   o.QueryTable("action_info").Filter("del_flag",0).All(&allActionList)
   //5:展示出对应的数据
   this.Data["roleInfo"]=roleInfo
   this.Data["roleExtActions"]=actions
   this.Data["allActionList"]=allActionList
   this.TplName="RoleInfo/ShowSetRoleAction.html"
}
```

对应的前端展示如下：

```javascript
<script type="text/javascript">
    function subForm() {
        $("#setActionForm").submit();
    }
    function afterSet(data) {
        window.parent.afterSet(data)
    }
</script>
```

```html
为角色<span style="font-size:14px;color:red;font-weight: bolder">{{.roleInfo.RoleName}}</span>分配权限
<form data-ajax="true" data-ajax-method="post" data-ajax-success="afterSet" data-ajax-url="/Admin/RoleInfo/SetRoleAction" id="setActionForm">
    <input type="hidden" name="roleId" value="{{.roleInfo.Id}}">
    {{range .allActionList}}
        {{if checkAction $.roleExtActions .Id}}
            <input type="checkbox" name="cba_{{.Id}}" checked="checked" value="{{.Id}}">{{.ActionInfoName}}
            {{else}}
            <input type="checkbox" name="cba_{{.Id}}"  value="{{.Id}}">{{.ActionInfoName}}
        {{end}}
    {{end}}
</form>
```

这里需要在main.go文件中定义视图函数

```go
//判断权限

func  CheckAction(userExtActionList[]*models.ActionInfo,actionId int)(b bool)  {
   b=false
   for i:=0;i<len(userExtActionList) ;i++  {
      if userExtActionList[i].Id==actionId {
         b=true
         break
      }
   }
   return
}
```

同时，完成视图函数的添加。

```go
func main() {
   beego.AddFuncMap("checkId",CheckId)
   beego.AddFuncMap("checkAction",CheckAction)
   beego.Run()
}
```

## 3:完成角色权限的分配

```go
//完成给角色分配权限信息
func(this *RoleInfoController) SetRoleAction()  {
   //获取角色编号
   roleId,_:=strconv.Atoi(this.GetString("roleId"))
   allKeys:=this.Ctx.Request.PostForm
   var actionIdList[]int
   for key,_:= range allKeys{
      if strings.Contains(key,"cba_") {
         id:=strings.Replace(key,"cba_","",-1)
         actionId,_:=strconv.Atoi(id)
         actionIdList=append(actionIdList,actionId)
      }
   }
   //查询角色的信息
   var roleInfo models.RoleInfo
   o:=orm.NewOrm()
   o.QueryTable("role_info").Filter("Id",roleId).One(&roleInfo)
   //查询出角色对应的权限
   o.LoadRelated(&roleInfo,"Actions")
   m2m:=o.QueryM2M(&roleInfo,"Actions")
   o.Begin()
   var err1 error
   var err2 error
   for _,action:=range roleInfo.Actions {
      _,err1=m2m.Remove(action)
   }
   var actionInfo models.ActionInfo
   for i:=0;i<len(actionIdList) ;i++  {
      o.QueryTable("action_info").Filter("Id",actionIdList[i]).One(&actionInfo)
      _,err2=m2m.Add(actionInfo)
   }
   if err1!=nil||err2!=nil{
      this.Data["json"]=map[string]interface{}{"flag":"no"}
      o.Rollback()
   }else{
      this.Data["json"]=map[string]interface{}{"flag":"yes"}
      o.Commit()
   }
   this.ServeJSON()

}
```

# 十：后台布局

## 1:传统布局方式

使用了JqueryEasyUi的easyui-layout进行布局，同时使用了easyui-tabs

1.1  在这里我们使用了layout布局中的full这个案例。

 1.2 将代码full案例中的代码拷贝过来，删除“east region” 这一项DIV。

并且在头部插入一张图片，代码如下：

```html
<body class="easyui-layout">
<div data-options="region:'north',border:false" style="height:60px;background:#B3DFDA;padding:10px">
    <img src="/static/img/logo.gif">
</div>
<div data-options="region:'west',split:true,title:'West'" style="width:150px;padding:10px;">west content</div>

<div data-options="region:'south',border:false" style="height:50px;background:#A9FACD;padding:10px;">south region</div>
<div data-options="region:'center',title:'Center'"></div>
</body>
```

1.3  插入图片后，图片并没有显示全，可以将样式中的”60px“修改一下，例如修改成"80px".同时发现右侧出现了滚动条，可以将该滚动条隐藏，需要添加如下的样式：overflow: hidden

```css
style="height:80px;background:#B3DFDA;padding:10px;overflow: hidden"
```

同时可以在顶部插入一段文字。

```html
<span style="font-size:30px;color: #0000FF;font-weight: bolder;margin-left: 400px;">ItcastCMS系统管理</span>
```

同时为了美观，可以插入背景图片。

```css
background-image: url('/static/img/bg.png')
```



1.4:现在可以制作左侧的内容了，左侧是一个折叠的菜单，这里用的效果是：accordion

这里将其DEMO中的basic.html代码拷贝过来。

并且将内容进行修改，如下所示：

```html
<div data-options="region:'west',split:true,title:'West'" style="width:150px;padding:5px;">

    <div class="easyui-accordion" style="width:auto;height:auto;">
        <div title="用户管理" data-options="iconCls:'icon-ok'" style="overflow:auto;padding:10px;">
            <a href="#" >用户管理</a>
        </div>
        <div title="角色管理" data-options="iconCls:'icon-help'" style="padding:10px;">
            <a href="#" >角色管理</a>
        </div>
        <div title="权限管理" data-options="iconCls:'icon-help'" style="padding:10px;">
            <a href="#" >权限管理</a>
        </div>
    </div>


</div>
```





1.5在中间的位置，加上iframe默认显示用户管理。

```html
<div data-options="region:'center',title:'Center'">
    <iframe src="/Admin/UserInfo/Index" frameborder="0" width="100%" height="100%"></iframe>
</div>
```

但是，我们需要加上页签，这样效果更好。

这里我们使用的是页签案例是：tabs/basic.html

将其代码拷贝过来

```html
<div class="easyui-tabs" style="width:700px;height:250px">
    <div title="用户管理" style="padding:10px">
        <iframe src="/Admin/UserInfo/Index" frameborder="0" width="100%" height="100%"></iframe>
    </div>

</div>
```



但是发现页签中的页面没有填充整个中间的内容，所以需要在页签中加上也给属性。为fit=true

```html
<div class="easyui-tabs" style="width:700px;height:250px" fit="true">
```

1.6 实现单击右侧的标题，其页面在中间位置展示。

具体实现如下代码所示：



页签的创建与选择：

```javascript
<script type="text/javascript">
    $(function () {
        bindEventClick();
    })
    function bindEventClick() {
        $(".detailLink").click(function () {
            var title=$(this).text();//获取超链接的文本
            var url=$(this).attr("url");//获取超链接中的url属性
            var isExt=$("#tt").tabs("exists",title);//根据标题判断页签是否存在
            if (isExt){
                $("#tt").tabs("select",title);//如果存在，将页签选中。
                return;//不在向下执行
            }
            //如果不存在页签就添加页签。
            //添加页签
            $("#tt").tabs('add', {
                title: title,//标题
                content: createContet(url),//内容
                closable: true //出现关闭按钮
            });
        });
    }
    //创建页签
    function createContet(url) {
        var content="<iframe src="+url+" frameborder=\"0\" width=\"100%\"  height=\"100%\"></iframe>";
        return content;
    }
</script>
```

## 2:经典Windows布局

1：导入JS和CSS文件

```html
<link href="/static/lib/ligerUI/skins/Aqua/css/ligerui-all.css" rel="stylesheet" />
<link href="/static/lib/ligerUI/skins/ligerui-icons.css" rel="stylesheet" />
<script src="/static/js/jquery.js"></script>
<script src="/static/lib/ligerUI/js/ligerui.min.js"></script>
```

2；参考案例进行布局。

需要添加如下的脚本：

```javascript
<script type="text/javascript">
    $(function () {
        $("#layout1").ligerLayout({
            minLeftWidth: 80,
            minRightWidth: 80,
            allowTopResize: false,
            topHeight:90
        });
    });

</script>
```

# 十一：为用户分配权限

思考：

在前面我们已经实现了用户，角色，权限这一条线的设计，已经完成了用户权限的分配。

为什么，还要在让用户和权限在关联呢？是否多此一举呢？

思考：用户表和权限表是什么关系？

思考：如果在多对多的中间表中额外添加别的字段，应该怎样设计？

例如，我们这里需要在中间表中加上一个字段叫IsPass,该字段的作用是，用来判断用户是否具有某个权限，还是禁止具有该权限。假设，该字段的值为1，表示用户具有该权限，如果为0，表示禁止具有该权限。

思考：

在上面我们提到，如果IsPass字段的值为1，表示用户具有该权限，如果为0表示禁止具有该权限。

那么问题是，如果按照用户，角色，权限，这一条线已经给用户分配了某个具体的权限,例如“用户管理”权限，但是按照“用户”，“权限”，这一条线，却禁止该用户具有“用户管理”的权限，也就是将IsPass字段设置为了0，那么问题是，该用户是否具有“用户管理”的权限呢？

在这里，我们的设置是“没有”。

模型设计

```go
type UserInfo struct {
   Id int
   UserName string //用户名
   UserPwd string//用户密码
   DelFlag int//删除标记
   Remark string //备注
   AddDate time.Time
   ModifDate time.Time
   Roles []*RoleInfo`orm:"rel(m2m)"`
   UserActions []*UserAction `orm:"reverse(many)"`
}
```



```go
type ActionInfo struct {
   Id             int
   Remark         string
   DelFlag        int
   AddDate        time.Time
   ModifDate      time.Time
   Url            string
   HttpMethod     string
   ActionInfoName string
   ActionTypeEnum int
   MenuIcon       string
   IconWidth      int
   IconHeight     int
   Roles[] *RoleInfo `orm:"reverse(many)"`
   UserActions []*UserAction `orm:"reverse(many)"`
}
```



/中间表模型。

```go
type UserAction struct {
   Id int
   IsPass int
   Users *UserInfo `orm:"rel(fk)"`
   Actions *ActionInfo `orm:"rel(fk)"`
}
```



## 1:前端处理

在前端添加一个ifrmae



```html
<div id="setActionDiv">
    <iframe id="setUserActionFrame" frameborder="0" width="100%" height="100%"></iframe>
</div>
```

单击“为用户分配权限”按钮时，弹出对应的窗口页面。在该窗口中，将显示出所有的权限。

```javascript
//为用户分配权限
function SetUserAction() {
    var rows = $('#tt').datagrid('getSelections');
    if(rows.length!=1){
        $.messager.alert("提示","请选择要分配权限的用户!!","error");
        return;
    }
    $("#setUserActionFrame").attr("src","/Admin/UserInfo/ShowUserAction?userId="+rows[0].Id);
    $("#setActionDiv").css("display","block");
    $('#setActionDiv').dialog({
        title: '为用户分配角色',
        width: 700,
        height: 500,
        collapsible: true,
        maximizable: true,
        resizable: true,
        modal: true,
        buttons: [{
            text: 'Ok',
            iconCls: 'icon-ok',
            handler: function () {

            }
        }, {
            text: 'Cancel',
            handler: function () {
                $('#setActionDiv').dialog('close');
            }
        }]
    });


}
```

## 2:查询出用户具有的权限

/Admin/UserInfo/ShowUserAction

根据传递过来的用户编号，首先查询出用户的信息，然后查询出对应的权限信息，并且将所有的权限信息都查询出来。

```go
//为用户分配权限
func (this *UserInfoController)ShowUserAction()  {
   //接收用户Id
   userId,_:=strconv.Atoi(this.GetString("userId"))
   o:=orm.NewOrm()
   var userInfo models.Userinfo
   //查询用户信息
   o.QueryTable("userinfo").Filter("Id",userId).One(&userInfo)
   //查询出用户已经有的权限编号。
   var userExtActions []models.UserAction
   o.QueryTable("user_action").Filter("users_id",userId).All(&userExtActions)
   //查询出所有的权限信息
   var allActionList []models.ActionInfo
   o.QueryTable("action_info").Filter("del_flag",0).All(&allActionList)
   this.Data["userInfo"]=userInfo
   this.Data["allActions"]=allActionList
   this.Data["userExtActions"]=userExtActions
   this.TplName="UserInfo/ShowSetUserAction.html"
}
```

## 3:展示权限信息

在ShowUserAction方法最后，我们指定了模板文件的路径。“UserInfo/ShowSetUserAction.html”

在该模板中，我们要实现的是：

第一：展示出所有的权限。

第二：判断用户具有哪些权限。

第三：用户虽然具有某些权限，但是，还要判断用户是具有该权限，还是禁止具有该权限（也就是IsPass字段的取值）

下面的模板视图的效果是：

遍历所有的权限，每循环一次，展示出权限的编号，权限的名称，对应的URL地址，同时要通过视图函数（checkUserAction）判断是否具有该权限。如果没有直接显示“允许”和“禁止”两个单选按钮。

如果具有对应的权限，还要判断是“允许”还是“禁止”，通过视图函数checkUserActionId来完成。

如果是“允许”，对应的复选框被选中，如果是“禁止”对应的复选框被选中。

```html
为用户<span style="font-size: 14px;color:red;font-weight: bolder">{{.userInfo.UserName}}</span>分配权限
  <table width="100%">
      <tr><td>编号</td><td>权限名称</td><td>Url</td><td>操作</td></tr>
      {{range .allActions}}
            <tr>
                <td>{{.Id}}</td>
                <td>{{.ActionInfoName}}</td>
                <td>{{.Url}}}</td>
                <td>
                    {{if checkUserAction $.userExtActions .Id}}
                        {{if checkUserActionId $.userExtActions .Id}}
                            <label for="cba_{{.Id}}">允许</label>
                            <input type="radio" value="true" class="selectActions" name="cba_{{.Id}}" ids="{{.Id}}" checked="checked">
                            <label for="cba_{{.Id}}">禁止</label>
                            <input type="radio" value="false" class="selectActions" name="cba_{{.Id}}" ids="{{.Id}}">
                      {{else}}
                            <label for="cba_{{.Id}}">允许</label>
                            <input type="radio" value="true" class="selectActions" name="cba_{{.Id}}" ids="{{.Id}}">
                            <label for="cba_{{.Id}}">禁止</label>
                            <input type="radio" value="false" class="selectActions" name="cba_{{.Id}}" ids="{{.Id}}" checked="checked">
                        {{end}}
                    {{else}}
                            <label for="cba_{{.Id}}">允许</label>

                        <input type="radio" value="true" class="selectActions" name="cba_{{.Id}}" ids="{{.Id}}">
                        <label for="cba_{{.Id}}">禁止</label>
                        <input type="radio" value="false" class="selectActions" name="cba_{{.Id}}" ids="{{.Id}}">
                     {{end}}
                    <input type="button" value="删除" class="btnClearActions" ids="{{.Id}}">

                </td>
            </tr>
       {{end}}
  </table>
```

对应的视图模板函数

```go
//判断用户是否具有某个权限
func CheckUserAction(userExtActionList[]models.UserAction,actionId int)(b bool) {
   b=false
   for i:=0;i<len(userExtActionList);i++ {
      if userExtActionList[i].Actions.Id==actionId {
         b=true
         break
      }
   }
   return
}
//判断用户具有某个权限是禁止还是允许
func CheckUserActionId(userExtActionList[]models.UserAction,actionId int)(b bool)  {
     b=false
   for i:=0;i<len(userExtActionList) ;i++  {
      if userExtActionList[i].Actions.Id==actionId {
         if userExtActionList[i].IsPass==1 {
            b=true
         }
         break//注意break的位置
      }

   }
   return
}
```

## 4:完成用户权限分配

当用户单击“允许”单选按钮时，或者是“禁止”按钮时，都是用来完成对权限的分配，只不过当单击“允许”单选按钮时，表示用户具有该权限，单击“禁止”时，表示用户禁止具有该权限。

所以在这里，我们需要给所有的“单选按钮”绑定对应的单击事件，我们在这里也给所有的单选按钮指定了类选择器为“.selectActions”.

前端处理

```javascript
<script type="text/javascript">
    $(function () {
        //完成用户权限分配
        $(".selectActions").click(function () {
            setUserAction($(this));
        })
       
    })

```

```javascript
给用户分配权限，实际上是将对应的数据，插入到用户和权限对应的中间表中。
所以在对应的方法setUserAction中，我们通过AJAX的方式，将权限的编号，用户的编号，以及isPass(表示是允许具有权限，还是禁止具有该权限)的值，发送到服务端。
在这里，我们给每个单选按钮指定了一个属性为ids,保存了对应的权限编号。这个属性是我们人为加上的，并不是单选按钮具有的属性，但是并不会出现错误。
isPass的值就是单选按钮的value属性的值，也就是true,或者是false.


//完成用户权限的分配
function setUserAction(control) {
    var actionId=control.attr("ids");
    var isPass=control.val()
    $.post("/Admin/UserInfo/SetUserAction",{"actionId":actionId,"isPass":isPass,"userId":{{.userInfo.Id}}},function (data) {
        if(data.flag=="ok"){
            $.messager.show({
                title: '提示',
                msg: '权限分配成功',
                showType: 'show'
            })
        }else{
            $.messager.show({
                title: '提示',
                msg: '权限分配失败',
                showType: 'show'
            })
        }
    })
}
```

服务端处理

服务端接收客户端传递过来的，权限编号，isPass的值，还有用户的编号。

在这里获取isPass的值，判断是true还是false,

然后，根据用户编号，和权限编号，从中间表user_action中去查询，看一下该用户是否具有对应的权限，如果有只需要将IsPass的值，修改一下就可以了。

如果没有，那么需要重新向中间表中添加记录。

```go
//完成用户权限的分配
func (this *UserInfoController)SetUserAction()  {
   actionId,_:=strconv.Atoi(this.GetString("actionId"))
   isPass,_:=this.GetBool("isPass")
   userId,_:=strconv.Atoi(this.GetString("userId"))
   var isExt int
   if isPass{
      isExt=1
   }else{
      isExt=0
   }
   o:=orm.NewOrm()
   var userAction models.UserAction
   o.QueryTable("user_action").Filter("users_id",userId).Filter("actions_id",actionId).One(&userAction)
   if userAction.Id>0{
      //如果用户有权限，直接修改
      userAction.IsPass=isExt
      o.Update(&userAction)
   }else{
      //如果没有权限，直接添加
      var actionInfo models.ActionInfo
      o.QueryTable("action_info").Filter("Id",actionId).One(&actionInfo)
      var userInfo models.Userinfo
      o.QueryTable("userinfo").Filter("Id",userId).One(&userInfo)
      userAction.IsPass=isExt
      userAction.Actions=&actionInfo
      userAction.Users=&userInfo
      o.Insert(&userAction)

   }
   this.Data["json"]=map[string]interface{}{"flag":"ok"}
   this.ServeJSON()
}
```

5:删除用户的权限

在前端的视图中，在每个权限的后面添加了一个删除按钮，单击该按钮，可以删除用户对应的权限。

所以，首先也是要给按钮绑定对应的单击事件。

前端处理

```javascript
$(function () {
    //完成用户权限分配
    $(".selectActions").click(function () {
        setUserAction($(this));
    })
    //清除用户权限
    $(".btnClearActions").click(function () {
        clearUserAction($(this));
    })
})
```

```javascript
将用户的编号，和权限编号传递到服务端。用来删除对应用户的权限。
这里需要注意的是，当服务端删除成功，会返回对应的消息，这时会调用AJAX的回调函数，在该回调函数中，我们需要将对应的单选按钮的选中状态取消。

//清除用户权限
function clearUserAction(control) {
    var actionId=control.attr("ids");
    var userId={{.userInfo.Id}}
    $.post("/Admin/UserInfo/DeleteUserAction",{"actionId":actionId,"userId":userId},function (data) {
        if(data.flag=="ok"){
            control.parent().find(".selectActions").removeAttr("checked");//注意parent()是方法,find()方法找类选择器，所以要加上"."
            $.messager.show({
                title: '提示',
                msg: '权限删除成功!!',
                showType: 'show'
            })
        }else{
            $.messager.show({
                title: '提示',
                msg: '权限删除失败',
                showType: 'show'
            })
        }
    })
}
```

服务端处理

/Admin/UserInfo/DeleteUserAction

服务端接收到用户编号和权限编号，删除对应的信息。

```go
//删除用户权限
func (this *UserInfoController)DeleteUserAction()  {
   actionId,_:=strconv.Atoi(this.GetString("actionId"))
   userId,_:=strconv.Atoi(this.GetString("userId"))
   o:=orm.NewOrm()
   var userAciton models.UserAction
   o.QueryTable("user_action").Filter("users_id",userId).Filter("actions_id",actionId).One(&userAciton)
   o.Delete(&userAciton)

   this.Data["json"]=map[string]interface{}{"flag":"ok"}
   this.ServeJSON()
}
```

# 十二：菜单权限过滤



用户登录成功，会跳转到我们整个项目中的首页，在该页面中，显示了所有的菜单。当单击，不同的菜单后，显示出对应的页面。但是问题是，不同的用户由于权限不同，登录成功，进入首页，看到的菜单项应该是不同的。那么下面，我们就来实现对应的菜单权限的过滤，首先，我们先来做一个简单的登录。

## 1：用户登录

1.1 前端页面的设计

导入JS文件

```javascript
<script src="/static/js/jquery.js" type="text/javascript"></script>
<script src="/static/js/jquery.unobtrusive-ajax.min.js" type="text/javascript"></script>
```

1.2:添加对应的CSS样式

```css
<style type="text/css">
    *
    {
        padding: 0;
        margin: 0;
    }
    body
    {
        text-align: center;
        background: #4974A4;
    }
    #login
    {
        width: 740px;
        margin: 0 auto;
        font-size: 12px;
    }
    #loginlogo
    {
        width: 700px;
        height: 100px;
        overflow: hidden;
        background: url('/static/img/login/logo.png') no-repeat;
        margin-top: 50px;
    }
    #loginpanel
    {
        width: 729px;
        position: relative;
        height: 300px;
    }
    .panel-h
    {
        width: 729px;
        height: 20px;
        background: url('/static/img/login/panel-h.gif') no-repeat;
        position: absolute;
        top: 0px;
        left: 0px;
        z-index: 3;
    }
    .panel-f
    {
        width: 729px;
        height: 13px;
        background: url('/static/img/login/panel-f.gif') no-repeat;
        position: absolute;
        bottom: 0px;
        left: 0px;
        z-index: 3;
    }
    .panel-c
    {
        z-index: 2;
        background: url('/static/img/login/panel-c.gif') repeat-y;
        width: 729px;
        height: 300px;
    }
    .panel-c-l
    {
        position: absolute;
        left: 60px;
        top: 40px;
    }
    .panel-c-r
    {
        position: absolute;
        right: 20px;
        top: 50px;
        width: 222px;
        line-height: 200%;
        text-align: left;
    }
    .panel-c-l h3
    {
        color: #556A85;
        margin-bottom: 10px;
    }
    .panel-c-l td
    {
        padding: 7px;
    }
    .login-text
    {
        height: 24px;
        left: 24px;
        border: 1px solid #e9e9e9;
        background: #f9f9f9;
    }
    .login-text-focus
    {
        border: 1px solid #E6BF73;
    }
    .login-btn
    {
        width: 114px;
        height: 29px;
        color: #E9FFFF;
        line-height: 29px;
        background: url('/static/img/login/login-btn.gif') no-repeat;
        border: none;
        overflow: hidden;
        cursor: pointer;
    }
    #txtUsername, #code, #txtPassword
    {
        width: 191px;
    }
    #logincopyright
    {
        text-align: center;
        color: White;
        margin-top: 50px;
    }
    a
    {
        color: Black;
    }
    a:hover
    {
        color: Red;
        text-decoration: underline;
    }
</style>
```



1.3 创建表单

```html
<body style="padding: 10px">

<div id="login">
    <div id="loginlogo">
    </div>
    <div id="loginpanel">
        <div class="panel-h">
        </div>
        <div class="panel-c">
            <div class="panel-c-l">

                <form  data-ajax="true" data-ajax-method="post" data-ajax-success="afterLogin" data-ajax-url="/Login/UserLogin" id="LoginForm">
                    <table cellpadding="0" cellspacing="0">
                        <tbody>
                        <tr>
                            <td align="left" colspan="2">
                                <h3>
                                    请使用ItcastCMS管理系统登录管理系统账号登录</h3>
                            </td>
                        </tr>
                        <tr>
                            <td align="right">
                                账号：
                            </td>
                            <td align="left">
                                <input type="text" name="LoginCode" id="LoginCode" class="login-text" />

                            </td>
                        </tr>
                        <tr>
                            <td align="right">
                                密码：
                            </td>
                            <td align="left">
                                <input type="password" name="LoginPwd" id="LoginPwd" value="123" class="login-text" />
                            </td>
                        </tr>

                        <tr>
                            <td>

                            </td>
                            <td>


                            </td>
                        </tr>
                        <tr>
                            <td align="center" colspan="2">
                                <input type="submit" id="btnLogin" value="登录" class="login-btn" />


                            </td>
                            <td>
                                <span id="errorMsg" style="font-size:14px;color:red;display:none"></span>
                            </td>
                        </tr>
                        </tbody>
                    </table>
                </form>
            </div>
            <div class="panel-c-r">
                <p>
                    请从左侧输入登录账号和密码登录</p>
                <p>
                    如果遇到系统问题，请联系网络管理员。</p>
                <p>
                    如果没有账号，请联系网站管理员。
                </p>
                <p>
                    ......</p>
            </div>
        </div>
        <div class="panel-f">
        </div>
    </div>
    <div id="logincopyright">
        Copyright @ 2018 itcast.cn
    </div>
</div>
</body>
```



1.4  登录成功后，执行AJAX的回调函数

```javascript
<script type="text/javascript">
    function afterLogin(data) {
        if(data.flag=="ok"){
            window.location.href="/Admin/Home/Index"
        }else{
            $("#errorMsg").css("display","block");
            $("#errorMsg").text("用户登录失败!!")
        }
    }
</script>
```

1.5 服务端实现用户登录。

判断输入的用户名和密码是否正确，如果正确将用户名和用户的编号存储到session中。

```go
func (this *LoginController)UserLogin()  {
  userName:=this.GetString("LoginCode")
  userPwd:=this.GetString("LoginPwd")
  o:=orm.NewOrm()
  var userInfo models.UserInfo
  o.QueryTable("user_info").Filter("user_name",userName).Filter("user_pwd",userPwd).One(&userInfo)
  if userInfo.Id>0{
   this.SetSession("userName",userName)
   this.SetSession("userId",userInfo.Id)
   this.Data["json"]=map[string]interface{}{"flag":"ok"}
  }else{
   this.Data["json"]=map[string]interface{}{"flag":"no"}
  }
  this.ServeJSON()
}
```

## 2：菜单权限的过滤

2.1：前端准备

通过上面的登录成功后，跳转到了/Admin/Home/Index，

在这个页面中，发起一个ajax请求,请求的地址对应的GetMenus方法，就是进行菜单权限过滤的方法。

最终该方法返回的是权限的名称，对应的图片地址，以及url地址。并将赋值给links数组，最后将该数组进行遍历，将其存储的数据打印在窗口中。

```javascript
$.post("/Admin/Home/GetMenus",{},function (data) {
    links=data.menus;
    linksInit();
    onResize();
})
```

2.2 进行菜单权限的过滤。

整个菜单权限过滤的基本思路是;

第一：获取登录用户的信息

第二：根据登录用户，查询出该登录用户具有的角色信息。

第三：根据角色信息，查询出对应的权限信息。

第四：对查询出的权限信息进行判断，判断一下是否为菜单权限，也就是ActionTypeEnum属性的取值是否为1.

第五：上面是按照“用户--角色--权限”，这条线进行权限过滤的，那么下面就是按照“用户--权限""这条线进行过滤。

第六：按照“用户---权限”这条线查询出用户对应的权限后，还要进行菜单权限的过滤。

第七：将两条线查询出的权限进行合并。

第八：去重操作

第九：过滤掉登录用户禁用的权限。也就是判断登录用户中关于isPass的取值。

第十：前端内容的修改。

 <1>从session中获取登录用户的信息。

```go
userId:=this.GetSession("userId")
//1:获取用户的信息。
o:=orm.NewOrm()
var userInfo models.UserInfo
o.QueryTable("user_info").Filter("id",userId).One(&userInfo)
```

<2>获取登录用户的角色信息。

```go
var roles []*models.RoleInfo
  o.LoadRelated(&userInfo,"Roles")
for _,role:= range userInfo.Roles{
   roles=append(roles,role)
}
```

<3>获取角色对应的权限信息。

```go
var actions []*models.ActionInfo
for i:=0;i<len(roles);i++ {
   o.LoadRelated(roles[i],"Actions")//注意，这里roles[i]不能在加取地址符号&，因为本身存储的就是*
   for _,action:= range roles[i].Actions {
       actions=append(actions,action)
   }
}
```

<4>找出上面权限中的菜单权限。

对actions切片集合中存储的权限信息进行过滤，通过循环的方式，每循环一次取出对应的值，判断其ActionTypeEnum属性的取值是否为1.

```go
var menuActions []*models.ActionInfo
for i:=0;i<len(actions);i++ {
   if actions[i].ActionTypeEnum==1{
      menuActions=append(menuActions,actions[i])
   }
}
```

<5>按照“用户--权限”这条线进行过滤，先按照这条线找出对应的权限。

```go
var subActions[]models.UserAction
o.QueryTable("user_action").Filter("users_id",userId).All(&subActions)
```

<6>按照“用户---权限”这条线查询出用户对应的权限后，还要进行菜单权限的过滤。

```go
var actionInfo models.ActionInfo
var subMenuActions []*models.ActionInfo
for i:=0;i<len(subActions) ;i++  {
   //这里注意，过滤的条件是action_info权限表中的id与中间表中的id进行比较。
   o.QueryTable("action_info").Filter("id",subActions[i].Actions.Id).Filter("action_type_enum",1).One(&actionInfo)
   subMenuActions=append(subMenuActions,&actionInfo)
}
```



<7>:将两条线查询出的权限进行合并。

```go
menuActions=append(menuActions,subMenuActions...)
```

```
第二个参数也可以直接写另一个切片，将它里面所有元素拷贝追加到第一个切片后面。
要注意的是，这种用法函数的参数只能接收两个slice，并且末尾要加三个点

```



<8>去重操作

```go
temp:=RemoveRepeatedElement(menuActions)
```

```go
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
```

(9)过滤掉登录用户禁用的权限。也就是判断登录用户中关于isPass的取值

这里

```go
 var subForActions[]models.UserAction
   //这里一定要注意，过滤的是当前登录用户中的禁用权限。
   o.QueryTable("user_action").Filter("is_pass",0).Filter("users_id",userId).All(&subForActions)
   if len(subForActions)>0{
      for i,action:= range temp{
         if CheckForAction(subForActions,action.Id) {
            temp=append(temp[:i],temp[i+1:]...)
         }
      }
   }
   this.Data["json"]=map[string]interface{}{"menus":temp}
this.ServeJSON()
```



```
//判断禁用权限。
 func CheckForAction(subForActions[]models.UserAction,actionId int)(b bool){
   b=false
    for i:=0;i<len(subForActions) ;i++  {
       if subForActions[i].Actions.Id==actionId {
         b=true;
         break
       }

    }
    return
 }
```

（10）前端的处理。

```go
$.post("/Admin/Home/GetMenus",{},function (data) {
    links=data.menus;
    linksInit();
    onResize();
})
```

现在已经将菜单权限返回了，然后赋值给了对应的数组。

# 十三 非菜单权限过滤

所谓的非菜单权限，指的就是不再主页面中显示的内容。例如，用户的添加，编辑，删除等操作。

对这样的权限应该怎样进行过滤呢？

在这里，可以通过过滤器来实现。

整个非菜单权限过滤的步骤如下：

1：从session中获取登录用户的名称。

2：获取用户请求的地址与请求的方式。

3：根据获取的请求的地址和请求的方式，从权限表中查询出对应的权限。

4：如果没有找到对应的权限跳转到登录页面。

5：如果找到了，就要判断当前登录用户是否具有该权限，可以现按照“用户--权限”这条线进行过滤。

6：如果按照“用户--权限”这条线，发现登录用户具有该权限，那么就要判断是否是“禁用”还是允许。

7：如果是“允许”就结束整个判断，如果是“禁止”，就要按照“用户--角色--权限”这条线进行判断。

8：如果按照“用户--角色--权限”这条线进行过滤，用户有权限，进行访问，如果没有权限跳转到登录页面。



第一：从session中获取登录用户的名称,并且进行判断，如果为空跳转到登录页面。

```go
userName:= ctx.Input.Session("userName")
if userName!=""{
    
}else{
    ctx.Redirect(302,"/Login/Index")
}
```

第二：如果session中有值，获取用户请求的地址与请求的方式。

```go
//接收用户请求的url地址
url:=ctx.Request.URL.Path
//接收请求的方法。
httpMethod:=ctx.Request.Method;
//根据请求的地址与请求的方法从权限表中，查找具体的权限信息。
```

第三:根据获取的请求的地址和请求的方式，从权限表中查询出对应的权限。

```go
var actionInfo models.ActionInfo
o:=orm.NewOrm()
o.QueryTable("action_info").Filter("url",url).Filter("http_method",httpMethod).One(&actionInfo)
```



第四：如果没有找到对应的权限跳转到登录页面

```go
if actionInfo.Id>0{
}else{
    ctx.Redirect(302,"/Login/Index")
}
```

第五：如果找到了，就要判断当前登录用户是否具有该权限，可以现按照“用户--权限”这条线进行过滤

```go
var userInfo models.UserInfo
//查询用户的信息.
o.QueryTable("user_info").Filter("user_name",userName).One(&userInfo)
//根据用户权限这条线进行过滤。
var userAction models.UserAction
o.QueryTable("user_action").Filter("users_id",userInfo.Id).Filter("actions_id",actionInfo.Id).One(&userAction)
```



第六：如果按照“用户--权限”这条线，发现登录用户具有该权限，那么就要判断是否是“禁用”还是允许。

```go
if userAction.Id>0{
   //判断是允许的还是禁止的。
   if userAction.IsPass==1{
      return
   }else {
      ctx.Redirect(302,"/Login/Index")
   }
}else{
    //按照用户--角色--权限进行判断。
    
}
```

第七：如果是“允许”就结束整个判断，如果是“禁止”，就要按照“用户--角色--权限”这条线进行判断

  1：找出登录用户具有的角色。

```go
var roles []*models.RoleInfo
   o.LoadRelated(&userInfo,"Roles")
for _,role:= range userInfo.Roles{
   roles=append(roles,role)
}
```

2；找出角色对应的权限。

```go
var actions []*models.ActionInfo
   var roleInfo models.RoleInfo
for _,role:=range roles{

              o.QueryTable("role_info").Filter("id",role.Id).One(&roleInfo)
   o.LoadRelated(&roleInfo,"Actions")
   for _,action:= range roleInfo.Actions{
     if action.Id==actionInfo.Id{//这里需要与用户访问的地址对应的权限进行比较。
      actions=append(actions,action)
     }
   }
   }
```

第八：如果按照“用户--角色--权限”这条线进行过滤，用户有权限，进行访问，如果没有权限跳转到登录页面

```go
if len(actions)<1{
   ctx.Redirect(302,"/Login/Index")
}
```



第九：完成过滤器的注册

```go
beego.InsertFilter("/Admin/*",beego.BeforeExec,FilterUserAction)
```





如下是完整的代码：

```go
func FilterUserAction(ctx *context.Context)  {
   userName:= ctx.Input.Session("userName")
   if userName!=""{
      if userName=="laowang"{
         return
      }
      //接收用户请求的url地址
      url:=ctx.Request.URL.Path
      //接收请求的方法。
      httpMethod:=ctx.Request.Method;
      //根据请求的地址与请求的方法从权限表中，查找具体的权限信息。
      var actionInfo models.ActionInfo
      o:=orm.NewOrm()
      o.QueryTable("action_info").Filter("url",url).Filter("http_method",httpMethod).One(&actionInfo)
      if actionInfo.Id>0{
         var userInfo models.UserInfo
         //查询用户的信息.
         o.QueryTable("user_info").Filter("user_name",userName).One(&userInfo)
         //根据用户权限这条线进行过滤。
         var userAction models.UserAction
         o.QueryTable("user_action").Filter("users_id",userInfo.Id).Filter("actions_id",actionInfo.Id).One(&userAction)
         //判断权限是允许的还是禁止的。
         //如果下面的条件成立，表示根据用户权限这一条线，找到对应的记录了。
         if userAction.Id>0{
            //判断是允许的还是禁止的。
            if userAction.IsPass==1{
               return
            }else {
               ctx.Redirect(302,"/Login/Index")
            }
         }else{
               //如果按照用户权限这一条线没有找到，那么按照用户，角色，权限这一条线进行查询。
               //获取登录用户具有的角色信息。
               var roles []*models.RoleInfo
               o.LoadRelated(&userInfo,"Roles")
            for _,role:= range userInfo.Roles{
               roles=append(roles,role)
            }
             //查询出角色对应的权限。
             var actions []*models.ActionInfo
             var roleInfo models.RoleInfo
            for _,role:=range roles{

                  o.QueryTable("role_info").Filter("id",role.Id).One(&roleInfo)
               o.LoadRelated(&roleInfo,"Actions")
               for _,action:= range roleInfo.Actions{
                 if action.Id==actionInfo.Id{
                  actions=append(actions,action)
                 }
               }
            }
            if len(actions)<1{
               ctx.Redirect(302,"/Login/Index")
            }

         }
      }else{
         ctx.Redirect(302,"/Login/Index")
      }
   }else{
      ctx.Redirect(302,"/Login/Index")
   }
}
```



# 十四：新闻类别管理

https://www.layui.com/   （文档：https://www.layui.com/doc/base/element.html）

在前面的课程中，我们已经完成了权限的讲解，在前面的课程中，主要大家要掌握的是，

第一：关于数据库中表的设计（或者也可以称之为模型的设计），例如，用户与角色的多对多关系，角色与权限的多对多关系，用户与权限的多对多关系（这里，我们为了在中间的表中加入额外的字段,中间模型是我们自己设计的）。

第二：前端AJAX的处理，与Jquery的基本操作要掌握。对jqueryEasyUI这种前端的框架（库），需要了解。能根据其提供的案例进行改造。

第三：理清楚整个权限的处理业务。

第四：针对数据库的查询等等操作。

下面我们主要学习一下关于新闻发布的内容。

## 1：新闻类别模型设计

```go
//文章类别
type ArticelClass struct {
   Id int//主键
   ClassName string//类别名称
   ParentId int //父类别的编号
   CreateUserId int //创建类别的用户编号
   CreateDate time.Time //创建时间
   DelFlag int //删除标记
   Remark string //备注
   }
```



思考：我们都知道一个类别A下面还有类别A1，并且A1类别下面有肯能有A11,也就是整个类别构建成了一个树，那么我们应该怎样存储这样的树形结构呢？

## 2：根类别信息展示

### 2.2.1 前端处理

在这里，与前面的设计是一样的，值不过，这里我们在页面上添加一个treegrid树形表格来展示数据。具体的做法如下。

在ArticelClassController中添加一个Index方法，该方法将呈现一个视图，在该视图中添加treegrid表格。

```go
func (this *ArticelClassController)Index()  {
   this.TplName="ArticelClass/Index.html"
}
```

对应的Index.html的视图的内容如下：

第一：添加对应的JS与CSS的引用。

```js
<script src="/static/js/jquery.js" type="text/javascript"></script>
<script src="/static/js/jquery.easyui.min.js" type="text/javascript"></script>
<script src="/static/js/easyui-lang-zh_CN.js" type="text/javascript"></script>
<link href="/static/css/themes/default/easyui.css" rel="stylesheet" />
<link href="/static/css/themes/icon.css" rel="stylesheet" />
```



第二：在页面上添加一个table表格标签。

```go
<table id="tt"></table>
```

第三：使用treegrid表格。

```js
function loadData() {
    $('#tt').treegrid({
        title: '栏目管理',
        iconCls: 'icon-save',
        width: 500,
        height: 350,
        nowrap: false,
        rownumbers: true,
        animate: true,
        collapsible: true,
        url: '/Admin/ArticelClass/ShowArticelClass',
        idField: 'Id',
        treeField: 'ClassName',
        lines: true,
        columns: [[
            { field: 'Id', title: '编号', width: 150, rowspan: 2 },
            { field: 'ClassName', title: '栏目名称', width: 120 },
            { field: 'Remark', title: '备注', width: 120, rowspan: 2 }

        ]],
        onClickRow: function (row) {
            //根据所单击的行，获取对应的子类别.
            $.post("/Admin/ArticelClass/ShowChildClass", { "id": row.Id }, function (data) {
                //先清空，后追加.如果没有数据不追加
                if (data.rows.length != 0) {
                    var nodes = $('#tt').treegrid('getChildren', row.Id);
                    for (var i = 0; i < nodes.length; i++) {
                        $('#tt').treegrid('remove', nodes[i].id);
                    }
                    $('#tt').treegrid('append', {
                        parent: row.Id,
                        data: data.rows
                    });
                }
            });
        },
        toolbar: [
            {
                id: 'btnAddParent',
                text: '添加根栏目',
                iconCls: 'icon-add',
                handler: function () {
                    addParentArticel();

                }
            }
            ,{
                id: 'btnAdd',
                text: '添加子栏目',
                iconCls: 'icon-add',
                handler: function () {
                    addChildArticel();

                }
            }, {
                id: 'btnEdit',
                text: '编辑',
                iconCls: 'icon-edit',
                handler: function () {
                    // showChildArticel();

                }
            }]
    })
}
```

同时在$(function)中调用loadData方法。

这里需要注意的是，这些代码都是固定的用法，大家不需要记忆，只需要拷贝过来进行修改就可以。而且，通过上面的代码，我们发现与dataGrid表格是非常相似，大家一定要能够触类旁通。

注意：在treegrid的表格中，我们加入了一个onClickRow事件，该事件是在单击treegrid表格中行的时候，触发。

在其对应的函数中，定义了一个AJAX的POST请求，其目的是查询根类别下的子类别。（这里默认行中只显示根类别，对应的子类别不显示，只有单击根类别后通过该POST请求，获取出对应的子类别。）

这里将返回的子类别添加到treegrid上，再添加之前先删除，然后再添加。（这些代码，也不要记忆，只要能看懂就行，因为都是treegrid提供的api操作，文档中都有。）

### 2.2.2  服务端处理

/Admin/ArticelClass/ShowArticelClass

当页面加载完成后，会自动向/Admin/ArticelClass/ShowArticelClass，发送请求，来获取数据。

下面看一下ArticelClass控制器下的ShowArticelClass方法的实现。

```go
//展示根类别的信息
func(this *ArticelClassController)ShowArticelClass(){
   o:=orm.NewOrm()
   var articelClasses[]models.ArticelClass
   o.QueryTable("articel_class").Filter("parent_id",0).All(&articelClasses)
   this.Data["json"]=map[string]interface{}{"rows":articelClasses}
   this.ServeJSON()
}
```

在上面的查询中，注意查询的条件，这里只查询根类别信息。

将上面查询出的结果生成JSON返回后，填充到treegrid表格中，在treegrid表中，每一行都只显示根类别的信息，但是我们给每一行加上了一个onClickRow事件，当单击行时触发。这里发送一个AJAX的POST请求，请求的地址

/Admin/ArticelClass/ShowChildClass，给该地址传递的参数是根类别的编号。

在ArticelClass控制器下的ShowChildClass方法中接收传递过来的根类别编号，查询出对应的子类别，然后返回，这时在treegrid树形表格中，根类别下面会展示出对应的子类别，这样就构建成了一个树形结构。

ArticelClass控制器下的ShowChildClass方法的具体代码实现如下：

```go
//展示根类别下的子类别。
func (this *ArticelClassController)ShowChildClass()  {
  cId,_:=this.GetInt("id")
  o:=orm.NewOrm()
   var allClass []models.ArticelClass
  o.QueryTable("articel_class").Filter("parent_id",cId).All(&allClass)
  this.Data["json"]=map[string]interface{}{"rows":allClass}
  this.ServeJSON()
}
```

上面的代码是接收传递过来的根类别编号，查询出对应的子类别信息，生成JSON并且返回。

在上面的代码中，要注意的是查询的条件。



## 3：根类别信息添加



### 3.1 主页面前端处理 

第一：在页面上添加一个DIV，在DIV中添加一个iframe,并且让该DIV隐藏。

```html
<div id="addParetnDiv">
    <iframe id="addParentFrame" frameborder="0" width="100%" height="100%"></iframe>
</div>
```



```javascript
$(function () {
    loadData();
    $("#addParetnDiv").css("display","none");//隐藏DIV

})
```

并且在treegrid中添加一个toolbar

```javascript
{
    id: 'btnAddParent',
    text: '添加根栏目',
    iconCls: 'icon-add',
    handler: function () {
        addParentArticel();

    }
}
```



对应的 addParentArticel();首先完成对addParentFrame这个iframe的src属性的设置，同时将窗口弹出了，该方法的代码实现如下：

```javascript
function addParentArticel() {
    //表示选择对应的根类别

    $("#addParentFrame").attr("src","/Admin/ArticelClass/ShowAddParentClass")
    $("#addParetnDiv").css("display","block");
    $('#addParetnDiv').dialog({
        title: '添加根类别信息',
        width: 300,
        height: 300,
        collapsible: true,
        maximizable: true,
        resizable: true,
        modal: true,
        buttons: [{
            text: 'Ok',
            iconCls: 'icon-ok',
            handler: function () {
                //提交表单。
                var childWindows=$("#addParentFrame")[0].contentWindow;//获取子窗体的windows对象。

                childWindows.subForm()//调用，子页面中的subForm方法，来完成表单的提交

            }
        }, {
            text: 'Cancel',
            handler: function () {
                $('#addParetnDiv').dialog('close');
            }
        }]
    });
}
```

最后，子页面处理完成后，会执行主页面中的如下的函数：

```go
//添加完根类别后调用该方法
function afterParentAdd(data){
    if(data.flag=="ok"){
        $('#addParetnDiv').dialog('close');
        $('#tt').treegrid('reload');
    }
}
```



将主页面的窗口关闭，并且刷新整个treegrid表格。

### 3.2 子页面的处理

上面，在主页面中，指定了一个叫addParentFrame的iframe标签,并且给iframe指定了src为/Admin/ArticelClass/ShowAddParentClass

对应的ShowAddParentClass方法指定了一个视图的地址，最终该视图会被填充到iframe中。

该方法的代码实现如下：

```go
func (this *ArticelClassController)ShowAddParentClass()  {
   this.TplName="ArticelClass/ShowAddParentClass.html"
}
```

ShowAddParentClass.html视图中呈现出就是一个添加表单，代码如下：

```html
<form  data-ajax="true" data-ajax-method="post" data-ajax-success="afterAdd" data-ajax-url="/Admin/ArticelClass/AddParentClass" id="addForm">
    <table>
        <tr><td>类别名称</td><td><input type="text" name="className"></td></tr>
        <tr><td>备注</td><td><input type="text" name="Remark" /></td></tr>
    </table>

</form>
```

注意：添加JS文件的引入：

```js
<script src="/static/js/jquery.js" type="text/javascript"></script>
<script src="/static/js/jquery.unobtrusive-ajax.min.js" type="text/javascript"></script>
```

在该子页面中定义如下的方法，有主页面调用该方法，完成表单的提交。

```javascript
function subForm() {
    $("#addForm").submit();
}
```

这时，会向/Admin/ArticelClass/AddParentClass这个页面发送post请求，完成根类别信息的添加。

AddParentClass方法的代码如下：

```go
//添加根类别
func (this *ArticelClassController)AddParentClass()  {
  var articelClassParent=models.ArticelClass{}
  articelClassParent.DelFlag=0
  articelClassParent.ClassName=this.GetString("className")
  articelClassParent.Remark=this.GetString("Remark")
  articelClassParent.ParentId=0
  articelClassParent.CreateDate=time.Now()
  value,ok:=this.GetSession("userId").(int)
  if ok{
     articelClassParent.CreateUserId=value
  }
  o:=orm.NewOrm()
 _,err:=o.Insert(&articelClassParent)
  if err==nil{
   this.Data["json"]=map[string]interface{}{"flag":"ok"}
  }else{
   this.Data["json"]=map[string]interface{}{"flag":"non"}
  }
  this.ServeJSON()
}
```

在上面的添加中，要注意如下两点内容。

第一：ParentId属性的取值为0，因为添加的是根类别。

第二： value,ok:=this.GetSession("userId").(int) ：表示断言，断言其session中取出的值。

根类别信息添加完成后，会调用回调函数afterAdd。在该函数中完成对主函数的调用，代码如下：

```javascript
function afterAdd(data) {
    window.parent.afterParentAdd(data);//调用父窗体的方法
}
```

afterParentAdd方法在主窗体中已经完成定义了，就是将窗口关闭，并且刷新treegrid表格。

## 4：子类别信息添加

### 4.1 主页面前端处理

在主页面中，添加一个toolBar,

```javascript
{
    id: 'btnAdd',
    text: '添加子栏目',
    iconCls: 'icon-add',
    handler: function () {
        addChildArticel();

    }
}
```



同时也要添加一个iframe.

```html
<div id="addChildDiv">
    <iframe id="addChildFrame" frameborder="0" width="100%" height="100%"></iframe>
</div>
```

并且让其隐藏。

```javascript
$(function () {
    loadData();
    $("#addParetnDiv").css("display","none");
    $("#addChildDiv").css("display","none");
})
```

对应的addChildArticel方法的实现：

```javascript
function addChildArticel() {
    var row = $('#tt').treegrid('getSelected');
    if (row != null) {
        $("#addChildFrame").attr("src", "/Admin/ArticelClass/ShowAddChildClass?cId="+row.Id);
        $("#addChildDiv").css("display", "block");
        $('#addChildDiv').dialog({
            title: '添加子类别信息',
            width: 300,
            height: 300,
            collapsible: true,
            maximizable: true,
            resizable: true,
            modal: true,
            buttons: [{
                text: 'Ok',
                iconCls: 'icon-ok',
                handler: function () {
                    //提交表单。
                    var childWindows = $("#addChildFrame")[0].contentWindow;//获取子窗体的windows对象。

                    childWindows.subForm()//提交表单

                }
            }, {
                text: 'Cancel',
                handler: function () {
                    $('#addChildDiv').dialog('close');
                }
            }]
        });
    }else{

        $.messager.alert("提示","请选择对应的根类别!!","error");

    }
}
```

在上面的代码中，首先判断是否选中了根类别，在选中的根类别下面添加对应的子类别。

所以在给iframe指定的src路径中，将根类别的编号传递到该地址中。

然后通过对话框，将iframe中嵌入的网页显示出来。

指定一个子窗体中添加完后，调用主窗体的方法。

```javascript
//添加完子类别后调用该方法。
function afterChiledAdd(data) {
    if(data.flag=="ok"){
        $('#addChildDiv').dialog('close');
        $('#tt').treegrid('reload');
    }
}
```

### 4.2  服务端处理

"/Admin/ArticelClass/ShowAddChildClass?cId="+row.Id

在ShowAddChildClass中接收传递过来的根类别的编号，将根类别的信息查询出来后，传递到ShowAddChildClass.html视图中。

```go
//展示添加子类别的页面。
func (this *ArticelClassController)ShowAddChildClass()  {
   cId,_:=this.GetInt("cId")
   o:=orm.NewOrm()
   var articelClass models.ArticelClass
   o.QueryTable("articel_class").Filter("id",cId).One(&articelClass)
   this.Data["classInfo"]=articelClass
   this.TplName="ArticelClass/ShowAddChildClass.html"
}
```

ShowAddChildClass.html视图中展示的内容如下：

```html
<div>为类别<span style="font-size: 14px; color: red; font-weight: bold">{{.classInfo.ClassName}}</span>添加子类别</div>
<form  data-ajax="true" data-ajax-method="post" data-ajax-success="afterAdd" data-ajax-url="/Admin/ArticelClass/AddChildClass" id="addForm">
    <input type="hidden" value="{{.classInfo.Id}}" name="cId">
    <table>
        <tr><td>类别名称</td><td><input type="text" name="className"></td></tr>
        <tr><td>备注</td><td><input type="text" name="Remark" /></td></tr>
    </table>

</form>
```

就是一个添加的表单。注意要引用对应的Js文件。

```js
<link href="/static/css/tableStyle.css" rel="stylesheet" />
<script src="/static/js/jquery.js" type="text/javascript"></script>
<script src="/static/js/jquery.unobtrusive-ajax.min.js" type="text/javascript"></script>
```

还需要定义一个提交表单的方法，

```javascript
function subForm() {
    $("#addForm").submit();
}
```

/Admin/ArticelClass/AddChildClass

在ArticelClass控制器中对应的AddChildClass方法中，完成子类别信息的添加

```go
//完成子类别添加
func (this *ArticelClassController)AddChildClass()  {
   o:=orm.NewOrm()
   cId,_:=this.GetInt("cId")
   className:=this.GetString("className")
   Remark:=this.GetString("Remark")
   var classInfo=models.ArticelClass{}
   userId:=this.GetSession("userId")
   if uId,ok:=userId.(int);ok{
      classInfo.CreateUserId=uId
   }
   classInfo.CreateDate=time.Now()
   classInfo.ClassName=className
   classInfo.ParentId=cId
   classInfo.DelFlag=0
   classInfo.Remark=Remark
    _,err:=o.Insert(&classInfo)
    if err==nil{
       this.Data["json"]=map[string]interface{}{"flag":"ok"}
   }else{
      this.Data["json"]=map[string]interface{}{"flag":"no"}
   }
   this.ServeJSON()


}
```



在上面的代码中，注意的是ParentId属性的值。

执行完添加后，会自动调用AJAX中指定的回调函数，afterAdd()

```javascript
function afterAdd(data) {
    window.parent.afterChiledAdd(data);//调用父窗体的方法
}
```

## 5:展示根类别下的子类别

在前面已经提到过，就是在整个页面中，默认只展示根类别信息，当单击根类别所在的行时，展示对应的子类别信息。

所以给treeeGrdid指定了单击行触发的事件。对应的会发送AJAX请求，将根类别的编号传递到服务端，获取该根类别下的子类别信息，最终添加到treegrid上，在添加之前，先将以前的内容删除掉，然后在添加。

这段代码不需要记忆，能看懂就可以，可以查文档。

```javascript
onClickRow: function (row) {
    //根据所单击的行，获取对应的子类别.
    $.post("/Admin/ArticelClass/ShowChildClass", { "id": row.Id }, function (data) {
        //先清空，后追加.如果没有数据不追加
        if (data.rows.length != 0) {
            var nodes = $('#tt').treegrid('getChildren', row.Id);
            for (var i = 0; i < nodes.length; i++) {
                $('#tt').treegrid('remove', nodes[i].id);
            }
            $('#tt').treegrid('append', {
                parent: row.Id,
                data: data.rows
            });
        }
    });
}
```

服务端处理：注意查询的条件

```go
//展示根类别下的子类别。
func (this *ArticelClassController)ShowChildClass()  {
  cId,_:=this.GetInt("id")
  o:=orm.NewOrm()
   var allClass []models.ArticelClass
  o.QueryTable("articel_class").Filter("parent_id",cId).All(&allClass)
  this.Data["json"]=map[string]interface{}{"rows":allClass}
  this.ServeJSON()
}
```



# 十五：新闻管理

## 1:模型设计

```go
//文章类别
type ArticelClass struct {
   Id int//主键
   ClassName string//类别名称
   ParentId int //父类别的编号
   CreateUserId int//创建类别的用户编号
   CreateDate time.Time//创建时间
   DelFlag int//删除标记
   Remark string//备注
   Artices []*ArticelInfo`orm:"reverse(many)"`
}
```



```go
//文章信息表
type ArticelInfo struct {
   Id int
   KeyWords string  //关键词
   Title string    //标题
   FullTitle string  //全标题
   Intro string  //导读
   ArticleContent string `orm:"type(text)"` //新闻内容
   Author string//作者
   Origin string//来源
   AddDate time.Time//添加日期
   ModifyDate time.Time//修改日期
   DelFlag int//删除标记
   PhotoUrl string//图片地址
   ArticelClasses []*ArticelClass`orm:"rel(m2m)"`
   
}
```

文章类别与文章信息表之间是什么关系？

## 2:新闻信息展示

### 2.1 前端处理

新闻信息的展示，与前面的设计是一样的。

在前端，创建一个dataGrid表格。

```javascript
function loadData() {
    $('#tt').datagrid({
        url: '/Admin/ArticelInfo/GetArticelInfo',
        title: '新闻数据表格',
        width: 700,
        height: 400,
        fitColumns: true, //列自适应
        nowrap: false,
        idField: 'Id',//主键列的列明
        loadMsg: '正在加载新闻的信息...',
        pagination: true,//是否有分页
        singleSelect: false,//是否单行选择
        pageSize: 5,//页大小，一页多少条数据
        pageNumber: 1,//当前页，默认的
        pageList: [5, 10, 15],
        queryParams: {},//往后台传递参数
        columns: [[//c.UserName, c.UserPass, c.Email, c.RegTime
            { field: 'ck', checkbox: true, align: 'left', width: 50 },
            { field: 'Id', title: '编号', width: 80 },
            { field: 'Title', title: '标题', width: 120 },
            { field: 'Author', title: '作者', width: 120 },
            { field: 'Origin', title: '来源', width: 120 },
            {
                field: 'AddDate', title: '时间', width: 80, align: 'right',
                formatter: function (value, row, index) {
                    return value.split("T")[0];
                }
            }

        ]],

        toolbar: [{
            id: 'btnDelete',
            text: '删除',
            iconCls: 'icon-remove',
            handler: function () {


            }
        }, {
            id: 'btnAdd',
            text: '添加',
            iconCls: 'icon-add',
            handler: function () {
                addArticel();//添加文章

            }
        }, {
            id: 'btnEdit',
            text: '编辑',
            iconCls: 'icon-edit',
            handler: function () {


            }
        }],
    });
}
```

需要在页面上添加一个table标签。

```html
<table id="tt" style="width: 700px;" title="标题，可以使用代码进行初始化，也可以使用这种属性的方式" iconcls="icon-edit">
```

同时在$(function{

})

中完成对loadData方法的调用。

### 2.2 服务端处理

在/Admin/ArticelInfo/GetArticelInfo对应的方法中完成数据的查询，生成JSON并返回。

```go
//获取列表示数据
func (this *ArticelInfoController)GetArticelInfo()  {
 pageIndex,_:=this.GetInt("page")
 pageSize,_:=this.GetInt("rows")
 start:=(pageIndex-1)*pageSize;
 o:=orm.NewOrm()
 var articels[]models.ArticelInfo
 o.QueryTable("articel_info").Filter("del_flag",0).OrderBy("id").Limit(pageSize,start).All(&articels)
 count,_:=o.QueryTable("articel_info").Filter("del_flag",0).Count()
 this.Data["json"]=map[string]interface{}{"rows":articels,"total":count}
 this.ServeJSON()

}
```

## 3:新闻信息的添加 

### 3.1 前端处理

在主页面中添加一个iframe,用来显示添加新闻的表单。

```html
<div id="addDiv">
    <iframe id="addFrame" frameborder="0" width="100%" height="100%"></iframe>
</div>
```

同时将上面的div隐藏。

```javascript
$("#addDiv").css("display","none");
```

接下来，给iframe指定src路径，并且弹出一个添加新闻的窗口。

```javascript
function addArticel() {
    $("#addFrame").attr("src","/Admin/ArticelInfo/ShowAddArticelInfo");
    $("#addDiv").css("display","block");
    $('#addDiv').dialog({
        title: '添加文章信息',
        width: 800,
        height: 700,
        collapsible: true,
        maximizable: true,
        resizable: true,
        modal: true,
        buttons: [{
            text: 'Ok',
            iconCls: 'icon-ok',
            handler: function () {
                //提交表单。
                var childWindows=$("#addFrame")[0].contentWindow;//获取子窗体的windows对象。

                childWindows.subForm()//提交表单

            }
        }, {
            text: 'Cancel',
            handler: function () {
                $('#addDiv').dialog('close');
            }
        }]
    });
}
```

在Admin/ArticelInfo/ShowAddArticelInfo对应的方法中,在指定视图页面之前，先将所有的类别信息查询出来。

在添加具体的新闻时可以选择对应的类别。

```go
func (this *ArticelInfoController)ShowAddArticelInfo()  {
    o:=orm.NewOrm()
    var articelClass []models.ArticelClass
    o.QueryTable("articel_class").Filter("parent_id__gte",1).All(&articelClass)
    this.Data["articelClass"]=articelClass
      this.TplName="ArticelInfo/ShowAddArticelInfo.html"
}
```

ShowAddArticelInfo.html

中展示展示添加的表单。

第一：引入所需要的js文件以及的css文件

```js
<script src="/static/js/jquery.js" type="text/javascript"></script>
<script src="/static/js/jquery.unobtrusive-ajax.min.js"></script>
<script src="/static/js/MyAjaxForm.js" type="text/javascript"></script>

<script src="/static/js/Ckeditor/ckeditor.js"></script>
<link href="/static/css/tableStyle.css" rel="stylesheet" />
```

第二：添加常用的CSS样式文件。

```css
<style type="text/css">
    textarea, select {
        padding: 2px;
        border: 1px solid;
        border-color: #666 #ccc #ccc #666;
        background: #F9F9F9;
        color: #333;
        resize: none;
        width: 100%;
    }

    .textbox {
        padding: 3px;
        border: 1px solid;
        border-color: #666 #ccc #ccc #666;
        background: #F9F9F9;
        color: #333;
        resize: none;
        width: 100%;
    }

    .textbox:hover, .textbox:focus, textarea:hover, textarea:focus {
        border-color: #09C;
        background: #F5F9FD;
    }
</style>
```

第三：添加表单。

```html
<body>
<div>
    <form  data-ajax="true" data-ajax-method="post" data-ajax-success="afterAdd" data-ajax-url="/Admin/ArticeInfo/AddArtice" id="form1">
        <table style="width:auto; margin: 0 auto">
            <tr>
                <td>简短标题:</td>

                <td>
                    <input type="text" name="title" maxlength="160" class="textbox" id="title" /></td>


            </tr>
            <tr>
                <td>完整标题:</td>
                <td colspan="4">
                    <input type="text" name="Fulltitle" class="textbox" id="Fulltitle" /></td>
            </tr>
            <tr>
                <td>归属栏目:</td>
                <td colspan="4">
                   {{range .articelClass}}
                        <input type="radio" value="{{.Id}}" name="className">{{.ClassName}}
                    {{end}}
                </td>
            </tr>
            <tr>
                <td>关 键 字</td>
                <td colspan="4">
                    <input type="text" name="KeyWords" class="textbox" id="KeyWords" /></td>
            </tr>
            <tr>
                <td>文章作者:</td>
                <td colspan="4">
                    <input type="text" name="Author" class="textbox" id="Author" />&nbsp;&nbsp;【<a href="#" class="authorClick">未知</a>】【<a href="#" class="authorClick">佚名</a>】【<a href="#" class="authorClick">Itcast</a>】</td>
            </tr>
            <tr>
                <td>文章来源:</td>
                <td colspan="4">
                    <input name="Origin" id="Origin"  value="" size="50" class="textbox" type="text">&nbsp;&nbsp;【<a href="#" class="originInfo">不详</a>】【<a href="#" class="originInfo">本站原创</a>】【<a href="#" class="originInfo">互联网</a>】</td>
            </tr>
            <tr>
                <td>文章导读</td>
                <td colspan="4">
                    <textarea class="textbox" name="Intro" style="width: 95%; height: 80px"></textarea></td>
            </tr>
            <tr>
                <td>文章内容</td>
                <td colspan="4">
                    <textarea class="ckeditor" onblur="getData2()" id="ArticleContent2" name="ArticleContent1" rows="30" cols="40" style="width: 95%; height: 80px"></textarea>
                    <script type="text/javascript">
                        //<![CDATA[
                        // Replace the <textarea id="editor1"> with an CKEditor instance.

                        var editor = CKEDITOR.replace('ArticleContent2');


                        //]]>
                    </script>
                    <textarea class="textbox"  name="ArticleContent" id="txtArticleContent" rows="30" cols="40" style="display:none"></textarea>
                </td>
            </tr>
            <tr>
                <td>图片地址:</td>
                <td colspan="4">
                    <input type="file" name="fileUp" />
                    <input type="button" value="上传图片" id="btnFileUp" />
                    <div id="showImage"></div>


                    <input name="PhotoUrl" id="PhotoUrl" value="" type="hidden">

                    <input name="InsertEditContent" id="InsertEditContent" value="1" type="checkbox">图片是否插入编辑器
                </td>
            </tr>
        </table>
    </form>
</div>
</div>

</body>
```

在上面的表单中，我们将对应的文章类别进行了循环遍历展示。

```go
{{range .articelClass}}
    <input type="radio" value="{{.Id}}" name="className">{{.ClassName}}
{{end}}
```



在上面的表单中，我们使用了CKEditor富文本编辑，通过该编辑器可以给输入的内容加粗，变红等操作。

那么怎样使用该富文本编辑器呢？

第一：导入ckeditor.js文件。

```js
<script src="/static/js/Ckeditor/ckeditor.js"></script>
```

第二：在表单中添加textarea元素，实际上CKEditor就是给textarea添加了很多的样式，效果。

第三：添加完textarea后，在其后面加上如下的js代码：

var editor = CKEDITOR.replace('ArticleContent2');

该代码的含义是使用CKEDITOR体会id为ArticleContent2的textarea元素。



为了方便输入，这里给一些标签加了单击事件，只要单击这些页签就可以直接将输入录入到文本框中。

```js
$(".authorClick").click(function () {
    $("#Author").val($(this).text());
});
//添加原创
$(".originInfo").click(function () {
    $("#Origin").val($(this).text());
});
```

### 3.2 图片文件上传

关于图片文件上传前面我们也已经实现过了。

第一:添加对应的js文件引入：

```js
<script src="/static/js/MyAjaxForm.js" type="text/javascript"></script>
```

第二：给上传按钮绑定单击事件。

```js
function bindFileUp() {
    $("#btnFileUp").click(function () {
        $("#form1").ajaxSubmit({
            success: function (data) {

                if (data.flag == "ok") {
                    var imgurl=data.msg.substr(1)
                    $("#showImage").html("<img src='" + imgurl + "' width='50px' height='50px'/>");
                    $("#PhotoUrl").val(imgurl);

                    //判断是否选择了"图片是否插入编辑器"
                    var flag = $("#InsertEditContent").is(":checked");
                    if (flag) {
                        var oEditor = CKEDITOR.instances.ArticleContent2;//找到编辑器
                        if (oEditor.mode == 'wysiwyg') {//what you see is what you get  所见即所得
                            var img = "<img src='"+imgurl+"'/>";
                            oEditor.insertHtml(img);//将上传成功的图片插入到编辑器中。
                        }
                        else
                            alert('You must be in WYSIWYG mode!');

                    }
                }
            },
            error: function (error) { alert(error); },
            url: '/Admin/ArticeInfo/FileUp', /*设置post提交到的页面*/
            type: "post", /*设置表单以post方法提交*/
            dataType: "json" /*设置返回值类型为文本*/
        });

    });
}
```

在上传成功后，会执行回调函数，在回调函数中，判断是否执行成功，如果执行成功，首先对返回的路径中的点（.）截取掉，同时组建一个<img>标签，插入到showImage整个div中，这时会在页面上看到用户上传成功的图片。同时将图片的路径存在隐藏域中，当提交表单时，将图片的地址提交到服务端，最终存储到数据库中。

这里与前面所讲的上传不同的地方是，这里需要判断是否选中了一个叫InsertEditContent的复选框，如果选中了该复选框还需要将上传成功的图片插入到富文本编辑器中。

在富文本编辑器中插入图片的过程如下：

第一：获取富文本编辑器。

 var oEditor = CKEDITOR.instances.ArticleContent2;//找到编辑器

CKEDITOR表示富文本编辑器，instances表示其对象，ArticleContent2：前面我们是给ArticleContent2这个textarea这个文本区域加上的富文本编辑器。

第二：判断其模式

CKEDITOR富文本编辑器有两种模式，一种是源码模式，通过该模式可以看到插入到编辑器中的内容的源代码，例如：给一段文字加粗使用的就是HTML中的<strong>标签。第二种模式，就是我们常用的模式就是编辑模式，也就是所见即所得模式 。只有在编辑模式下，才允许动态的向CKEDITOR中插入数据。

   if (oEditor.mode == 'wysiwyg'){
}

第三：进行图片的插入

这里需要构建一个img标签，然后调用insertHtml方法将图片插入到CKEDITOR中。

```js
if (oEditor.mode == 'wysiwyg') {//what you see is what you get  所见即所得
    var img = "<img src='"+imgurl+"'/>";
    oEditor.insertHtml(img);//将上传成功的图片插入到编辑器中。
}
```



这里使用的是insertHtml方法将图片的<img>标签插入到ckeditor编辑器中。

关于CKEDITOR编辑器的其它的方法，可以查看文档或者是案例。



服务端具体上传功能的实现：

对应的url地址为：/Admin/ArticeInfo/FileUp

具体上传的代码如下：

```go
func (this *ArticelInfoController)FileUp()  {
   f,h,err:=this.GetFile("fileUp")
   if err!=nil{
      this.Data["json"]=map[string]interface{}{"flag":"no","msg":"文件上传失败!!"}
   }else{
      //获取文件名称
      fileName:=h.Filename
      //获取扩展名
      fileExt:=path.Ext(fileName)
      if fileExt!=".jpg"||fileExt!=".png" {
         //获取上传文件的大小
         fileSize:=h.Size
         if fileSize<50000000 {
            //构建存储的目录
            dir:="./static/fileUp/"+strconv.Itoa(time.Now().Year())+"/"+time.Now().Month().String()+"/"+strconv.Itoa(time.Now().Day())+"/"
            _,err:=os.Stat(dir)
            if err!=nil{//表示没有文件目录
               os.MkdirAll(dir,os.ModePerm)
            }
            //文件重名
            newFileName:=strconv.Itoa(time.Now().Year())+time.Now().Month().String()+strconv.Itoa(time.Now().Day())+strconv.Itoa(time.Now().Hour())+strconv.Itoa(time.Now().Minute())+strconv.Itoa(time.Now().Minute())+strconv.Itoa(time.Now().Second())
            //构建完成路径
            fullDir:=dir+newFileName+fileExt
            err1:=this.SaveToFile("fileUp",fullDir)//保存文件
            if err1==nil {
               this.Data["json"]=map[string]interface{}{"flag":"ok","msg":fullDir}
            }else{
               this.Data["json"]=map[string]interface{}{"flag":"no","msg":"文件上传失败!!"}
            }

         }else{
            this.Data["json"]=map[string]interface{}{"flag":"no","msg":"文件类太大!!"}
         }
      }else{
         this.Data["json"]=map[string]interface{}{"flag":"no","msg":"文件类型错误!!"}
      }
   }
   this.ServeJSON()

   defer  f.Close()
}
```

对于上传文件的具体实现，前面已经讲解过，这里不在做过多的解释。

### 3.3: 服务端处理

在将表单中的数据提交到服务端前，还需要做一些处理。在服务端是无法直接接受到CKEditor中的内容的，所以表单中又添加了一个textarea元素。

```html
<textarea class="textbox"  name="ArticleContent" id="txtArticleContent" rows="30" cols="40" style="display:none"></textarea>
```

目的是，在CKEDITOR中输入的内容自动的填充到id为txtArticleContent的textarea中，在服务端接收的是该textarea的值。

所以，下面我们要给CKEDITOR绑定一个事件为blur事件，该事件是在失去焦点的时候出发。

具体的实现方式如下：

```js
CKEDITOR.instances["ArticleContent2"].on("blur", function () {
    //获取编辑器内容。
    var oEditor = CKEDITOR.instances.ArticleContent2;//找到编辑器
    var content = oEditor.getData();//获取编辑器中的数据
    $("#txtArticleContent").val(content);//将获取到的值插入到txtArticleContent整个文本域中
    // console.log($("#txtArticleContent").val())

});
```

上面代码的含义是：首先找到编辑器，然后获取编辑器中的内容数据（getData()函数的作用就是获取数据）

将获取到的数据插入到文本域中。

接下来当提交表单时，服务端就可以进行处理了：

/Admin/ArticeInfo/AddArtice

在AddArtice方法中，接收到表单提交过来的数据，保存到数据库中。

```go
 var articelInfo=models.ArticelInfo{}
 articelInfo.DelFlag=0
 articelInfo.AddDate=time.Now()
 articelInfo.ArticleContent=this.GetString("ArticleContent")
 articelInfo.Origin=this.GetString("Origin")
 articelInfo.Title=this.GetString("title")
 articelInfo.PhotoUrl=this.GetString("PhotoUrl")
 articelInfo.KeyWords=this.GetString("KeyWords")
 articelInfo.Intro=this.GetString("Intro")
 articelInfo.Author=this.GetString("Author")
 articelInfo.FullTitle=this.GetString("FullTitle")
 articelInfo.ModifyDate=time.Now()
 o:=orm.NewOrm()
 num,_:=o.Insert(&articelInfo)
//获取类别编号
 classId,_:=this.GetInt("className")
 //查询类别的信息。
 var classInfo models.ArticelClass
 o.QueryTable("articel_class").Filter("id",classId).One(&classInfo)
//创建M2M对象
m2m:=o.QueryM2M(&articelInfo,"ArticelClasses")
m2m.Add(classInfo)
```

注意：这里需要保存新闻类别信息，也就是刚插入的新闻是属于哪个类别的，并且由于类别与新闻是多对多的关系，所以这里需要创建m2m对象来进行保存新闻对应的类别信息。



### 4:生成静态页面

大家可以可以考虑下，像新闻的详细页面，小说的页面等，有什么样的特点？

这些页面显示的文字内容比较多，同时发现这些内容并不是经常修改，并且经常被访问。

如果直接将这些内容从数据库中查询，效率会比较低。

所以，我们可以将这些经常查询，但是不是经常修改，内容量比较大的页面，做成静态页面。

所谓的静态页面，就是.html页面，页面中展示的新闻，小说内容等数据都是写死在页面上的。其实就是大家前面学的html内容。

具体的制作过程是：

第一：先指定模板文件（模板文件就是一个html网页，该模板文件不需要我们自己设计，有专门的前端人员进行设计）。该模板文件完成了整个页面的布局，并且在该页面中有占位符。这些占位符需要被具体的数据替换。所谓占位符，其实就是要用具体数据填充的位置，这里先放上一个符号（这个符号大家可以随便定义），后面用具体的数据替换掉就可以了。

第二：当将新闻的数据插入到数据库后，读取该模板文件，并且用具体的数据替换掉模板文件中的占位符。

接下来，我们实现一下具体的代码：

在完成新闻数据的插入后，立即调用生成静态页面的方法：

```go
CreateStaticPage(int(num))//生成静态页面
```

该方法的具体实现如下

在看具体代码之前，先说一下具体的思路：

1：根据传递过来的刚插入的新闻编号，查询出具体的新闻信息。

2：指定模板路径，并且读取模板文件。

3:对模板文件中的占位符，用具体的数据进行替换。

4：根据日期创建文件夹（这里的日期，我们使用添加新闻的日期）。在前面我们也使用这种方式来创建文件夹，但是月份得到的是英文的大写。

  所以，接下来我们要替换成具体的数字。

5：完成文件夹的创建。

6：构建完整的路径

7：完成文件写入操作。

```go
func CreateStaticPage(aId int )  {
   //1:根据传递过来的文章编号，查询对应的文章信息
   var articelInfo models.ArticelInfo
    o:=orm.NewOrm()
    o.QueryTable("articel_info").Filter("id",aId).One(&articelInfo)
    dir:="./static/ArticelTemplate/ArticelTemplateInfo.html"
    file,err:=os.Open(dir)
    defer  file.Close()
    if err!=nil{
       beego.Info("文件打开失败")
   }else{
      //读取打开文件中的内容。
      content,_:=ioutil.ReadAll(file)
      //将读取的内容转换成文本字符串。
       articeContent:=string(content)
      articeContent=strings.Replace(articeContent,"$Title",articelInfo.Title,-1)
      articeContent=strings.Replace(articeContent,"$Origin",articelInfo.Origin,-1)
      articeContent=strings.Replace(articeContent,"$ArticleContent",articelInfo.ArticleContent,-1)
      articeContent=strings.Replace(articeContent,"$AddDate",articelInfo.AddDate.Format("2006-01-02"),-1)
      //创建文件夹。
      month:=articelInfo.AddDate.Month().String()
      var m int
      for i:=0;i<len(months) ; i++ {
         if months[i]==month {
            m=i;
            break;
         }
      }
      m=m+1
      var dirDict string
      if m<10{
         md:="0"+strconv.Itoa(m)
         dirDict="./static/Articel/"+strconv.Itoa(articelInfo.AddDate.Year())+"/"+md+"/"+strconv.Itoa(articelInfo.AddDate.Day())+"/"
      }else{
         dirDict="./static/Articel/"+strconv.Itoa(articelInfo.AddDate.Year())+"/"+strconv.Itoa(m)+"/"+strconv.Itoa(articelInfo.AddDate.Day())+"/"
      }
      _,err:=os.Stat(dirDict)
      if err!=nil{
         os.MkdirAll(dirDict,os.ModePerm)
      }
      fullDir:=dirDict+strconv.Itoa(articelInfo.Id)+".html"
      if ioutil.WriteFile(fullDir,[]byte(articeContent),0644)==nil{
         beego.Info("写入文件成功")
      }

   }



}
//定义日期切片
var months = []string{
   "January",
   "February",
   "March",
   "April",
   "May",
   "June",
   "July",
   "August",
   "September",
   "October",
   "November",
   "December",
}
```

创建完成后名，怎样访问静态页面呢？可以在主页中的表格中添加链接

，可以在datagrid中加上一列。对返回的日期格式进行处理，构建响应的路径。

```js
{
    field: 'showDetail', title: '详细', width: 80, align: 'right',
    formatter: function (value, row, index) {
        var d=row.AddDate.split("T")[0];
        var year=d.split("-")[0];
        var month=d.split("-")[1];
        var day=d.split("-")[2];
        url=year+"/"+month+"/"+day;
       var str = "<a href='#' ids='" + row.Id + "' class='details' showAddData='" +url + "'>详细</a>";
            return str;
    }
}
```

并且给datagrid表格加上一个事件为onLoadSuccess，这个事件指是当表格中的数据全部加载完成后触发。

```js
//载入成功以后触发
onLoadSuccess: function () {

    $(".details").click(function () {
        var articelId = $(this).attr("ids");
        var showAddData = $(this).attr("showAddData");//2015/1/12
        var dir = "/static/Articel/"+showAddData+"/"+articelId+".html";
        window.open(dir);
    });
}
```

当表格中的数据全部加载完后，会触发该事件，那么可以获取添加完的“详细”按钮链接，并且绑定单击事件。

然后构建完整路径。

### 5：发布评论

模型设计：

```go
//文章评论
type ArticelComment struct {
   Id int `from:"-"`
   Msg string
   AddDate time.Time
   IsPass int
   Articel *ArticelInfo `orm:"rel(fk)"`
}
```



```go
//文章信息表
type ArticelInfo struct {
   Id int
   KeyWords string  //关键词
   Title string    //标题
   FullTitle string  //全标题
   Intro string  //导读
   ArticleContent string `orm:"type(text)"` //新闻内容
   Author string//作者
   Origin string//来源
   AddDate time.Time//添加日期
   ModifyDate time.Time//修改日期
   DelFlag int//删除标记
   PhotoUrl string//图片地址
   ArticelClasses []*ArticelClass`orm:"rel(m2m)"`
   Comments[]*ArticelComment `orm:"reverse(many)"`
}
```

思考：评论与文章的关系。



在模板文件中，单击发布评论按钮时，将文章的编号，和用户输入的评论内容，通过AJAX的方式发送到服务端，服务端接收后，添加到数据库中。

给发布评论按钮绑定单击事件：将评论内容以及文章的编号发送到服务端。思考：怎样获取文章编号呢？

```js
function addComment() {
    $("#btnComment").click(function () {
        var msg=$("#commentMsg").val();
        if(msg!=""){
            $.post("/Admin/ArticelInfo/AddMsg",{"msg":msg,"articelId":$articelId},function (data) {
                if(data.flag=="ok"){
                    loadComment();//添加完成后，重新调用该方法。
                }else{
                    $.messager.alert("提示","评论错误!!","error")
                }

            })
        }else{
            $.messager.alert("提示","评论内容不能为空!!","error")

        }
    })
}
```

服务端处理：

```go
//添加评论内容
func (this *ArticelInfoController)AddMsg()  {
   msg:=this.GetString("msg")
   articelId,_:=this.GetInt("articelId")
   var articelComment=models.ArticelComment{}
   articelComment.AddDate=time.Now()
   articelComment.IsPass=1
   articelComment.Msg=msg
   o:=orm.NewOrm()
   var articelInfo models.ArticelInfo
   o.QueryTable("articel_info").Filter("id",articelId).One(&articelInfo)
   articelComment.Articel=&articelInfo
   _,err:=o.Insert(&articelComment)
   if err==nil{
      this.Data["json"]=map[string]interface{}{"flag":"ok"}
   }else{
      this.Data["json"]=map[string]interface{}{"flag":"no"}
   }
   this.ServeJSON()
}
```

评论的加载：

前端处理：将文章的编号发送到服务端，看一下该文章下面有多少评论。然后将服务端返回的评论数据添加到

“<li>”标签上，最终追加到<ul>上。

```js
//加载评论
function loadComment() {
    $.post("/Admin/ArticelInfo/LoadMsg",{"articelId":$articelId},function (data) {
        var  serverDataLength=data.msg.length;
        for (var i=0;i<serverDataLength;i++){
            $("<li>"+data.msg[i].AddDate.split("T")[0]+":"+data.msg[i].Msg+ "</li>").appendTo("#articelCommentList")
        }

    })
}
```

服务端处理：

根据传递过来的文章编号，查询出对应的评论内容：

```go
func (this *ArticelInfoController)LoadMsg()  {
   articelId,_:=this.GetInt("articelId")
   beego.Info(articelId)
   o:=orm.NewOrm()
   var comments[]models.ArticelComment
   o.QueryTable("articel_comment").Filter("articel_id",articelId).All(&comments)
   this.Data["json"]=map[string]interface{}{"msg":comments}
   this.ServeJSON()
}
```

