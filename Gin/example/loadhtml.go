package example

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

//LoadHTML ...
func LoadHTML() {
	r := gin.Default()

	//use LoadHTMLGlob() and LoadHTMLFiles() to render HTML
	r.LoadHTMLGlob("templates/**/*")
	// r.LoadHTMLFiles() //must some strings ,or return false

	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})

	r.Run(":8080")
}

//DefinedSelfHTML ...
func DefinedSelfHTML() {
	r := gin.Default()

	// /static是路由对应于文件./static,这样使用html就能正确找到文件,加载static文件下所有静态文件
	// /static/favicon.ico想正确显示，必须删除缓存，重新进入index页面，最好换一个浏览器登入。
	// first is 路由，second是当前文件系统搜索路径
	r.Static("/static", "./static")

	// r.StaticFile("/static/favicon.ico", "./static/favicon.ico")
	//设置自定义的funcMap，tmpl模板通过调用safe渲染页面
	r.SetFuncMap(template.FuncMap{"safe": func(str string) template.HTML {
		return template.HTML(str)
	}})

	r.LoadHTMLFiles("templates/index.tmpl")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<href> Are you OK??<href>") //最后一个指定在tmpl使用的string信息
	})

	r.GET("/json", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg.Name = "大黄"
		msg.Message = "yellow"
		msg.Age = 18
		c.JSON(http.StatusOK, msg)
	})

	r.POST("/json", func(c *gin.Context) {
		//从c.Request.Body读取数据
		b, _ := c.GetRawData()

		//解析json数据为map或者结构体
		var m map[string]interface{}
		_ = json.Unmarshal(b, &m)
		c.JSON(http.StatusOK, m)
	})
	//需要下载google的goprotobuf库
	// r.GET("/protobuf", func(c *gin.Context) {
	// 	// type Test struct {
	// 	// 	Label *string
	// 	// 	Reps  []int64
	// 	// }
	// 	reps := []int64{int64(1), int64(2)}
	// 	label := "test"
	// 	data := &protoexample.Test{
	// 		Label: &label,
	// 		Reps:  reps,
	// 	}
	// 	c.ProtoBuf(http.StatusOK, data)
	// })

	//获取路由？之后的querystring
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "大黄")
		address := c.Query("address")
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	//发送POST请求，使用form表单提交数据
	r.POST("/user/search", func(c *gin.Context) {
		username := c.PostForm("username")
		address := c.PostForm("address")
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	//发送GET请求，从url路径中获取参数
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"address":  address,
		})
	})
	//每个engine使用不同的端口号，否则会冲突
	r.Run(":8081")
}
