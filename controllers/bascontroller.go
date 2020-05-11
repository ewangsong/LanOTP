package controllers

import (
	"radiusweb/models"

	"github.com/astaxie/beego"
)

type BasController struct {
	beego.Controller
}

//bas方法
func (c *BasController) Getbas() {
	c.Layout = "layout_base.html"
	c.TplName = "bas.html"
	ShowBas := models.BasRead()

	c.Data["ShowBas"] = ShowBas
}

//添加bas
func (c *BasController) AddBas() {
	c.Layout = "layout_base.html"
	c.TplName = "bas_add.html"

}
func (c *BasController) PostAddBas() {
	name := c.GetString("bas_name")
	ip_addr := c.GetString("ip_addr")
	secret := c.GetString("bas_secret")
	port := c.GetString("coa_port")

	if models.BasInsert(name, ip_addr, secret, port) {
		beego.Info(name)
		c.Redirect("/admin/bas", 302)
	} else {
		beego.Info("添加错误，请重新添加")
	}
	c.Layout = "layout_base.html"
	c.TplName = "bas_add.html"
}

//更新bas
func (c *BasController) UpdateBas() {
	c.Layout = "layout_base.html"
	c.TplName = "bas_update.html"
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("获取basID错误", err)
		return
	}
	c.Data["bas1"] = models.BasRead2(id)
}
func (c *BasController) PostUpdateBas() {
	c.Layout = "layout_base.html"
	c.TplName = "bas_update.html"
	id, err := c.GetInt("id")
	beego.Info(id)
	if err != nil {
		beego.Info("获取basID错误", err)
		return
	}
	name := c.GetString("bas_name")
	ip_addr := c.GetString("ip_addr")
	secret := c.GetString("bas_secret")
	port := c.GetString("coa_port")
	models.BasUpdate(id, name, ip_addr, secret, port)

	c.Redirect("/admin/bas", 302)

}

//删除bas
func (c *BasController) DeleteBas() {
	c.Layout = "layout_base.html"
	c.TplName = "bas.html"
	id, err := c.GetInt("id")
	beego.Info(id)
	if err != nil {
		beego.Info("获取删除basID错误", err)
		return
	}
	models.BsaDelete(id)
	c.Redirect("/admin/bas", 302)
}
