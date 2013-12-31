package controllers

import (
	"github.com/astaxie/beego"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"zituoblog/common"
)

var (
	server  = "/"
	basedir = "upload/"
)

var maximage uint

type UploadMessage struct {
	Errcode int
	Message string
	Url     string
}

type UpLoadHandel struct {
	beego.Controller
}

//图片上传
func (this *UpLoadHandel) Post() {
	file, handel, err := this.GetFile("imgFile")
	if err != nil {
		common.Check(err)
		return
	}
	defer file.Close()

	now := time.Now()

	dir := basedir + now.Format("2006-01") + "/" + now.Format("02")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		common.Check(err)
		return
	}

	filename := handel.Filename
	split_part := strings.Split(filename, ".")
	ext := "." + strings.ToLower(split_part[1])

	if !common.IsImageFile(ext) {
		ms := UploadMessage{1, "can't upload non-image file", ""}
		this.Data["json"] = &ms
		this.ServeJson()
		return
	}

	filename = common.GetGuid() + ext
	isimg := true

	if ext == ".gif" {

		data, err := ioutil.ReadAll(file)
		if err != nil {
			common.Check(err)
			return
		}
		err = ioutil.WriteFile(dir+"/"+filename, data, 0777)
		if err != nil {
			common.Check(err)
			return
		}
	} else {

		tempfilename := "temp_" + filename
		data, err := ioutil.ReadAll(file)
		if err != nil {
			common.Check(err)
			return
		}
		err = ioutil.WriteFile(dir+"/"+tempfilename, data, 0777)
		if err != nil {
			common.Check(err)
			return
		}

		isimg = WaterMark(dir+"/"+tempfilename, dir+"/"+filename, ext)

		err = os.Remove(dir + "/" + tempfilename)
		if err != nil {
			common.Check(err)
		}
	}

	if isimg {
		imgurl := server + dir +"/"+ filename
		ms := UploadMessage{0, "", imgurl}
		this.Data["json"] = &ms
		this.ServeJson()
	} else {
		ms := UploadMessage{1, "can't upload non-image file", ""}
		this.Data["json"] = &ms
		this.ServeJson()
	}

}


//图片略缩
func Resize(img image.Image) (m image.Image) {
	maximage = 850
	if img.Bounds().Dx() > 850 {
		return resize.Resize(maximage, 0, img, resize.MitchellNetravali)
	} else {
		return img
	}	
}

//图片水印
func WaterMark(tempfilepath string, newfilepath string, ext string) bool {

	imgb, _ := os.Open(tempfilepath)
	img, _, err := image.Decode(imgb)
	if err != nil {
		common.Check(err)
		imgb.Close()
		return false
	}
	//略缩后再加水印
	img = Resize(img)
	defer imgb.Close()

	wmb, err := os.Open("static/watermark/text.png")
	watermark, _ := png.Decode(wmb)
	if err != nil {
		common.Check(err)
		wmb.Close()
		return false
	}
	defer wmb.Close()

	offset := image.Pt(img.Bounds().Dx()-watermark.Bounds().Dx()-10, img.Bounds().Dy()-watermark.Bounds().Dy()-10)
	b := img.Bounds()
	m := image.NewNRGBA(b)

	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	imgw, _ := os.Create(newfilepath)
	jpeg.Encode(imgw, m, &jpeg.Options{100})
	defer imgw.Close()
	return true

}
