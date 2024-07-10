package sudokucalc

import (
	"log"
	"strings"
)

type DataSetQuery struct {
	NumberOfDigits                     int      `json:"numberOfDigits"`                            //The number of digits
	Value                              int      `json:"value"`                                     //The Value for the digits to sum to
	NumbersToExclude                   []string `json:"excludedNumbers"`                           //Numbers which cannot be part of the solution
	NumbersToInclude                   []string `json:"includedNumbers"`                           //Numbers which must be part of the solution
	GetDigitsInAllCombinations         bool     `json:"getDigitsInAllCombinations,omitempty"`      //Return list of numbers present in all combinations
	GetDigitsAbsentFromAllCombinations bool     `json:"getDigitsNotFromAllCombinations,omitempty"` //Return list of numbers absent from all combinations
}

type DataSetQueryResponse struct {
	Combinations                  []string `json:"combinations"`                            // Valid Combinations
	DigitsInAllCombinations       []string `json:"digitsInAllCombinations,omitempty"`       //Digits which are present in all combinations
	DigitsAbsentInAllCombinations []string `json:"digitsAbsentInAllCombinations,omitempty"` //Digits which are absent from all combinations
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

	err := dsq.Validate()
	if err != nil {
		return response, err
	}

	combinations := ds[dsq.NumberOfDigits][dsq.Value]

	response.Combinations = GetValidCombinations(combinations, dsq.NumbersToExclude, dsq.NumbersToInclude)

	if dsq.GetDigitsInAllCombinations {
		response.GetDigitsInAllCombinations()
	}

	if dsq.GetDigitsAbsentFromAllCombinations {
		response.GetDigitsAbsentInAllCombinations()
	}

	return response, nil
}

func (response *DataSetQueryResponse) GetDigitsInAllCombinations() {

	eligibleNumbers := setupValidationNumberList(all)

	for k := range eligibleNumbers {
		if eligibleNumbers[k] {
			eligibleNumbers[k] = ConfirmDigitInCombinations(k, response.Combinations)
		}
		if eligibleNumbers[k] {
			response.DigitsInAllCombinations = append(response.DigitsInAllCombinations, k)
		}
	}
}

func ConfirmDigitInCombinations(testDigit string, combinations []string) bool {
	for _, item := range combinations {
		if !strings.Contains(item, testDigit) {
			return false
		}
	}
	return true
}

func (response *DataSetQueryResponse) GetDigitsAbsentInAllCombinations() {

	eligibleNumbers := setupValidationNumberList(all)
	response.DigitsAbsentInAllCombinations = []string{}

	for k := range eligibleNumbers {
		if eligibleNumbers[k] {
			eligibleNumbers[k] = ConfirmDigitAbsentFromAllCombinations(k, response.Combinations)
		}
		if eligibleNumbers[k] {
			response.DigitsAbsentInAllCombinations = append(response.DigitsAbsentInAllCombinations, k)
		}
	}
}

func ConfirmDigitAbsentFromAllCombinations(testDigit string, combinations []string) bool {
	for _, item := range combinations {
		if strings.Contains(item, testDigit) {
			return false
		}
	}
	return true
}

func (dsq DataSetQuery) Validate() error {

	var err error
	var x int

	err = ValidateNumberRange(1, 9, dsq.NumberOfDigits)
	if err != nil {
		return err
	}

	err = ValidateNumberRange(3, 45, dsq.Value)
	if err != nil {
		return err
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
