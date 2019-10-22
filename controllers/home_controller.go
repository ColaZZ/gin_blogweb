package controllers

import (
	"gin_blogweb/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HomeGet(c *gin.Context){
	isLogin := GetSession(c)

	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0{
		page = 1
	}
	HomeFooterPageCode := models.ConfigHomeFooterPageCoder(page)
	articleList, _ := models.FindArticleWithPage(page)
	content := models.MakeHomeBlocks(articleList, isLogin)

	c.HTML(http.StatusOK, "home.html", gin.H{
		"IsLogin": isLogin,
		"Content": content,
		"HasFooter":true,
		"PageCode": HomeFooterPageCode,
	})
}
