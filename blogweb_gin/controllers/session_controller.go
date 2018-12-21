package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

//type BaseController struct {
//	beego.Controller
//	IsLogin   bool
//	Loginuser interface{}
//}




func SessionMiddleware() gin.HandlerFunc {

	// Do some initialization logic here

	// Foo()

	return func(c *gin.Context) {

		session := sessions.Default(c)
		loginuser := session.Get("loginuser")

		fmt.Println("loginuser:", loginuser)
		if loginuser != nil {
			//this.IsLogin = true
			//this.Loginuser = loginuser

			c.Set("IsLogin", true)
			c.Set("Loginuser", loginuser)

			//c.JSON(http.StatusOK, gin.H{"IsLogin": true, "Loginuser": loginuser})
		} else {
			c.Set("IsLogin", false)
			//c.JSON(http.StatusOK, gin.H{"IsLogin": false})
			//this.IsLogin = false
		}
		islogin, _ := c.Get("IsLogin")
		loginuser2, _ := c.Get("Loginuser")
		fmt.Println("middleware...", islogin, loginuser2)

		c.Next()


	}

}
//获取session
func GetSession(c *gin.Context) bool {

	session := sessions.Default(c)
	loginuser := session.Get("loginuser")

	fmt.Println("loginuser:", loginuser)

	if loginuser != nil {
		//this.IsLogin = true
		//this.Loginuser = loginuser

		return true
		//c.JSON(http.StatusOK, gin.H{"IsLogin": true, "Loginuser": loginuser})
	} else {
		return false
		//c.JSON(http.StatusOK, gin.H{"IsLogin": false})
		//this.IsLogin = false
	}
	//this.Data["IsLogin"] = this.IsLogin
}
