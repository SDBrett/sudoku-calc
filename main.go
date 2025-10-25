package main

import (
	"net/http"
	"os"
	"strconv"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	response, err := ds.Query(dsq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func main() {
	// Get port from environment variable, default to 8080
	port := os.Getenv("CALC_PORT")
	if port == "" {
		port = "8080"
	}

	// Validate port is a valid number
	if _, err := strconv.Atoi(port); err != nil {
		port = "8080"
	}

	r := gin.Default()

	// Serve static files (CSS, JS)
	r.Static("/static", "./static")

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Serve the main web interface
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// API endpoints
	r.GET("/dataset", getDataset)
	r.POST("/dataset", queryDataSet)

	// Listen and Server on the configured port
	r.Run(":" + port)
}
