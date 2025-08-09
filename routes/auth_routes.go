package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/patil-prathamesh/Blogify/controllers"
)

func AuthRoutes(c *gin.Engine) {
	c.POST("/login", controllers.Login)
	c.POST("/signup", controllers.SignUp)
	c.POST("/refresh_token", controllers.Login)
}