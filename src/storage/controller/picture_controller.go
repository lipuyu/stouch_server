package controller

import (
	"github.com/gin-gonic/gin"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"stouch_server/src/common/base"
	"stouch_server/src/common/er"
	"stouch_server/src/common/utils"
	"stouch_server/src/core"
	model2 "stouch_server/src/storage/model"
	service2 "stouch_server/src/storage/service"
	"strconv"
	"strings"
)

func getBy(c *gin.Context){
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, "")
		return
	}
	picture := model2.Picture{Id: id}
	if ok, _ := core.Orm.Get(&picture); ok {
		c.JSON(http.StatusOK, re.NewByData(gin.H{"picture": picture}))
	} else {
		c.JSON(http.StatusOK, er.SourceNotExistError)
	}
}

func post(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
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
	c.JSON(http.StatusOK, re.NewByData(gin.H{"picture": *picture}))
}

func postEditor(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
	}
	img, _, err := image.DecodeConfig(file)
	width, height, size := img.Width, img.Height, fileHeader.Size
	file.Seek(0, 0)
	md5 := utils.GetMD5(file)
	file.Seek(0, 0)
	sr := strings.Split(fileHeader.Filename, ".")
	picture := &model2.Picture{Width: width, Height: height, Size: size, Md5: md5, Format: sr[len(sr)-1]}
	if service2.GetOrSave(md5+"."+sr[len(sr)-1], file) {
		if _, err := core.Orm.Get(picture); err != nil {
		}
	} else {
		if _, err := core.Orm.Insert(picture); err != nil {
		}
	}
	c.JSON(http.StatusOK, gin.H{"default": "https://lipuyu.oss-cn-shanghai.aliyuncs.com/" + picture.Md5 + "." + picture.Format})
}
