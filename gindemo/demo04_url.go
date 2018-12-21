package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func main() {

	router := gin.Default()
	router.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Guest") //可设置默认值

		//nickname := c.Query("nickname") // 是 c.Request.URL.Query().Get("nickname") 的简写

		c.String(http.StatusOK, fmt.Sprintf("Hello %s ", name))
	})

	router.Run(":9527")
}

