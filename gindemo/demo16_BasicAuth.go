package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模拟私有数据
var secrets = gin.H{
	"hanru":    gin.H{"email": "hanru@163.com", "phone": "123433"},
	"wangergou": gin.H{"email": "wangergou@example.com", "phone": "666"},
	"ruby":   gin.H{"email": "ruby@guapa.com", "phone": "523443"},
}

func main() {
	r := gin.Default()

	// 使用 gin.BasicAuth 中间件，设置授权用户
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"hanru":    "hanru123",
		"wangergou": "1234",
		"ruby":   "hello2",
		"lucy":   "4321",
	}))

	// 定义路由
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取提交的用户名（AuthUserKey）
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	r.Run(":8080")
}
