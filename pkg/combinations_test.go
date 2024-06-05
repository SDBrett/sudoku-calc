package sudokucalc

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

type ValidCombinationTestCase struct {
	Name             string
	Combinations     []string
	NumbersToInclude []string
	NumbersToExclude []string
	Expected         []string
}

func TestGenerateNumberLists(t *testing.T) {
	t.Run("Test two digit list", func(t *testing.T) {

		got := generateNumberLists(1, nil)
		want := []string{"12", "13", "14", "15", "16", "17", "18", "19", "23", "24", "25", "26", "27", "28", "29", "34", "35", "36", "37", "38", "39", "45", "46", "47", "48", "49", "56", "57", "58", "59", "67", "68", "69", "78", "79", "89"}

		assertNumberList(t, got, want)
	})
}

func TestValidation(t *testing.T) {

	numberRangeTestCases := []NumberRangeTestCases{
		{
			Name:        "Valid number within range",
			Min:         1,
			Max:         45,
			Given:       9,
			ExpectError: false,
		},
		{
			Name:        "Valid number below range",
			Min:         1,
			Max:         45,
			Given:       -1,
			ExpectError: true,
		},
		{
			Name:        "Valid number above range",
			Min:         1,
			Max:         9,
			Given:       10,
			ExpectError: true,
		},
	}

	for _, test := range numberRangeTestCases {
		t.Run(test.Name, func(t *testing.T) {
			got := ValidateNumberRange(test.Min, test.Max, test.Given)
			assertError(t, got, test.ExpectError)
		})
	}
}

func TestCombinations(t *testing.T) {

	got := GenerateDataSet()
	want := DataSet{
		4: {
			21: []string{"1389", "1479", "1569", "1578", "2379", "2469", "2478", "2568", "3459", "3468", "3567"},
		},
		7: {
			30: []string{"1234569", "1234578"},
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

func TestGetValidCombinations(t *testing.T) {
	testCases := []ValidCombinationTestCase{
		{
			Name:             "4 digit combinations for 21",
			Combinations:     []string{"1389", "1479", "1569", "1578", "2379", "2469", "2478", "2568", "3459", "3468", "3567"},
			NumbersToInclude: []string{"1"},
			NumbersToExclude: []string{"5", "6"},
			Expected:         []string{"1389", "1479"},
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

func assertNumberList(t testing.TB, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
