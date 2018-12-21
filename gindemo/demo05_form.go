package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func main() {

	router := gin.Default()
	//form
	router.POST("/form", func(c *gin.Context) {
		type1 := c.DefaultPostForm("type", "alert") //可设置默认值
		username := c.PostForm("username")
		password := c.PostForm("password")

		//hobbys := c.PostFormMap("hobby")
		//hobbys := c.QueryArray("hobby")
		hobbys := c.PostFormArray("hobby")

		c.String(http.StatusOK, fmt.Sprintf("type is %s, username is %s, password is %s,hobby is %v", type1, username, password,hobbys))

	})

	router.Run(":9527")
}
