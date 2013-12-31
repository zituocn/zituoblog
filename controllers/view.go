package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"zituoblog/models"
)

type ViewHandel struct {
	beego.Controller
}

func (this *ViewHandel) Get() {
	this.TplNames = "_view.html"

	id := this.Ctx.Input.Params[":id"]
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
		return
	}
	info, err := models.GetNewsInfo(tid)
	if err != nil {
		this.Redirect("/", 302)
	}

	this.Data["info"] = info
}
