package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/patil-prathamesh/Blogify/controllers"
	"github.com/patil-prathamesh/Blogify/middlware"
)

func PostRoutes(c *gin.Engine) {
	c.Use(middlware.Authenticate)
	c.Use(middlware.IsValid)
	posts := c.Group("/posts")
	{
		posts.POST("/", controllers.CreatePost)
		posts.PUT("/:post_id", controllers.UpdatePost)
		posts.DELETE("/:post_id", controllers.DeletePost)
		posts.DELETE("/", controllers.DeleteAllPosts)
		posts.GET("/", controllers.ListAllPosts)
	}

}
