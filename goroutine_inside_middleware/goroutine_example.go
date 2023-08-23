package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	router := gin.Default()

	router.GET("/long_async", func(context *gin.Context) {
		contextCopy := context.Copy()

		go func() {
			time.Sleep(5 * time.Second) // simulate long task with time sleep

			log.Println("Done! in path" + contextCopy.Request.URL.Path) // use copied context in goroutine
		}()
	})

	router.GET("/long_sync", func(context *gin.Context) {

		time.Sleep(5 * time.Second) // simulate long task with time sleep

		log.Println("Done! in path" + context.Request.URL.Path) // use copied context in goroutine
	})

	router.Run(":8080")
}
