package main

import "gin_blogweb/routers"

func main(){
	router := routers.InitRouter()

	router.Static("/static", "./static")
	_ = router.Run(":8081")
}
