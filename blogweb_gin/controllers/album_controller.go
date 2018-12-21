package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"blogweb_gin/models"
)



func  AlbumGet(c *gin.Context) {
	//获取session
	islogin := GetSession(c)
	albums,_ := models.FindAllAlbums()


	c.HTML(http.StatusOK, "album.html", gin.H{"IsLogin": islogin,"Album":albums})
}
