package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strconv"
	"zituoblog/common"
	"zituoblog/models"
)

type AdminTestHandel struct {
	beego.Controller
}

//退出管理
type AdminLogoutHandel struct {
	beego.Controller
}

//管理登录
type AdminLoginHandel struct {
	beego.Controller
}

//管理界面
type AdminMainHandel struct {
	beego.Controller
}

type AdminLeftHandel struct {
	beego.Controller
}

type AdminNewsHandel struct {
	beego.Controller
}

func (this *AdminLogoutHandel) Get() {
	this.Ctx.SetCookie("username", "", -1, "/")
	this.Ctx.WriteString("<script>top.location.href='/webadmin';</script>")
	return
}

func (this *AdminLoginHandel) Get() {
	this.TplNames = "_login.html"
}

func (this *AdminMainHandel) Get() {
	if !checkLoginStatus(this.Ctx) {
		this.Ctx.WriteString("<script>alert('请登录后再操作..');top.location.href='/webadmin';</script>")
		return
	}
	this.TplNames = "_adminmain.html"
}
func (this *AdminLeftHandel) Get() {
	this.TplNames = "_adminleft.html"
}

//文章管理Get请求的处理
func (this *AdminNewsHandel) Get() {
	object := this.Ctx.Input.Params[":object"]
	action := this.Ctx.Input.Params[":action"]

	if !checkLoginStatus(this.Ctx) {
		this.Ctx.WriteString("<script>alert('请登录后再操作..');top.location.href='/webadmin';</script>")
		return
	}

	if object == "news" {
		switch action {
		case "add":
			this.TplNames = "_admin_addnews.html"
			classlist, err := models.GetClassList("id")
			common.Check(err)
			this.Data["classlist"] = classlist
		case "list":
			pagestr := this.Ctx.Input.Params[":page"]
			page, _ := strconv.ParseInt(pagestr, 10, 64)
			if page < 1 {
				page = 1
			}
			var pagesize int64
			pagesize = 20
			offset := (page - 1) * pagesize

			this.TplNames = "_admin_newslist.html"
			newslist, count := models.GetNewsPageList(pagesize, offset, 0)
			this.Data["newslist"] = newslist
			pager := common.PageList(pagesize, page, count, false, "/webadmin/news/list")
			this.Data["pager"] = pager
		case "edit":
			this.TplNames = "_admin_addnews.html"
			classlist, err := models.GetClassList("id")
			common.Check(err)

			idstr := this.Ctx.Input.Params[":id"]
			tid, _ := strconv.ParseInt(idstr, 10, 64)
			newsinfo, _ := models.GetNewsInfo(tid)
			this.Data["newsinfo"] = newsinfo
			this.Data["classlist"] = classlist
		case "delete":
			idstr := this.Ctx.Input.Params[":id"]
			tid, _ := strconv.ParseInt(idstr, 10, 64)
			err := models.DeleteNewsInfo(tid)
			common.Check(err)
			this.Redirect("/webadmin/news/list", 302)
		}
	}
}

//文章管理Post请求的处理
func (this *AdminNewsHandel) Post() {
	if !checkLoginStatus(this.Ctx) {
		this.Ctx.WriteString("<script>alert('请登录后再操作..');top.location.href='/webadmin';</script>")
		return
	}
	object := this.Ctx.Input.Params[":object"]
	action := this.Ctx.Input.Params[":action"]
	if object == "news" {
		switch action {
		case "add":
			title := this.Input().Get("title")
			content := this.Input().Get("content")
			anthor := this.Input().Get("anthor")
			cidstr := this.Input().Get("cid")
			cid, _ := strconv.ParseInt(cidstr, 10, 64)
			source := this.Input().Get("source")
			var err error
			err = models.AddNewsInfo(0, cid, title, content, anthor, source)
			if err != nil {
				common.Check(err)
				this.Ctx.WriteString("<script>alert('添加失败，请重试...');self.location.href='/webadmin/news/add';</script>")
			} else {
				this.Redirect("/webadmin/news/add", 302)
			}
			return
		case "edit":
			idstr := this.Ctx.Input.Params[":id"]
			id, _ := strconv.ParseInt(idstr, 10, 64)
			title := this.Input().Get("title")
			content := this.Input().Get("content")
			anthor := this.Input().Get("anthor")
			cidstr := this.Input().Get("cid")
			cid, _ := strconv.ParseInt(cidstr, 10, 64)
			source := this.Input().Get("source")
			var err error
			err = models.AddNewsInfo(id, cid, title, content, anthor, source)
			if err != nil {
				common.Check(err)
				this.Ctx.WriteString("<script>alert('修改失败，请重试...');self.location.href='/webadmin/news/edit/'+idstr;</script>")
			} else {
				this.Redirect("/webadmin/news/edit/"+idstr, 302)
			}
			return
		}
	}
}

//提交登录
func (this *AdminLoginHandel) Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")

	if username == beego.AppConfig.String("username") && password == beego.AppConfig.String("password") {
		this.Ctx.SetCookie("username", username, "/")
		this.Redirect("/webadmin/main", 302)
	} else {
		this.Redirect("/webadmin", 302)
	}

	return
}

//用户状态验证
func checkLoginStatus(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}

	username := ck.Value

	return username == beego.AppConfig.String("username")
}
