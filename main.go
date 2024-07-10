package main

import (
	"net/http"

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

	r := gin.Default()
	r.GET("/dataset", getDataset)
	r.POST("/dataset", queryDataSet)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
