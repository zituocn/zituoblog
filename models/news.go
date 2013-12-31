package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//添加文章
func AddNewsInfo(id int64, cid int64, title, content, anthor, source string) error {
	o := orm.NewOrm()
	newsinfo := &NewsInfo{
		Id:       id,
		Cid:      cid,
		Title:    title,
		Content:  content,
		Poster:   "阿布",
		Posterid: 1,
		Views:    0,
		Anthor:   anthor,
		Source:   source,
		Addtime:  time.Now(),
		Status:   0,
	}
	var err error
	if newsinfo.Id == 0 {
		_, err = o.Insert(newsinfo)
	} else {
		_, err = o.Update(newsinfo)
	}

	//更新分类表上，最后发帖时间
	_, _ = o.Raw("update class_info set Lastposttime=? where id= ? ", time.Now(), cid).Exec()
	return err
}

//获取某个文章详情
func GetNewsInfo(tid int64) (info *NewsInfo, err error) {
	o := orm.NewOrm()
	info = new(NewsInfo)
	qs := o.QueryTable("news_info")

	err = qs.Filter("id", tid).One(info)

	//回写点击量
	_, _ = o.Raw("update news_info set views=views+1 where id= ? ", tid).Exec()
	return info, err
}

//删除文章
func DeleteNewsInfo(tid int64) error {
	o := orm.NewOrm()
	info := &NewsInfo{Id: tid}
	_, err := o.Delete(info)
	return err
}

///后面列表页的分页列表
///cid有可能=0
func GetNewsPageList(pagesize, page, cid int64) ([]*NewsInfo, int64) {
	o := orm.NewOrm()
	newslist := make([]*NewsInfo, 0)
	var ct int64
	if cid == 0 {
		o.QueryTable("news_info").OrderBy("-id").Limit(pagesize, page).All(&newslist)
		ct, _ = o.QueryTable("news_info").Count()
	} else {

		o.QueryTable("news_info").OrderBy("-id").Limit(pagesize, page).Filter("cid", cid).All(&newslist)
		ct, _ = o.QueryTable("news_info").Filter("cid", cid).Count()
	}
	return newslist, ct

}

//按分类取出文章，并且分页列表
func GetNewsListByCid(pagesize, page, cid int64) ([]*NewsInfo, int64) {
	o := orm.NewOrm()
	newslist := make([]*NewsInfo, 0)
	o.QueryTable("news_info").OrderBy("-id").Limit(pagesize, page).Filter("cid", cid).All(&newslist, "Id", "Title", "Addtime")
	ct, _ := o.QueryTable("news_info").Filter("cid", cid).Count()
	return newslist, ct
}

//翻页方式取出所有文章
func GetNewsList() (newslist []*NewsInfo, err error) {
	o := orm.NewOrm()

	newslist = make([]*NewsInfo, 0)

	qs := o.QueryTable("news_info")

	_, err = qs.OrderBy("-id").All(&newslist)

	return newslist, err
}
