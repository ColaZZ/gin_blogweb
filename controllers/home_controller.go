package controllers

import (
	"gin_blogweb/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeGet(c *gin.Context){
	isLogin := GetSession(c)

	page := 1
	articleList, _ := models.FindArticleWithPage(page)
	models.MakeHomeBlocks(articleList, isLogin)

	c.HTML(http.StatusOK, "home.html", gin.H{"IsLogin": isLogin})
}
