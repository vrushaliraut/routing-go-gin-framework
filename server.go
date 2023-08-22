package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK,
			gin.H{
				"message": "Hello World..!",
			})
	})

	//post
	r.POST("users", func(context *gin.Context) {
		//create a user
	})

	//put
	r.PUT("/users/:id", func(context *gin.Context) {
		//update an existing user
	})

	r.DELETE("/users/:id", func(context *gin.Context) {
		//delete an existing user
	})

	//Handle Dynamic parameters
	r.GET("/users/:id", func(context *gin.Context) {
		id := context.Param("id")
		context.String(http.StatusOK, "Hello user ::-> %s", id)
	})

	//handle query parameters in routes
	r.GET("/users", func(context *gin.Context) {
		name := context.Query("name")
		age := context.Query("age")

		//handle request with name and age query parameters
		context.String(http.StatusOK, "Hello World with name..! %s %s ", name, age)
	})

	//query + post form
	r.POST("/post", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page", "0")
		name := context.PostForm("name")
		message := context.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	//handle route group
	v1 := r.Group("/v1")
	{
		v1.GET("/users", func(context *gin.Context) {
			//handle request for version 1 of users route
		})

		v1.GET("/articles", func(context *gin.Context) {
			//handle request for version 1 of users route
		})
	}
	r.Run()
	//Advanced routing
	// HTTP redirects

	/*router := gin.Default()
	r.GET("/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.google.com/")
	})

	r.Run(":8080")*/

}
