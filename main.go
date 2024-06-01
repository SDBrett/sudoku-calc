package main

import (
	"encoding/json"
	"fmt"

	sudokucalc "github.com/sdbrett/sudoku-calc/pkg"
)

func main() {

	dataSet := sudokucalc.GenerateDataSet()

	jsonString, _ := json.Marshal(dataSet)

	fmt.Printf("%s", jsonString)
}
