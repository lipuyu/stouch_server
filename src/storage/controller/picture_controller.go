package controller

import (
	"github.com/gin-gonic/gin"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"stouch_server/src/common/er"
	"stouch_server/src/common/re"
	"stouch_server/src/common/utils"
	"stouch_server/src/core"
	"stouch_server/src/storage/model"
	"stouch_server/src/storage/service"
	"strconv"
	"strings"
)

func GetBy(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, "")
		return
	}
	picture := model.Picture{Id: id}
	if ok, _ := core.Orm.Get(&picture); ok {
		c.JSON(http.StatusOK, re.Data(gin.H{"picture": picture}))
	} else {
		c.JSON(http.StatusOK, er.SourceNotExistError)
	}
}

func Post(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
	}
	img, _, err := image.DecodeConfig(file)
	width, height, size := img.Width, img.Height, fileHeader.Size
	file.Seek(0, 0)
	md5 := utils.GetMD5(file)
	file.Seek(0, 0)
	sr := strings.Split(fileHeader.Filename, ".")
	picture := &model.Picture{Width: width, Height: height, Size: size, Md5: md5, Format: sr[len(sr)-1]}
	if service.GetOrSave(md5+"."+string(sr[len(sr)-1]), file) {
		if _, err := core.Orm.Get(picture); err != nil {
		}
	} else {
		if _, err := core.Orm.Insert(picture); err != nil {
		}
	}
	c.JSON(http.StatusOK, re.Data(gin.H{"picture": *picture}))
}

func PostEditor(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
	}
	img, _, err := image.DecodeConfig(file)
	width, height, size := img.Width, img.Height, fileHeader.Size
	file.Seek(0, 0)
	md5 := utils.GetMD5(file)
	file.Seek(0, 0)
	sr := strings.Split(fileHeader.Filename, ".")
	picture := &model.Picture{Width: width, Height: height, Size: size, Md5: md5, Format: sr[len(sr)-1]}
	if service.GetOrSave(md5+"."+sr[len(sr)-1], file) {
		if _, err := core.Orm.Get(picture); err != nil {
		}
	} else {
		if _, err := core.Orm.Insert(picture); err != nil {
		}
	}
	c.JSON(http.StatusOK, gin.H{"default": "https://stouch.oss-cn-beijing.aliyuncs.com/" + picture.Md5 + "." + picture.Format})
}
