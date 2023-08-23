package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery()) // middleware recovery from panics and write

	//authorization group
	/*
		authorized := router.Group("/")
		authorized.Use(AuthRequired())
		{
			authorized.POST("/login", loginEndpoint)
			authorized.POST("/submit", submitEndpoint)
			authorized.POST("/read", readEndpoint)

			// nested group
			testing := authorized.Group("testing")
			// visit 0.0.0.0:8080/testing/analytics
			testing.GET("/analytics", analyticsEndpoint)
		}

		router.Run(":8080")

	*/

}
