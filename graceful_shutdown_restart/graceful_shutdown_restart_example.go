package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//without context

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		context.String(http.StatusOK, "Welcome Gin Server")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	//Initialize server in goroutine so that it won't block graceful shutdown

	go func() {
		err := server.ListenAndServe()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s \n", err)
		}
	}()

	//wait for interrupt signal to gracefullt shutsown the server with a timeout of 5 seconds

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal("Server forced to shutdown", err)
	}
	log.Println("Shutting existing")
}
