package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AboutMeGet(c *gin.Context){
	isLogin := GetSession(c)

	c.HTML(http.StatusOK, "about_me.html", gin.H{
		"IsLogin": isLogin,
		"wechat":"",
		"qq":"QQ：",
		"tel":"Tel：",
	})
}
