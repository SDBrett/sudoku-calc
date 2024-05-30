package main

import (
	"encoding/json"
	"fmt"

	combinations "github.com/sdbrett/sudoku-calc/pkg/combinations"
)

func main() {

	dataSet := combinations.GenerateDataSet()

	jsonString, _ := json.Marshal(dataSet)

	fmt.Printf("%s", jsonString)
}
