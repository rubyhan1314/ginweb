package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"fmt"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, file.Filename)

		/*
        也可以直接使用io操作，拷贝文件数据。
        out, err := os.Create(filename)
        defer out.Close()
        _, err = io.Copy(out, file)
        */
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8080")
}
