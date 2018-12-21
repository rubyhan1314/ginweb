package controllers

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"strconv"
	"blogweb_gin/models"
	"net/http"
)



//点击删除后重定向到首页
func  DeleteArticleGet(c *gin.Context) {

	idstr := c.Query("id")
	id, _ := strconv.Atoi(idstr)
	fmt.Println("删除 id:", id)



	_, err := models.DeleteArticle(id)
	if err != nil {
		log.Println(err)
	}
	//c.JSON(http.StatusOK, gin.H{"IsLogin": islogin})
	c.Redirect(http.StatusMovedPermanently,"/")
}
