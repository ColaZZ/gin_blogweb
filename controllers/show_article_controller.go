package controllers

import (
	"gin_blogweb/models"
	"gin_blogweb/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ShowArticleGet(c *gin.Context) {
	isLogin := GetSession(c)
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	art := models.QueryArticleWithId(id)

	c.HTML(http.StatusOK, "show_article.html", gin.H{
		"isLogin": isLogin,
		"Title": art.Title,
		"Content": utils.SwitchMarkdownToHtml(art.Content),
	})
}
