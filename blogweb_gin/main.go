package main

import (
	"blogweb_gin/routers"
	"blogweb_gin/database"
)

func main() {
	database.InitMysql()
	router := routers.InitRouter()

	//静态资源
	router.Static("/static", "./static")

	router.Run(":8081")
}

