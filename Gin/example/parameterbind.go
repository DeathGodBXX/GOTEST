package example

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Login struct defined
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	ID       int
}

//ParameterBind explaination
func ParameterBind() {
	r := gin.Default()

	//JSON数据解析
	r.POST("/LoginJSON", func(c *gin.Context) {
		var login Login
		//shouldBind根据Content-type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"User":     login.User,
				"Password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})

	//form表单解析
	r.POST("/LoginFORM", func(c *gin.Context) {
		var login Login
		//ShouldBind根据Content-type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("form-data login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"User":     login.User,
				"Password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})

	r.GET("/LoginFORM", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"User":     login.User,
				"Password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}
	})

	r.Run(":8082")
}
