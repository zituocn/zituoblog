package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strings"
	"time"
	"zituoblog/common"
)

//后台用户结构
type UserInfo struct {
	Id       int
	Username string
	Password string
}

//文章分类
type ClassInfo struct {
	Id           int64
	Cname        string
	Ename        string
	Topics       int       `orm:"index"`
	Lastposttime time.Time `rom:"index"`
}

const (
	_DB_NAME        = "data/zituoblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

//文章
type NewsInfo struct {
	Id       int64
	Cid      int64
	Title    string
	Content  string `orm:"size(8000)"`
	Poster   string
	Posterid int
	Views    int `orm:"index"`
	Anthor   string
	Source   string
	Addtime  time.Time
	Status   int
}

func RegisterDB() {
	if !common.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(NewsInfo), new(ClassInfo), new(UserInfo))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
	orm.RunSyncdb("default", false, false)
}

//返回导航链接
func GetNavString() (html string) {
	classlist, err := GetClassList("-lastposttime")
	common.Check(err)
	list := []string{}
	for _, classinfo := range classlist {
		list = append(list, fmt.Sprintf("<li><a href=\"/category/%s\">%s</a></li>", classinfo.Ename, classinfo.Cname))
	}
	html = strings.Join(list, "\n")
	return html
}
