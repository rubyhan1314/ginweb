package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"blogweb_gin/models"
	"net/http"
)

func UpdateArticleGet(c *gin.Context) {

	//获取session
	islogin := GetSession(c)

	idstr := c.Query("id")
	id, _ := strconv.Atoi(idstr)
	fmt.Println(id)

	//获取id所对应的文章信息
	art := models.QueryArticleWithId(id)

	c.HTML(http.StatusOK, "write_article.html", gin.H{"IsLogin": islogin, "Title": art.Title, "Tags": art.Tags, "Short": art.Short, "Content": art.Content, "Id": art.Id})
}

//修改文章
func UpdateArticlePost(c *gin.Context) {

	idstr := c.Query("id")
	id, _ := strconv.Atoi(idstr)
	fmt.Println("postid:", id)

	//获取浏览器传输的数据，通过表单的name属性获取值
	title := c.PostForm("title")
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")

	//实例化model，修改数据库
	art := models.Article{id, title, tags, short, content, "", 0}
	_, err := models.UpdateArticle(art)

	//返回数据给浏览器
	response := gin.H{}
	if err == nil {
		//无误
		response = gin.H{"code": 1, "message": "更新成功"}
	} else {
		response = gin.H{"code": 0, "message": "更新失败"}
	}

	c.JSON(http.StatusOK, response)
}
