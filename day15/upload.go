package main

/**
 * gin 文件上传
 */

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ok(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
	})
}

func fail500(context *gin.Context, err error) {
	context.JSON(http.StatusInternalServerError, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  err,
	})
}

func uploadHandler(context *gin.Context) {
	file, err := context.FormFile("file")
	if err != nil {
		fail500(context, err)
		return
	}
	err = context.SaveUploadedFile(file,
		fmt.Sprintf("./day15/static/images/%s", file.Filename))
	if err != nil {
		fail500(context, err)
		return
	}
	ok(context)
}

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("./day15/template/**")
	engine.GET("/upload", func(context *gin.Context) {
		context.HTML(http.StatusOK, "upload.html", nil)
	})
	engine.POST("/upload", uploadHandler)
	engine.Run()
}
