package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSession(c *gin.Context) bool {
	session := sessions.Default(c)
	loginuser := session.Get("loginuser")
	if loginuser != nil {
		return true
	} else {
		return false
	}
}

func HomeGet(c *gin.Context){
	isLogin := GetSession(c)
	c.HTML(http.StatusOK, "home.html", gin.H{"IsLogin": isLogin})
}
