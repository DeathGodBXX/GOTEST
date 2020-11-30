package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//将前缀相同的路由放在一个group下
	v1 := r.Group("/v1")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	r.Run(":8080")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "Nick")
	c.String(http.StatusOK, fmt.Sprintln(name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "小明")
	c.String(http.StatusOK, fmt.Sprintln(name))
}
