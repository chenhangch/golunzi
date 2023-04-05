package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/",func(ctx *gin.Context) {
		time.Sleep(5* time.Second)
		ctx.String(http.StatusOK, "Welcome")
	})

	srv := &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe();err != nil && err != http.ErrServerClosed{
			log.Fatal("listen: %s\n",err)
		}	
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit,os.Interrupt)
	<-quit
	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx);err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server Exiting")

}