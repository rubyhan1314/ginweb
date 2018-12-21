package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
	"fmt"
)


func ExitGet(c *gin.Context)  {


	//清除该用户登录状态的数据
	session := sessions.Default(c)
	session.Delete("loginuser")
	session.Save()
	//session.Clear()

	fmt.Println("delete session...",session.Get("loginuser"))
	c.Redirect(http.StatusMovedPermanently,"/")
}
