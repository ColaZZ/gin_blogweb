package controllers

import (
	"gin_blogweb/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateArticleGet(c *gin.Context){
	isLogin := GetSession(c)
	idStr := c.Query("id")
	id, _ := strconv.Atoi(idStr)

	art := models.QueryArticleWithId(id)
	c.HTML(http.StatusOK, "write_article.html", gin.H{
		"IsLogin": isLogin,
		"Title": art.Title,
		"Tags": art.Tags,
		"Short": art.Short,
		"Content": art.Content,
		"Id": art.Id,
	})
}

func UpdateArticlePost(c *gin.Context){
	idStr := c.Query("id")
	id, _ := strconv.Atoi(idStr)
	title := c.PostForm("title")
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")
	
	art := models.Article{
		Id:         id,
		Title:      title,
		Tags:       tags,
		Short:      short,
		Content:    content,
		Author:     "",
		CreateTime: 0,
	}

	_, err := models.UpdateArticle(art)
}
