package controllers

import (
	"fmt"
	"gin_blogweb/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func UploadPost(c *gin.Context){
	fileHeader, err := c.FormFile("upload")
	if err != nil {
		responseErr(c, err)
		return
	}

	now := time.Now()
	fileType := "other"
	// 文件后缀判断img
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg"{
		fileType = "img"
	}
	// 拼接文件夹路径，创建文件夹
	fileDir := fmt.Sprintf("/static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		responseErr(c, err)
		return
	}

	// 文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePathStr := filepath.Join(fileDir, fileName)

	// 将上传的文件拷贝到本地
	_ = c.SaveUploadedFile(fileHeader, filePathStr)

	// 插入数据库
	if fileType == "img"{
		album := models.Album{
			Id:         0,
			FilePath:   filePathStr,
			Filename:   fileName,
			CreateTime: 0,
		}
		_, _ = models.InsertAlbum(album)
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "上传成功"})
}

func responseErr(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": err})
}
