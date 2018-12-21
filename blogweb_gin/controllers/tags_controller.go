package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"blogweb_gin/models"
	"net/http"
)



func TagsGet(c *gin.Context) {
	//获取session
	islogin := GetSession(c)

	tags := models.QueryArticleWithParam("tags")
	fmt.Println(models.HandleTagsListData(tags))

	//返回html
	c.HTML(http.StatusOK, "tags.html", gin.H{"Tags": models.HandleTagsListData(tags),"IsLogin":islogin})
}
