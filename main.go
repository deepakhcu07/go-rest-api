package main

import (
	"context"
	"fmt"
	"github.com/deepakhcu07/go-rest-api/api"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	host = "0.0.0.0"
	port = 34567
)


func main() {
	fmt.Printf("\n Starting Go Rest API Server \n ")

	// Initialize all the singleton object

	// Starting the Server
	// Starting the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		fmt.Printf(" System Call: %v, Received at : %d", oscall, time.Now().Unix())
		cancel()
	}()

	if err := serve(ctx); err != nil {
		fmt.Printf("\n Error occurs while booting the server : %v ", err)
		log.Fatalf("error occurs while booting the server, Error: %v ", err)
	}

}


func serve(ctx context.Context) (err error) {
	router, err := initRouter()
	if err != nil {
		return err
	}

	serviceHost := fmt.Sprintf("%s:%d", host, port)
	srv := &http.Server{
		Addr:    serviceHost,
		Handler: router,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error on listening:%s \n ", err)
		}
	}()

	fmt.Printf("\n Server Started : %v", serviceHost)

	<-ctx.Done()

	fmt.Printf("\n Server Stopped ")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutdown); err != nil {
		fmt.Printf("\n Server shutdown failed ..: %s", err)
	}

	fmt.Printf("\n Server Exited Properly ")

	if err == http.ErrServerClosed {
		err = nil
	}
	return
}


func initRouter() (*gin.Engine, error) {
	r := gin.Default()



	if err := api.Health().Routes(r); err != nil {
		return nil, fmt.Errorf("error occured while initializing health routes. %w", err)
	}

	return r, nil
}
