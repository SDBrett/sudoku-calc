package sudokucalc

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
)

type dataSetQueryTestCase struct {
	Name        string
	DSQ         DataSetQuery
	ExpectError bool
}

type NumberRangeTestCases struct {
	Name        string
	Min         int
	Max         int
	Given       int
	ExpectError bool
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
				NumbersToExclude: []string{"8", "9"},
				NumbersToInclude: []string{"1", "2"},
			},
			ExpectError: false,
		},
		{
			Name: "invalid query",
			DSQ: DataSetQuery{
				NumberOfDigits:   3,
				Value:            6,
				NumbersToExclude: []string{"8", "9", "45"},
				NumbersToInclude: []string{"1", "2"},
			},
			ExpectError: true,
		},
	}
	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			got := test.DSQ.Validate()
			assertError(t, got, test.ExpectError)
		})
	}
}

func TestCombinationValidation(t *testing.T) {

	t.Run("test invalid combination", func(t *testing.T) {
		_, err := GetValueOfCombination("12d")
		expectError := true

		assertError(t, err, expectError)
	})

	t.Run("test valid combination", func(t *testing.T) {
		got, err := GetValueOfCombination("1234")
		want := 10
		expectError := false

		assertError(t, err, expectError)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
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
		NumbersToInclude: []string{"1"},
		NumbersToExclude: []string{"5", "6"},
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

func assertError(t testing.TB, err error, want bool) {
	t.Helper()

	if want && err == nil {
		t.Errorf("expected error but didn't get one")
	}
	if !want && err != nil {
		t.Errorf("no error expected got %s", err)
	}
}
