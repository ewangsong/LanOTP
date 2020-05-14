package controllers

import (
	"encoding/json"
	"ewangsong/LanOTP/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

//TokenController 令牌Controller
type TokenController struct {
	beego.Controller
}

//ShowToken 令牌列表
func (c *TokenController) ShowToken() {
	c.Layout = "layout_base.html"
	if c.LayoutSections == nil {
		c.LayoutSections = make(map[string]string)
	}
	//c.LayoutSections["HeadCss"] = "users_head.html"
	c.TplName = "token_list.html"
	// //定义显示记录和页数

	pno, _ := c.GetInt("pno") //获取当前请求页
	var tokens []models.WsOtp
	var conditions string         //定义日志查询条件,格式为 " and name='zhifeiya' and age=12 "
	var po models.PageOptions     //定义一个分页对象
	po.TableName = "ws_otp"       //指定分页的表名
	po.EnableFirstLastLink = true //是否显示首页尾页 默认false
	po.EnablePreNexLink = true    //是否显示上一页下一页 默认为false
	po.Conditions = conditions    // 传递分页条件 默认全表
	po.Currentpage = int(pno)     //传递当前页数,默认为1
	po.PageSize = 20              //页面大小  默认为20
	//返回分页信息,
	//第一个:为返回的当前页面数据集合,ResultSet类型
	//第二个:生成的分页链接
	//第三个:返回总记录数
	//第四个:返回总页数
	totalItem, _, rs, pagerhtml := models.GetPagerLinks(&po, c.Ctx)
	rs.QueryRows(&tokens)     //把当前页面的数据序列化进一个切片内
	c.Data["tokens"] = tokens //把当前页面的数据传递到前端
	c.Data["pagerhtml"] = pagerhtml
	c.Data["totalItem"] = totalItem
	c.Data["PageSize"] = po.PageSize

}

//PostShowToken 令牌列表查询
func (c *TokenController) PostShowToken() {

	c.Layout = "layout_base.html"
	c.TplName = "token_list.html"

	tokensn := c.GetString("tokensn")
	username := c.GetString("username")

	var conditions string
	if tokensn == "" && username == "" {
		c.Redirect("/admin/token", 302)
	}
	if tokensn == "" {
		conditions = "AND binding_user='" + username + "'"
	} else if username == "" {
		conditions = "AND otp_sn='" + tokensn + "'"

	} else {
		conditions = "AND otp_sn='" + tokensn + "' AND binding_user='" + username + "'"
	}

	pno, _ := c.GetInt("pno") //获取当前请求页
	var tokens []models.WsOtp
	//var conditions []string      //定义日志查询条件,格式为 " and name='zhifeiya' and age=12 "
	var po models.PageOptions     //定义一个分页对象
	po.TableName = "ws_otp"       //指定分页的表名
	po.EnableFirstLastLink = true //是否显示首页尾页 默认false
	po.EnablePreNexLink = true    //是否显示上一页下一页 默认为false
	po.Conditions = conditions    // 传递分页条件 默认全表
	po.Currentpage = int(pno)     //传递当前页数,默认为1
	po.PageSize = 10              //页面大小  默认为20
	fmt.Println(conditions)
	//返回分页信息,
	//第一个:为返回的当前页面数据集合,ResultSet类型
	//第二个:生成的分页链接
	//第三个:返回总记录数
	//第四个:返回总页数
	totalItem, _, rs, pagerhtml := models.GetPagerLinks(&po, c.Ctx)
	rs.QueryRows(&tokens)     //把当前页面的数据序列化进一个切片内
	c.Data["tokens"] = tokens //把当前页面的数据传递到前台
	c.Data["pagerhtml"] = pagerhtml
	c.Data["totalItem"] = totalItem
}

//编辑用户
func (c *TokenController) DetailToken() {
	c.Layout = "layout_base.html"
	c.TplName = "token_detail.html"

	id, err := c.GetInt("token_id")

	if err != nil {
		beego.Info("获取用户ID错误", err)
		return
	}
	c.Data["token"] = models.TokenRead(id)

}

//更改token
func (c *TokenController) UpdateToken() {
	c.Layout = "layout_base.html"
	c.TplName = "token_update.html"

	id, err := c.GetInt("token_id")
	if err != nil {
		beego.Info("获取用户ID错误", err)
		return
	}
	c.Data["token"] = models.TokenRead(id)

}
func (c *TokenController) PostUpdateToken() {
	c.Layout = "layout_base.html"
	c.TplName = "token_update.html"
	id, err := c.GetInt("token_id")
	if err != nil {
		beego.Info("获取用户ID错误", err)
		return
	}
	name := c.GetString("name")
	models.TokenUdate(id, name)
	url := c.Ctx.Input.URI()
	c.Redirect(url, 302)
}

//删除token
func (c *TokenController) DeleteToken() {
	c.Layout = "layout_base.html"
	c.TplName = "token_list.html"
	id, err := c.GetInt("token_id")
	if err != nil {
		beego.Info("获取用户ID错误", err)
		return
	}
	models.TokenDelete(id)
	c.Redirect("/admin/token", 302)
}

//添加token

func (c *TokenController) AddToken() {
	c.Layout = "layout_base.html"
	if c.LayoutSections == nil {
		c.LayoutSections = make(map[string]string)
	}
	c.LayoutSections["HeadCss"] = "token_head.html"
	c.TplName = "token_add.html"

}
func (c *TokenController) PostAddToken() {
	type RET struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Url  string `json:"url"`
	}
	var ret RET

	c.Layout = "layout_base.html"
	c.TplName = "token_add.html"

	typeid := c.GetString("typeid")
	username := c.GetString("username")
	tid, _ := strconv.Atoi(typeid)

	tokenid, bo := models.AddToken(username, tid)
	if bo {
		ret.Code = 0
		ret.Msg = "ok"
		ret.Url = "/admin/token/detail?token_id=" + strconv.Itoa(tokenid) //跳转到指定token绑定用户界面
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
