package sudokucalc

import (
	"testing"

	"github.com/sdbrett/sudoku-calc/pkg/utils"
)

type dataSetQueryTestCase struct {
	Name   string
	DSQ    DataSetQuery
	Expect error
}

func BenchmarkGenerateDataSet(b *testing.B) {
	GenerateDataSet()
}

func TestDataSetValidation(t *testing.T) {

	testCases := []dataSetQueryTestCase{
		{
			Name: "valid query",
			DSQ: DataSetQuery{
				NumberOfDigits:   3,
				Value:            6,
				NumbersToExclude: NumberList{"8", "9"},
				NumbersToInclude: NumberList{"1", "2"},
			},
			Expect: nil,
		},
		{
			Name: "invalid query",
			DSQ: DataSetQuery{
				NumberOfDigits:   3,
				Value:            6,
				NumbersToExclude: NumberList{"8", "9", "45"},
				NumbersToInclude: NumberList{"1", "2"},
			},
			Expect: utils.ErrExceedRange,
		},
	}
	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			got := test.DSQ.Validate()
			assertError(t, got, test.Expect)
		})
	}

}
