package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// 下面测试静态文件服务
	// 显示当前文件夹下的所有文件/或者指定文件
	//定义多文件的路径,使用的是系统的路径(绝对,相对地址都可以)
	router.StaticFS("/showDir", http.Dir("."))
	//router.StaticFS("/files", http.Dir("/bin"))
	//Static提供给定文件系统根目录中的文件。
	router.Static("/files", "/bin")



	router.StaticFile("/image", "./assets/miao.jpg")

	router.Run(":8080")
}
