package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AlbumGet(c *gin.Context){
	isLogin := GetSession(c)

	c.HTML(http.StatusOK, "album.html", gin.H{"IsLogin": isLogin})
}