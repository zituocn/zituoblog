package models

import (
	"github.com/astaxie/beego/orm"
)

//获取所有文章分类
func GetClassList(order string) ([]*ClassInfo, error) {
	if len(order) == 0 {
		order = "id"
	}
	o := orm.NewOrm()
	classlist := make([]*ClassInfo, 0)
	qs := o.QueryTable("class_info")
	_, err := qs.OrderBy(order).All(&classlist)
	return classlist, err
}

//按ename取某分类信息
func GetClassInfo(ename string) (*ClassInfo, error) {
	o := orm.NewOrm()
	info := new(ClassInfo)
	err := o.QueryTable("class_info").Filter("ename", ename).One(info)
	return info, err
}
