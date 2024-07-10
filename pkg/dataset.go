package sudokucalc

import (
	"log"
)

type DataSetQuery struct {
	NumberOfDigits   int      `json:"numberOfDigits"`  //The number of digits
	Value            int      `json:"value"`           //The Value for the digits to sum to
	NumbersToExclude []string `json:"excludedNumbers"` //Numbers which cannot be part of the solution
	NumbersToInclude []string `json:"includedNumbers"` //Numbers which must be part of the solution
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

func (ds DataSet) Query(dsq DataSetQuery) ([]string, error) {

	err := dsq.Validate()
	if err != nil {
		return nil, err
	}

	combinations := ds[dsq.NumberOfDigits][dsq.Value]

	nl := GetValidCombinations(combinations, dsq.NumbersToExclude, dsq.NumbersToInclude)

	return nl, nil
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
