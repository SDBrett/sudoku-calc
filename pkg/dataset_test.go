package sudokucalc

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
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

func TestFullDataSet(t *testing.T) {

	mockDS := importDataSet(t)
	testDS := GenerateDataSet()

	if !reflect.DeepEqual(testDS, mockDS) {
		t.Errorf("generated DS does not match imported")
	}
}

func TestDataSetQuery(t *testing.T) {

	dsq := DataSetQuery{
		NumberOfDigits:   4,
		Value:            21,
		NumbersToInclude: NumberList{"1"},
		NumbersToExclude: NumberList{"5", "6"},
	}

	ds := GenerateDataSet()

	got, err := ds.Query(dsq)
	if err != nil {
		t.Errorf("got error %s", err)
	}

	fmt.Printf("%v", got)

}

func importDataSet(t testing.TB) DataSet {
	t.Helper()
	var ds DataSet
	path := "testdata/dataset_testdata1"
	i, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("error reading file %s", path)
	}

	err = json.Unmarshal(i, &ds)

	if err != nil {
		t.Errorf("unable to create %T from test dataset", ds)
	}

	return ds

}
