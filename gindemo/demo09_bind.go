package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-gonic/gin/binding"
)

type Login struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	//1.binding JSON
	// Example for binding JSON ({"user": "hanru", "password": "hanru123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		//其实就是将request中的Body中的数据按照JSON格式解析到json变量中
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.User != "hanru" || json.Password != "hanru123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})


	// 2. Example for binding XML (
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>user</user>
	//		<password>123</user>
	//	</root>)
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "hanru" || xml.Password != "hanru123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// 3. Form 绑定普通表单的例子
	// Example for binding a HTML form (user=hanru&password=hanru123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		//方法一：对于FORM数据直接使用Bind函数, 默认使用使用form格式解析,if c.Bind(&form) == nil
		// 根据请求头中 content-type 自动推断.
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "hanru" || form.Password != "hanru123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// 4.
	router.POST("/login", func(c *gin.Context) {
		var form Login
		//方法二: 使用BindWith函数,如果你明确知道数据的类型
		// 你可以显式声明来绑定多媒体表单：
		// c.BindWith(&form, binding.Form)
		// 或者使用自动推断:
		if c.BindWith(&form, binding.Form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in ..... "})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})



	// 5.URI
	router.GET("/:user/:password", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"username": login.User, "password": login.Password})
	})



	router.Run(":8080")
}
