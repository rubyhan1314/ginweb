package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"time"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		//request in
		fmt.Println("before middleware...")
		//设置request变量到Context的Key中,通过Get等函数可以取得
		c.Set("request", "client_request")
		//发送request之前

		//发送requst之后
		c.Next()


		//response out
		// 这个c.Write是ResponseWriter,我们可以获得状态等信息

		status := c.Writer.Status()
		fmt.Println("after middleware,", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	router := gin.Default()

	router.Use(MiddleWare())
	{
		router.GET("/middleware", func(c *gin.Context) {
			//获取gin上下文中的变量
			request := c.MustGet("request").(string)
			req, _ := c.Get("request")
			fmt.Println("request:",request)
			c.JSON(http.StatusOK, gin.H{
				"middile_request": request,
				"request":         req,
			})
		})

		router.GET("/before", MiddleWare(), func(c *gin.Context) {
			request := c.MustGet("request").(string)
			c.JSON(http.StatusOK, gin.H{
				"middile_request": request,
			})
		})
	}





	router.Run(":8080")
}
