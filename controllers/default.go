package controllers

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"radiusweb/libs"
	"radiusweb/models"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

//默认的get和post方法
func (c *MainController) Get() {
	c.Redirect("/login", 302)
}
func (c *MainController) Post() {
}

//登入get页面方法
func (c *MainController) Login() {
	c.TplName = "login.html"
}

//登入post页面方法
func (c *MainController) PostLogin() {
	type RET struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	var ret RET

	Name1 := c.GetString("username")
	Passwd1 := c.GetString("password")

	o := orm.NewOrm()
	admin := models.WsAdmin{Name: Name1}
	err := o.Read(&admin, "Name")

	if err != nil || admin.Password != Passwd1 {
		beego.Error(err)
		ret.Code = 1
		ret.Msg = "用户名密码不符"
	} else {

		ret.Code = 0
		ret.Msg = "ok"
		//设置session
		c.SetSession("username", Name1)
	}
	//对应view/login.html中的js doLogin()函数
	b, err := json.Marshal(ret)
	if err == nil {
		c.Ctx.WriteString(string(b))
	} else {
		c.Ctx.WriteString("{code:1,msg:\"JSON ERROR\"}")
	}
	c.TplName = "login.html"
}

//主页方法
func (c *MainController) Getindex() {
	c.Layout = "layout_base.html"
	c.TplName = "template_ui.html"

}

//关于about
func (c *MainController) ShowAbout() {

	c.Layout = "layout_base.html"
	// if c.LayoutSections == nil {
	// 	c.LayoutSections = make(map[string]string)
	// }
	// c.LayoutSections["ui"] = "template_ui.html"
	c.TplName = "about.html"
}

//修改密码
func (c *MainController) ShowChangePassword() {
	c.Layout = "layout_base.html"

	c.TplName = "password.html"
}

func (c *MainController) PostChangePassword() {
	c.Layout = "layout_base.html"
	if c.LayoutSections == nil {
		c.LayoutSections = make(map[string]string)
	}
	c.LayoutSections["HeadCss"] = "password_head.html"
	c.TplName = "password.html"
	p1 := c.GetString("tr_user_pass")
	p2 := c.GetString("tr_user_pass_chk")
	if models.AdminUpdate(p1, p2) == "ok" {
		c.Data["T"] = "更改成功"
	} else {
		c.Data["T"] = "更改失败，请重新更改"
	}

}

//退出登入
func (c *MainController) LogOut() {
	c.DelSession("username")  //删除session
	c.Redirect("/login", 302) //重定向到登入界面

}

//操作日志界面
func (c *MainController) ShowLog() { //get请求
	c.Layout = "layout_base.html"
	c.TplName = "log.html"
	pno, _ := c.GetInt("pno") //获取当前请求页
	var ShowLog []models.WsLog
	var conditions []string = make([]string, 2) //定义日志查询条件,格式为 " and name='zhifeiya' and age=12 "
	var po models.PageOptions                   //定义一个分页对象
	po.TableName = "ws_log"                     //指定分页的表名
	po.EnableFirstLastLink = true               //是否显示首页尾页 默认false
	po.EnablePreNexLink = true                  //是否显示上一页下一页 默认为false
	po.Conditions = conditions                  // 传递分页条件 默认全表
	po.Currentpage = int(pno)                   //传递当前页数,默认为1
	po.PageSize = 10                            //页面大小  默认为20

	//返回分页信息,
	//第一个:为返回的当前页面数据集合,ResultSet类型
	//第二个:生成的分页链接
	//第三个:返回总记录数
	//第四个:返回总页数
	totalItem, _, rs, pagerhtml := models.GetPagerLinks(&po, c.Ctx)
	rs.QueryRows(&ShowLog)      //把当前页面的数据序列化进一个切片内
	c.Data["ShowLog"] = ShowLog //把当前页面的数据传递到前台
	c.Data["pagerhtml"] = pagerhtml
	c.Data["totalItem"] = totalItem

}

func (c *MainController) PostLog() { //post请求
	c.Layout = "layout_base.html"
	c.TplName = "log.html"
}

//系统服务页

func (c *MainController) ShowSuperrpc() {
	c.Layout = "layout_base.html"
	c.TplName = "wait.html"
}

//控制面板
func (c *MainController) ShowDashboard() {
	c.Layout = "layout_base.html"
	if c.LayoutSections == nil {
		c.LayoutSections = make(map[string]string)
	}
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["CPU"] = c.GetCpu()
	c.Data["ServerTime"] = c.GetServerTime()
	c.Data["MemInfo"] = c.GetMemInfo()
	c.Data["Disk"] = c.GetDiskUsage()

	c.TplName = "dashboard.html"
}

//服务器时间
func (c *MainController) GetServerTime() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}
func (c *MainController) GetCliOutput(cmd string) string {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return fmt.Sprintf("Failed to execute command: %s", cmd)
	}
	return string(out)
}

//内存信息
func (c *MainController) GetMemInfo() [][]string {
	cmd := "cat /proc/meminfo"
	tmp := c.GetCliOutput(cmd)
	var ret [][]string
	arr := strings.Split(tmp, "\n")
	for _, v := range arr {
		if strings.HasPrefix(v, "MemTotal") || strings.HasPrefix(v, "MemFree") || strings.HasPrefix(v, "Cached") {
			//v = strings.Replace(v,"  ","",-1)
			//fmt.Println(v)
			fields := strings.Fields(v)
			a := []string{fields[0], fields[1]}
			ret = append(ret, a)
		}
	}

	return ret
}

//CPU信息
func (c *MainController) GetCpu() string {

	cmd := "cat /proc/cpuinfo | egrep '^model name' | uniq | awk '{print substr($0, index($0,$4))}'"
	return c.GetCliOutput(cmd)
}

//磁盘信息
func (c *MainController) GetDiskUsage() *libs.DiskStatus {
	dk := libs.DiskUsage("/")
	return &dk
}
