package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"image"
	"stouch_server/common/er"
	"stouch_server/common/utils"
	"stouch_server/conf"
	"stouch_server/storage/model"
	"stouch_server/storage/service"
	"strings"
)

type PictureController struct{
	Ctx iris.Context
}

func (c *PictureController) GetBy(id int64) interface{}{
	picture := model.Picture{Id: id}
	if ok, _ := conf.Orm.Get(&picture); ok {
		return er.NoError.SetData(map[string]model.Picture{"picture": picture})
	} else  {
		return er.SourceNotExistError
	}
}

func (c *PictureController) Post() interface{}{
	file, fileHeader, err := c.Ctx.FormFile("file")
	file.Seek(0, 0)
	img, _, err := image.Decode(file)
	width, height, size := img.Bounds().Max.X, img.Bounds().Max.Y, fileHeader.Size
	file.Seek(0, 0)
	md5 := utils.GetMD5(file)
	file.Seek(0, 0)
	sr := strings.Split(fileHeader.Filename, ".")
	picture := &model.Picture{Width: width, Height: height, Size:size, Md5:md5, Format: sr[len(sr) - 1]}
	if  _, err := conf.Orm.Insert(picture); err == nil { } else {
		conf.Logger.Error(err)
	}
	if err != nil {
		fmt.Println(fileHeader.Filename, err)
	}
	service.Save(fileHeader.Filename, file)
	return fileHeader.Filename
}
