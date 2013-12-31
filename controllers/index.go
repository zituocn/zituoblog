package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"net/http"
	"strings"
	"time"
	"zituoblog/common"
	"zituoblog/models"
)

type HomeHandel struct {
	beego.Controller
}

func (this *HomeHandel) Get() {
	this.TplNames = "_index.html"
	classlist, err := models.GetClassList("-lastposttime")
	common.Check(err)
	list := []string{}
	for _, classinfo := range classlist {
		item := "<div class=\"box\"><div class=\"list\">"
		item = item + fmt.Sprintf("<div class=\"title\">%s</div>", classinfo.Cname)
		item = item + "<ul>"
		newslist, _ := models.GetNewsListByCid(10, 0, classinfo.Id)
		for _, newsinfo := range newslist {
			item = item + fmt.Sprintf("<li><a href=\"/view/%d\" target=\"_blank\">%s</a></li>", newsinfo.Id, newsinfo.Title)
		}
		item = item + "</ul>"
		item = item + "</div>"
		item = item + "</div>"
		list = append(list, item)
	}
	this.Data["IndexListString"] = strings.Join(list, "\n")

}

//统一的404页面

func Page_not_found(w http.ResponseWriter, r *http.Request) {
	tpl_404 := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title>页面不存在</title>
<style type="text/css">
*{padding:0px;margin:0px;font-size:12px;font-family:Tahoma;}
body{ background:#fff;}
#error{background:#2290CF; margin:280px auto;width:520px;color:#fff;padding:10px 15px 10px 15px; overflow:hidden;}
#error a{color:#FFFF00;}
#error h2{font-size:16px;font-family:'Microsoft YaHei';}
#error ul{margin:10px 20px 10px 26px;}
#error li{line-height:22px;}
#error .info{font-style:italic;}
</style>
</head>
<body>
<div id="error">
<h2>对不起，您所访问的页面不存在...</h2>
<ul>
<li><a href="http://zituo.net">立即访问首页</a></li>
<li>3秒后自动返回首页...</li>
<li>管理员QQ：301109640 </li>
<li>电子邮箱：abu#zituo.net</li>
<li>出错时间：{{.datetime}} </li>
</ul>
<script style="text/javascript">
    setTimeout("jumppage()", 3000);
    function jumppage() {
        top.location.href = "http://zituo.net";
    }
</script> 
</div>
</body>
</html>`

	t := template.New("page_404")
	t, _ = t.Parse(tpl_404)
	w.WriteHeader(http.StatusNotFound)
	data := make(map[string]interface{})
	data["datetime"] = time.Now().Format("2006-01-02 15:04:05")
	t.Execute(w, data)
}
