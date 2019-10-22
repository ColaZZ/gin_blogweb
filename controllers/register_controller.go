package controllers

import (
	"fmt"
	"gin_blogweb/models"
	"gin_blogweb/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RegisterGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}

func RegisterPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	fmt.Println(username, password, repassword)

	id := models.QueryUserWithUsername(username)
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "该用户已经注册过"})
		return
	}

	password = utils.MD5(password)
	user := models.User{0, username, password, 0, time.Now().Unix()}
	_, err := models.InsertUser(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "注册失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "注册成功"})
	}

}

