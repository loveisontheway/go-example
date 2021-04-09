package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// *** 上传单个文件 ***
/*
	multipart/form-data格式用于文件上传
	gin文件上传与原生的net/http方法类似，不同在于gin把原生的request封装到c.Request中
 */
func main() {
	// 上传特定文件
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file: %v", err)
		}
		//headers.Size 获取文件大小
		if headers.Size > 1024*1024*2 {
			fmt.Println("文件太大了")
			return
		}
		//headers.Header.Get("Content-Type")获取上传文件的类型
		if headers.Header.Get("Content-Type") != "image/png" {
			fmt.Println("只允许上传png图片")
			return
		}
		c.SaveUploadedFile(headers, "./video/"+headers.Filename)
		c.String(http.StatusOK, headers.Filename)
	})
	r.Run()
}
