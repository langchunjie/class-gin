package main

import (
	"github.com/gin-gonic/gin"
	// "class-gin/controller"
	"net/http"
)

func main()  {
	router := gin.Default()

	router.LoadHTMLGlob("../template/*")
	router.GET("/class", class)

	// 路由组
/* 	user := router.Group("/user")
	{	// 请求参数在请求路径上
		user.GET("/get/:id/:username",controller.QueryById)
		user.GET("/query",controller.QueryParam)
		user.POST("/insert",controller.InsertNewUser)
		user.GET("/form",controller.RenderForm)
		user.POST("/form/post",controller.PostForm)
	} */

/* /* 	file := router.Group("/file")
	{
		// 跳转上传文件页面
		file.GET("/view",controller.RenderView)
		// 根据表单上传
		file.POST("/insert",controller.FormUpload)
		file.POST("/multiUpload",controller.MultiUpload)
		// base64上传
		file.POST("/upload",controller.Base64Upload)
	} */
 */
	// 指定地址和端口号
	router.Run(":9090")
}

func class(context *gin.Context) {
	println(">>>> class function start <<<<")

	context.JSON(http.StatusOK,gin.H{
		"code":200,
		"success":true,
	})
}

