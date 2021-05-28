package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//NormalRoute ...
func NormalRoute() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	r.GET("/test2", func(c *gin.Context) {
		c.Request.URL.Path = "/test"
		r.HandleContext(c)
	})

	r.Run(":8084")
}
