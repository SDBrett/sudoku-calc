package combinations

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/sdbrett/sudoku-calc/pkg/utils"
)

type dataSetQueryTestCase struct {
	Name   string
	DSQ    DataSetQuery
	Expect error
}

type ValidCombinationTestCase struct {
	Name             string
	Combinations     NumberList
	NumbersToInclude NumberList
	NumbersToExclude NumberList
	Expected         NumberList
}

func TestGenerateNumberLists(t *testing.T) {
	t.Run("Test two digit list", func(t *testing.T) {

		got := generateNumberLists(1, nil)
		want := NumberList{"12", "13", "14", "15", "16", "17", "18", "19", "23", "24", "25", "26", "27", "28", "29", "34", "35", "36", "37", "38", "39", "45", "46", "47", "48", "49", "56", "57", "58", "59", "67", "68", "69", "78", "79", "89"}

		assertNumberList(t, got, want)
	})

}

func TestGetValues(t *testing.T) {

	got := getCombinationsForValue()

	want := ValueCombinations{}
	want[3] = NumberList{"12"}
	want[4] = NumberList{"13"}
	want[5] = NumberList{"14", "23"}
	want[6] = NumberList{"15", "24", "123"}
	want[45] = NumberList{"123456789"}

	for k, v := range want {
		assertNumberList(t, got[k], v)
	}
}

func TestCombinations(t *testing.T) {

	got := GenerateDataSet()
	want := DataSet{
		4: ValueCombinations{
			21: NumberList{"1389", "1479", "1569", "1578", "2379", "2469", "2478", "2568", "3459", "3468", "3567"},
		},
		7: ValueCombinations{
			30: NumberList{"1234569", "1234578"},
		},
	}

	for k, v := range want {
		for x, y := range v {
			if !reflect.DeepEqual(got[k][x], y) {
				t.Errorf("got %v want %v", got[k][x], y)
			}
		}
	}
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

func TestGetValidCombinations(t *testing.T) {
	testCases := []ValidCombinationTestCase{
		{
			Name:             "4 digit combinations for 21",
			Combinations:     NumberList{"1389", "1479", "1569", "1578", "2379", "2469", "2478", "2568", "3459", "3468", "3567"},
			NumbersToInclude: NumberList{"1"},
			NumbersToExclude: NumberList{"5", "6"},
			Expected:         NumberList{"1389", "1479"},
		},
	}

	for _, test := range testCases {

		t.Run(test.Name, func(t *testing.T) {
			got := GetValidCombinations(test.Combinations, test.NumbersToExclude, test.NumbersToInclude)
			sort.Strings(got)
			fmt.Println(got)
			assertNumberList(t, got, test.Expected)
		})
	}
}

func assertNumberList(t testing.TB, got, want NumberList) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
