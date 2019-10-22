package controllers

import (
	"gin_blogweb/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TagsGet(c *gin.Context){
	isLogin := GetSession(c)

	tags := models.QueryArticleWithParam("tags")

	c.HTML(http.StatusOK, "tags.html", gin.H{
		"IsLogin": isLogin,
		"Tags": models.HandleTagsListData(tags),
	})
}
