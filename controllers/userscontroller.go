package controllers

import (
	"encoding/json"
	"radiusweb/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

//用户列表
func (c *UserController) ShowUsers() {
	c.Layout = "layout_base.html"
	if c.LayoutSections == nil {
		c.LayoutSections = make(map[string]string)
	}
	//c.LayoutSections["HeadCss"] = "users_head.html"
	c.TplName = "users_list.html"
	// ShowUsers, count := models.UsersRead()

	// //定义显示记录和页数

	// c.Data["ShowUsers"] = ShowUsers
	// c.Data["Count"] = count
	// c.Data["PageCount"] = PageCount

	pno, _ := c.GetInt("pno") //获取当前请求页
	var ShowUsers []models.WsUsers
	var conditions []string = make([]string, 2) //定义日志查询条件,格式为 " and name='zhifeiya' and age=12 "
	var po models.PageOptions                   //定义一个分页对象
	po.TableName = "ws_users"                   //指定分页的表名
	po.EnableFirstLastLink = true               //是否显示首页尾页 默认false
	po.EnablePreNexLink = true                  //是否显示上一页下一页 默认为false
	po.Conditions = conditions                  // 传递分页条件 默认全表
	po.Currentpage = int(pno)                   //传递当前页数,默认为1
	po.PageSize = 20                            //页面大小  默认为20
	//返回分页信息,
	//第一个:为返回的当前页面数据集合,ResultSet类型
	//第二个:生成的分页链接
	//第三个:返回总记录数
	//第四个:返回总页数
	totalItem, _, rs, pagerhtml := models.GetPagerLinks(&po, c.Ctx)
	rs.QueryRows(&ShowUsers)        //把当前页面的数据序列化进一个切片内
	c.Data["ShowUsers"] = ShowUsers //把当前页面的数据传递到前台
	c.Data["pagerhtml"] = pagerhtml
	c.Data["totalItem"] = totalItem
	c.Data["PageSize"] = po.PageSize

}

//用户列表查询用户
func (c *UserController) PostShowUsers() {

	c.Layout = "layout_base.html"
	c.TplName = "users_list.html"

	var conditions []string = make([]string, 2)
	name := c.GetString("name")
	realname := c.GetString("realname")
	if name == "" && realname == "" {
		c.Redirect("/admin/users", 302)
	}

	if name == "" {
		//c.Data["ShowUsers"] = models.UsersRead4(realname)
		//realname = " and " + "real_name=" + `"` + realname + `"`
		conditions[1] = "%" + realname + "%"
		conditions[0] = "where real_name like ?"
	} else if realname == "" {
		// c.Data["ShowUsers"] = models.UsersRead3(name)
		conditions[1] = "%" + name + "%"
		conditions[0] = "where name like ?"

	}

	pno, _ := c.GetInt("pno") //获取当前请求页
	var ShowUsers []models.WsUsers
	//var conditions []string      //定义日志查询条件,格式为 " and name='zhifeiya' and age=12 "
	var po models.PageOptions     //定义一个分页对象
	po.TableName = "ws_users"     //指定分页的表名
	po.EnableFirstLastLink = true //是否显示首页尾页 默认false
	po.EnablePreNexLink = true    //是否显示上一页下一页 默认为false
	po.Conditions = conditions    // 传递分页条件 默认全表
	po.Currentpage = int(pno)     //传递当前页数,默认为1
	po.PageSize = 10              //页面大小  默认为20

	//返回分页信息,
	//第一个:为返回的当前页面数据集合,ResultSet类型
	//第二个:生成的分页链接
	//第三个:返回总记录数
	//第四个:返回总页数
	totalItem, _, rs, pagerhtml := models.GetPagerLinks(&po, c.Ctx)
	rs.QueryRows(&ShowUsers)        //把当前页面的数据序列化进一个切片内
	c.Data["ShowUsers"] = ShowUsers //把当前页面的数据传递到前台
	c.Data["pagerhtml"] = pagerhtml
	c.Data["totalItem"] = totalItem
}

//编辑用户
func (c *UserController) DetailUsers() {
	c.Layout = "layout_base.html"
	c.TplName = "users_detail.html"

	id, err := c.GetInt("user_id")

	if err != nil {
		beego.Info("获取用户ID错误", err)
		return
	}
	c.Data["user"] = models.UsersRead2(id)

}

//更改user
func (c *UserController) UpdateUsers() {
	c.Layout = "layout_base.html"
	c.TplName = "users_update.html"

	id, err := c.GetInt("user_id")
	if err != nil {
		beego.Info("获取用户ID错误", err)
		return
	}
	c.Data["user1"] = models.UsersRead2(id)

}
func (c *UserController) PostUpdateUsers() {
	c.Layout = "layout_base.html"
	c.TplName = "users_update.html"
	id, err := c.GetInt("user_id")
	if err != nil {
		beego.Info("获取用户ID错误", err)
		return
	}
	name := c.GetString("name")
	realname := c.GetString("realname")
	password := c.GetString("new_password")
	models.UserUdate(id, realname, name, password)
	url := c.Ctx.Input.URI()
	c.Redirect(url, 302)
}

//删除用户
func (c *UserController) DeleteUser() {
	c.Layout = "layout_base.html"
	c.TplName = "users_list.html"
	id, err := c.GetInt("user_id")
	if err != nil {
		beego.Info("获取用户ID错误", err)
		return
	}
	models.UserDelete(id)
	c.Redirect("/admin/users", 302)
}

//添加用户

func (c *UserController) AddUser() {
	c.Layout = "layout_base.html"
	if c.LayoutSections == nil {
		c.LayoutSections = make(map[string]string)
	}
	c.LayoutSections["HeadCss"] = "users_head.html"
	c.TplName = "user_add.html"

}
func (c *UserController) PostAddUser() {
	type RET struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Url  string `json:"url"`
	}
	var ret RET

	c.Layout = "layout_base.html"
	c.TplName = "user_add.html"
	realname := c.GetString("realname")
	name := c.GetString("username")
	password := c.GetString("password")

	if models.UserInsert(realname, name, password) {
		url := models.UsersRead5(name)
		ret.Code = 0
		ret.Msg = "ok"
		ret.Url = url

	} else {
		ret.Code = 1
		ret.Msg = "账号已存在或者输入错误"
	}

	b, err := json.Marshal(ret)
	if err == nil {
		c.Ctx.WriteString(string(b))
	} else {
		c.Ctx.WriteString("{code:1,msg:\"JSON ERROR\"}")
	}

}
