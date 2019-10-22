package routers

import (
	"gin_blogweb/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	store := cookie.NewStore([]byte("loginuser"))
	router.Use(sessions.Sessions("mysession", store))
	{
		// 注册
		router.GET("/register", controllers.RegisterGet)
		router.POST("/register", controllers.RegisterPost)

		// 登陆
		router.GET("/login", controllers.LoginGet)
		router.POST("/login", controllers.LoginPost)

		//首页
		router.GET("/", controllers.HomeGet)

		//退出
		router.GET("/exit", controllers.ExitGet)

		v1 := router.Group("/article")
		{
			v1.GET("/add", controllers.AddArticleGet)
			v1.POST("/add", controllers.AddArticlePost)
		}
	}
	return router
}
