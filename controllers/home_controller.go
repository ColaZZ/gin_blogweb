package controllers

import (
	"gin_blogweb/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HomeGet(c *gin.Context){
	isLogin := GetSession(c)

	tag := c.Query("tag")
	page, _ := strconv.Atoi(c.Query("page"))

	var articleList []models.Article
	var hasFooter bool

	if len(tag) >0{
		articleList, _ = models.QueryArticlesWithTag(tag)
		hasFooter = false
	}else{
		articleList, _ = models.FindArticleWithPage(page)
		hasFooter = true
	}

	if page <= 0{
		page = 1
	}
	HomeFooterPageCode := models.ConfigHomeFooterPageCoder(page)
	content := models.MakeHomeBlocks(articleList, isLogin)

	c.HTML(http.StatusOK, "home.html", gin.H{
		"IsLogin": isLogin,
		"Content": content,
		"HasFooter":hasFooter,
		"PageCode": HomeFooterPageCode,
	})
}
