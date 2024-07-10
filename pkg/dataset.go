package sudokucalc

import (
	"github.com/sdbrett/sudoku-calc/pkg/utils"
)

type DataSet map[int]ValueCombinations

type DataSetQuery struct {
	NumberOfDigits   int        //The number of digits
	Value            int        //The Value for the digits to sum to
	NumbersToExclude NumberList //Numbers which cannot be part of the solution
	NumbersToInclude NumberList //Numbers which must be part of the solution
}

// Generate the DataSet of all possible combinations
func GenerateDataSet() DataSet {
	combinationsByValue := getCombinationsForValue()

	// Initialize dataSet
	dataSet := newDataSet()

	// k is the sum of the items in the array
	// v is the list of combinations
	for k, v := range combinationsByValue {
		for i := 0; i < len(v); i++ {
			// v[i] is an individual combination for total k
			dataSet[len(v[i])].Update(k, v[i])
		}
	}

	return dataSet
}

// Returns possible values and their combinations with a given
// number of digits.
func (dc DataSet) Search(value int) (ValueCombinations, error) {

	nl, ok := dc[value]

	if !ok {
		return nil, ErrNotFound
	}
	return nl, nil
}

func (dc DataSet) Query(dsq DataSetQuery) (NumberList, error) {

	nl := NumberList{}

	err := dsq.Validate()
	if err != nil {
		return nil, err
	}

	return nl, nil
}

func (dsq DataSetQuery) Validate() error {

	var err error
	var x int

	err = utils.ValidateNumberRange(1, 9, dsq.NumberOfDigits)
	if err != nil {
		return err
	}

	err = utils.ValidateNumberRange(3, 45, dsq.Value)
	if err != nil {
		return err
	}

	for _, v := range dsq.NumbersToInclude {
		x, err = utils.StringToInt(v)
		if err != nil {
			return err
		}
		err = utils.ValidateNumberRange(1, 9, x)
		if err != nil {
			return err
		}
	}

	for _, v := range dsq.NumbersToExclude {
		x, err = utils.StringToInt(v)
		if err != nil {
			return err
		}
		err = utils.ValidateNumberRange(1, 9, x)
		if err != nil {
			return err
		}
	}

	return nil
}

func newDataSet() DataSet {
	dataSet := DataSet{
		2: ValueCombinations{},
		3: ValueCombinations{},
		4: ValueCombinations{},
		5: ValueCombinations{},
		6: ValueCombinations{},
		7: ValueCombinations{},
		8: ValueCombinations{},
		9: ValueCombinations{},
	}

	return dataSet
}
