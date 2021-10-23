package main

/**
 * gin框架
 */

import (
	"fmt"
	"gitee.com/yousj/golang_study/day12"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	db, err := dbtool.ConnMysql()
	if err != nil {
		fmt.Println(err)
		return
	}

	//books, err := dbtool.Query2Json(db)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	books := dbtool.Query(db)

	engine := gin.Default()

	// 响应json格式数据
	//engine.GET("/index", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusOK,
	//		"msg":  "hello gin!",
	//	})
	//})

	//engine.GET("/book", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, books)
	//})
	// 模板
	engine.LoadHTMLGlob("./day14/template/**")
	engine.GET("/book", func(context *gin.Context) {
		context.HTML(http.StatusOK, "book.tmpl", books)
	})
	engine.Run()
}
