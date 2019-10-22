package controllers

import (
	"gin_blogweb/models"
	"gin_blogweb/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})
}

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	id := models.QueryUserWithParams(username, utils.MD5(password))

	if id > 0 {
		session := sessions.Default(c)
		session.Set("loginuser", username)
		_ = session.Save()
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code":1, "message": "登陆失败"})
	}
}
