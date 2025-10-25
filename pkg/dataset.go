package sudokucalc

import (
	"errors"
	"fmt"
	"log"
)

type DataSetQuery struct {
	NumberOfDigits                   int      `json:"numberOfDigits"`                             //The number of digits
	Value                            int      `json:"value"`                                      //The Value for the digits to sum to
	NumbersToExclude                 []string `json:"excludedNumbers"`                            //Numbers which cannot be part of the solution
	NumbersToInclude                 []string `json:"includedNumbers"`                            //Numbers which must be part of the solution
	GetNumbersNotPresent             bool     `json:"getNumbersNotPresent,omitempty"`             //Return list of numbers which are not present in any result
	GetNumbersPresentAllCombinations bool     `json:"getNumbersPresentAllCombinations,omitempty"` //Return list of numbers which are present in all combinations
}

type DataSetQueryResponse struct {
	Combinations                    []string `json:"combinations"`                              // Valid Combinations
	DigitsInAllCombinations         []string `json:"digitsInAllCombinations,omitempty"`         //Digits which are present in all combinations
	DigitsAbsentFromAllCombinations []string `json:"digitsAbsentFromAllCombinations,omitempty"` //Digits which are absent from all combinations
}

// Generate the DataSet of all possible combinations
func GenerateDataSet() DataSet {

	// Initialize dataSet
	dataSet := newDataSet()
	numberList := all

	for numberOfDigits := 2; numberOfDigits < 10; numberOfDigits++ {
		idx := 1
		numberList = generateNumberLists(idx, numberList)
		for _, combination := range numberList {
			value, err := GetValueOfCombination(combination)
			if err != nil {
				log.Fatalf("number combination %s list contains invalid character", combination)
			}
			dataSet.UpdateValueCombination(numberOfDigits, value, combination)
		}
		idx++
	}
	return dataSet
}

func (ds DataSet) UpdateValueCombination(NumberOfDigits, Value int, Combination string) {
	nl, ok := ds[NumberOfDigits][Value]
	if ok {
		ds[NumberOfDigits][Value] = append(nl, Combination)
	} else {
		ds[NumberOfDigits][Value] = []string{Combination}
	}
}

func (ds DataSet) Query(dsq DataSetQuery) (DataSetQueryResponse, error) {

	response := DataSetQueryResponse{}

	log.Printf("DEBUG: Received query - GetNumbersNotPresent: %v, GetNumbersPresentAllCombinations: %v", dsq.GetNumbersNotPresent, dsq.GetNumbersPresentAllCombinations)

	err := dsq.Validate()
	if err != nil {
		return response, err
	}

	combinations := ds[dsq.NumberOfDigits][dsq.Value]

	response.Combinations = GetValidCombinations(combinations, dsq.NumbersToExclude, dsq.NumbersToInclude)
	if len(response.Combinations) == 0 {
		return response, errors.New("no combinations found")
	}

	if dsq.GetNumbersPresentAllCombinations || dsq.GetNumbersNotPresent {
		combinationAnalysis := AnalyzeCombinations(response.Combinations, dsq.NumberOfDigits)
		if dsq.GetNumbersNotPresent {
			response.DigitsAbsentFromAllCombinations = combinationAnalysis.NumbersNotPresent
		}
		if dsq.GetNumbersPresentAllCombinations {
			response.DigitsInAllCombinations = combinationAnalysis.NumbersInAllCombinations
		}
	}

	return response, nil
}

func (dsq DataSetQuery) Validate() error {

	var err error
	var x int

	err = ValidateNumberRange(2, 9, dsq.NumberOfDigits)
	if err != nil {
		return err
	}

	// Calculate valid range based on number of digits
	minSum := calculateMinSum(dsq.NumberOfDigits)
	maxSum := calculateMaxSum(dsq.NumberOfDigits)
	err = ValidateNumberRange(minSum, maxSum, dsq.Value)
	if err != nil {
		return fmt.Errorf("value %d is outside valid range for %d digits (%d-%d)", dsq.Value, dsq.NumberOfDigits, minSum, maxSum)
	}

	for _, v := range dsq.NumbersToInclude {
		x, err = StringToInt(v)
		if err != nil {
			return err
		}
		err = ValidateNumberRange(1, 9, x)
		if err != nil {
			return err
		}
	}

	for _, v := range dsq.NumbersToExclude {
		x, err = StringToInt(v)
		if err != nil {
			return err
		}
		err = ValidateNumberRange(1, 9, x)
		if err != nil {
			return err
		}
	}

	return nil
}

// calculateMinSum calculates the minimum possible sum for a given number of digits
func calculateMinSum(numberOfDigits int) int {
	sum := 0
	for i := 1; i <= numberOfDigits; i++ {
		sum += i
	}
	return sum
}

// calculateMaxSum calculates the maximum possible sum for a given number of digits
func calculateMaxSum(numberOfDigits int) int {
	sum := 0
	for i := 9 - numberOfDigits + 1; i <= 9; i++ {
		sum += i
	}
	return sum
}

func newDataSet() DataSet {

	// value combinations can only contain 2 - 9 digits.
	// Declared here to remove need for add / update functions for dataset
	dataSet := DataSet{
		2: map[int][]string{},
		3: map[int][]string{},
		4: map[int][]string{},
		5: map[int][]string{},
		6: map[int][]string{},
		7: map[int][]string{},
		8: map[int][]string{},
		9: map[int][]string{},
	}

	return dataSet
}

// CombinationAnalysis holds the results of analyzing combinations
type CombinationAnalysis struct {
	NumbersNotPresent        []string
	NumbersInAllCombinations []string
}

// AnalyzeCombinations analyzes the given combinations to find numbers not present and numbers present in all combinations
func AnalyzeCombinations(combinations []string, numberOfDigits int) CombinationAnalysis {
	analysis := CombinationAnalysis{
		NumbersNotPresent:        []string{},
		NumbersInAllCombinations: []string{},
	}

	if len(combinations) == 0 {
		// If no combinations, all numbers 1-9 are not present
		for i := 1; i <= 9; i++ {
			analysis.NumbersNotPresent = append(analysis.NumbersNotPresent, fmt.Sprintf("%d", i))
		}
		return analysis
	}

	// Track which numbers appear in each combination
	numberCounts := make(map[string]int)
	totalCombinations := len(combinations)

	// Count occurrences of each number across all combinations
	for _, combination := range combinations {
		for _, char := range combination {
			number := string(char)
			numberCounts[number]++
		}
	}

	// Find numbers not present in any combination
	for i := 1; i <= 9; i++ {
		numberStr := fmt.Sprintf("%d", i)
		if numberCounts[numberStr] == 0 {
			analysis.NumbersNotPresent = append(analysis.NumbersNotPresent, numberStr)
		}
	}

	// Find numbers present in all combinations
	for i := 1; i <= 9; i++ {
		numberStr := fmt.Sprintf("%d", i)
		if numberCounts[numberStr] == totalCombinations {
			analysis.NumbersInAllCombinations = append(analysis.NumbersInAllCombinations, numberStr)
		}
	}

	return analysis
}
