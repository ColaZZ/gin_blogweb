package main

import (
	"gin_blogweb/database"
	"gin_blogweb/routers"
)

func main(){
	database.InitMysql()
	router := routers.InitRouter()

	router.Static("/static", "./static")
	_ = router.Run(":8081")
}
