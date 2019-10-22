package routers

import (
	"gin_blogweb/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
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

			//显示文章内容
			v1.GET("/show/:id", controllers.ShowArticleGet)

			//更新文章
			v1.GET("/update", controllers.UpdateArticleGet)
			v1.POST("/update", controllers.UpdateArticlePost)

			// 删除文章
			v1.GET("/delete", controllers.DeleteArticleGet)
		}

		// 显示文章详情
		router.GET("/show/:id", controllers.ShowArticleGet)

		// 标签功能
		router.GET("/tags", controllers.TagsGet)
	}
	return router
}
