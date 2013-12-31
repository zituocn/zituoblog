package common

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	//"github.com/astaxie/beego"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func IsImageFile(ext string) bool {
	exts := ".jpg|.jpeg|.gif|.png|.bmp"
	return strings.Contains(exts, ext)
}

func GetGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}

//错误处理
func Check(err error) {
	if err != nil {
		panic(err)
		//beego.Error(err.Error())
		os.Exit(1)
	}
}

//Md5编码
func Md5(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	return fmt.Sprintf("%x", m.Sum(nil))
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//日期时间格式化
func DateT(t time.Time, format string) string {
	res := strings.Replace(format, "MM", t.Format("01"), -1)
	res = strings.Replace(res, "M", t.Format("1"), -1)
	res = strings.Replace(res, "DD", t.Format("02"), -1)
	res = strings.Replace(res, "D", t.Format("2"), -1)
	res = strings.Replace(res, "YYYY", t.Format("2006"), -1)
	res = strings.Replace(res, "YY", t.Format("06"), -1)
	res = strings.Replace(res, "HH", fmt.Sprintf("%02d", t.Hour()), -1)
	res = strings.Replace(res, "H", fmt.Sprintf("%d", t.Hour()), -1)
	res = strings.Replace(res, "hh", t.Format("03"), -1)
	res = strings.Replace(res, "h", t.Format("3"), -1)
	res = strings.Replace(res, "mm", t.Format("04"), -1)
	res = strings.Replace(res, "m", t.Format("4"), -1)
	res = strings.Replace(res, "ss", t.Format("05"), -1)
	res = strings.Replace(res, "s", t.Format("5"), -1)
	return res
}

//字串截取
func SubString(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

//显示分页链接
func PageList(pagesize, page, recordcount int64, first bool, path string) (pager string) {
	if recordcount == 0 {
		return ""
	}

	var pagecount int64
	pagecount = 0

	if recordcount%pagesize == 0 {
		pagecount = recordcount / pagesize
	} else {
		pagecount = (recordcount / pagesize) + 1
	}

	pager = "<span class=\"page-numbers\">" + strconv.FormatInt(page, 10) + "/" + strconv.FormatInt(pagecount, 10) + "</span>"

	if pagecount < 2 {
		return "<span class=\"page-numbers\">共有" + strconv.FormatInt(recordcount, 10) + "条数据</span>"
	}

	if page > 1 {
		if page == 2 {
			pager = pager + "<a href=\"" + path + "\" class=\"page-numbers\">上一页</a>"
		} else {
			pager = pager + "<a href=\"" + path + "/" + strconv.FormatInt(page-1, 10) + "\" class=\"page-numbers\">上一页</a>"
		}
	} else {
		pager = pager + "<a href=\"" + path + "\" class=\"page-numbers\">上一页</a>"
	}

	if page < pagecount {
		pager = pager + "<a href=\"" + path + "/" + strconv.FormatInt(page+1, 10) + "\" class=\"page-numbers\">下一页</a>"
	} else {
		pager = pager + "<a href=\"" + path + "/" + strconv.FormatInt(pagecount, 10) + "\" class=\"page-numbers\">下一页</a>"
	}
	pager = pager
	return pager

}
