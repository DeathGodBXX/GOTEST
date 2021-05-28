package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func mains() {
	r := gin.Default()
	//1.
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	//2.获取API中的信息
	r.GET("/user/:name/*action", func(c *gin.Context) {
		//name对应/后面的value，action对应/value
		name := c.Param("name")
		action := c.Param("action")
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" "+action)

	})

	//3.获取query指定位置的信息，如果没有返回默认值  /person?name=xiaoming
	r.GET("/person", func(c *gin.Context) {
		name := c.DefaultQuery("name", "小白")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})

	//4.先解析模板文件，再使用HTML渲染
	// r.LoadHTMLGlob("index.html")
	r.LoadHTMLGlob("../templates/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "")
	})

	//5.获取表单数据和单个文件
	r.MaxMultipartMemory = 1024
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		file, err := c.FormFile("upload")
		if err != nil {
			c.String(500, "上传文件出错")
		}
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,types:%s", username, password, types))
	})

	//6.上传多个文件
	// r.LoadHTMLFiles("home.html")
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", "")
	})

	r.MaxMultipartMemory = 1024 * 1024
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintln(err.Error()))
		}
		files := form.File["upload"]
		for _, file := range files {
			if err = c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintln(err.Error()))
			}
		}
		c.String(http.StatusOK, fmt.Sprintln(len(files)))
	})
	r.Run(":8080")
}
