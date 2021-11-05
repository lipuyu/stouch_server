package controller

import (
	"github.com/kataras/iris"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"stouch_server/src/common/base"
	"stouch_server/src/common/er"
	"stouch_server/src/common/utils"
	"stouch_server/src/core"
	model2 "stouch_server/src/storage/model"
	service2 "stouch_server/src/storage/service"
	"strings"
)

type PictureController struct {
	Ctx iris.Context
}

func (c *PictureController) GetBy(id int64) interface{} {
	picture := model2.Picture{Id: id}
	if ok, _ := core.Orm.Get(&picture); ok {
		return re.NewByData(map[string]model2.Picture{"picture": picture})
	} else {
		return er.SourceNotExistError
	}
}

func (c *PictureController) Post() interface{} {
	file, fileHeader, err := c.Ctx.FormFile("file")
	if err != nil {
	}
	img, _, err := image.DecodeConfig(file)
	width, height, size := img.Width, img.Height, fileHeader.Size
	file.Seek(0, 0)
	md5 := utils.GetMD5(file)
	file.Seek(0, 0)
	sr := strings.Split(fileHeader.Filename, ".")
	picture := &model2.Picture{Width: width, Height: height, Size: size, Md5: md5, Format: sr[len(sr)-1]}
	if service2.GetOrSave(md5+"."+string(sr[len(sr)-1]), file) {
		if _, err := core.Orm.Get(picture); err != nil {
		}
	} else {
		if _, err := core.Orm.Insert(picture); err != nil {
		}
	}
	return re.NewByData(iris.Map{"picture": *picture})
}

func (c *PictureController) PostEditor() interface{} {
	file, fileHeader, err := c.Ctx.FormFile("file")
	if err != nil {
	}
	img, _, err := image.DecodeConfig(file)
	width, height, size := img.Width, img.Height, fileHeader.Size
	file.Seek(0, 0)
	md5 := utils.GetMD5(file)
	file.Seek(0, 0)
	sr := strings.Split(fileHeader.Filename, ".")
	picture := &model2.Picture{Width: width, Height: height, Size: size, Md5: md5, Format: sr[len(sr)-1]}
	if service2.GetOrSave(md5+"."+string(sr[len(sr)-1]), file) {
		if _, err := core.Orm.Get(picture); err != nil {
		}
	} else {
		if _, err := core.Orm.Insert(picture); err != nil {
		}
	}
	return iris.Map{"default": "https://lipuyu.oss-cn-shanghai.aliyuncs.com/" + picture.Md5 + "." + picture.Format}
}
