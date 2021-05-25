package example

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//UpLoadFile ...
func UpLoadFile() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/upload*") //路径开始不要带"/"
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "uploadfile.html", gin.H{
			"title":   "主题",
			"message": "新兴",
		})
	})

	//处理form表单请求
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("f1") //读取单一文件
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}

		fmt.Println(file.Filename)

		dst := fmt.Sprint("./files", file.Filename)
		c.SaveUploadedFile(file, dst)

		ioreader, _ := file.Open()

		defer ioreader.Close()

		content := make([]byte, 10) //content设定不能过大，否则文件末尾会有很多/u0000补充
		ioreader.Read(content)
		c.JSON(http.StatusOK, string(content))

	})

	r.GET("/uploadmany", func(c *gin.Context) {
		c.HTML(http.StatusOK, "uploadmany.html", gin.H{
			"title": "nothing",
			"ok":    "okkkkk",
		})
	})

	r.POST("/uploadmany", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["f1"]
		for _, file := range files {
			dst := fmt.Sprint("./", file.Filename) //库函数只能上传到当前root目录下
			fmt.Println(dst)
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	//重定向
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.liwenzhou.com/posts/Go/Gin_framework/#autoid-0-6-2")
	})

	r.GET("/test2", func(c *gin.Context) {
		c.Request.URL.Path = "/uploadmany"
		r.HandleContext(c) //路由跳转，重定向
	})

	r.LoadHTMLGlob("templates/*.html")

	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"message": "ok",
			})
		})
		userGroup.GET("/home", func(c *gin.Context) {
			c.HTML(http.StatusOK, "home.html", gin.H{
				"su": "true",
			})
		})
	}

	r.LoadHTMLGlob("templates/*/*.html")

	//errror???????????
	subGroup := userGroup.Group("/sub")
	{
		subGroup.GET("/posts/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "posts/index.html", gin.H{
				"everything": "ok",
				"title":      "主题",
			})
		})
		subGroup.GET("/users/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "users/index.html", gin.H{ //name 前面不能加"/"
				"title": "theme",
			})
		})
	}

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "users/index.html", gin.H{
			"err": "NOT FOUND",
		})
	})

	r.Run(":8083")
}
