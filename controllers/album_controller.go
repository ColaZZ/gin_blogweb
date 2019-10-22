package controllers

import (
	"gin_blogweb/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AlbumGet(c *gin.Context){
	isLogin := GetSession(c)
	albums,_ := models.FindAllAlbums()

	c.HTML(http.StatusOK, "album.html", gin.H{"IsLogin": isLogin, "Albums": albums})
}