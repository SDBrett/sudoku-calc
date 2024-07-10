package main

import (
	"encoding/json"
	"fmt"

	combinations "github.com/sdbrett/sudoku-calc/pkg/combinations"
)

func main() {

	twoDigitList := combinations.GetCombinations(1, nil)
	threeDigitList := combinations.GetCombinations(2, twoDigitList)
	fourDigitList := combinations.GetCombinations(3, threeDigitList)
	fiveDigitList := combinations.GetCombinations(4, fourDigitList)
	sixDigitList := combinations.GetCombinations(5, fiveDigitList)
	sevenDigitList := combinations.GetCombinations(6, sixDigitList)
	eightDigitList := combinations.GetCombinations(7, sevenDigitList)
	nineDigitList := combinations.GetCombinations(8, eightDigitList)

	twoDigitCombinations := combinations.GetValues(twoDigitList)
	threeDigitCombinations := combinations.GetValues(threeDigitList)
	fourDigitCombinations := combinations.GetValues(fourDigitList)
	fiveDigitCombinations := combinations.GetValues(fiveDigitList)
	sixDigitCombinations := combinations.GetValues(twoDigitList)
	sevenDigitCombinations := combinations.GetValues(sixDigitList)
	eightDigitCombinations := combinations.GetValues(eightDigitList)
	nineDigitCombinations := combinations.GetValues(nineDigitList)

	c := combinations.Combinations{}
	c.Add("2", twoDigitCombinations)
	c.Add("3", threeDigitCombinations)
	c.Add("4", fourDigitCombinations)
	c.Add("5", fiveDigitCombinations)
	c.Add("6", sixDigitCombinations)
	c.Add("7", sevenDigitCombinations)
	c.Add("8", eightDigitCombinations)
	c.Add("9", nineDigitCombinations)

	jsonString, _ := json.Marshal(c)

	fmt.Printf("%s", jsonString)
}
