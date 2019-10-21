package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExitGet(c *gin.Context){
	session := sessions.Default(c)
	session.Delete("loginuser")
	_ = session.Save()

	c.Redirect(http.StatusMovedPermanently, "/")
}
