package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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

