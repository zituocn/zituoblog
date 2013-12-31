package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"zituoblog/common"
	"zituoblog/models"
)

type CategoryHandel struct {
	beego.Controller
}

func (this *CategoryHandel) Get() {
	this.TplNames = "_categroy.html"

	ename := this.Ctx.Input.Params[":ename"]
	pagestr := this.Ctx.Input.Params[":page"]
	page, _ := strconv.ParseInt(pagestr, 10, 64)
	if page < 1 {
		page = 1
	}

	classinfo, err := models.GetClassInfo(ename)
	if err != nil {
		this.Redirect("/", 302)
	}

	var pagesize int64
	pagesize = 20

	list := []string{}
	var item string
	offset := (page - 1) * pagesize
	newslist, count := models.GetNewsListByCid(pagesize, offset, classinfo.Id)
	list = append(list, "<ul>")
	for _, newsinfo := range newslist {
		item = item + fmt.Sprintf("<li><a href=\"/view/%d\" target=\"_blank\">%s</a></li>", newsinfo.Id, newsinfo.Title)
	}
	list = append(list, item)
	list = append(list, "</ul>")

	pager := common.PageList(pagesize, page, count, false, "/category/"+classinfo.Ename)
	this.Data["pager"] = pager
	this.Data["classinfo"] = classinfo
	this.Data["ListString"] = strings.Join(list, "\n")
}
