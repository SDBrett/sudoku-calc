package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	sudokucalc "github.com/sdbrett/sudoku-calc/pkg"
)

var ds = sudokucalc.GenerateDataSet()

// getAlbums responds with the list of all albums as JSON.
func getDataset(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, ds)
}

func queryDataSet(c *gin.Context) {
	var dsq sudokucalc.DataSetQuery
	var err error
	if err = c.BindJSON(&dsq); err != nil {
		return
	}

	response, err := ds.Query(dsq)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusAccepted, response)
}

func main() {

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	r := gin.Default()
	r.GET("/dataset", getDataset)
	r.POST("/dataset", queryDataSet)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r.Handler(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-gracefulShutdown

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
