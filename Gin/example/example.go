package example

import "github.com/gin-gonic/gin"

//Testexample ...
func Testexample() {
	//启动默认的路由引擎
	r := gin.Default()

	//GET:请求方式 /hello:请求路由
	//处理函数
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":     "Hello World",
			"are you ok?": "饮茶先了，做卵子做啥",
			"做卵子":         "sure",
		})
	})

	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})

	r.Run(":9090")
}
