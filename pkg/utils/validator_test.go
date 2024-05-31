package utils

import "testing"

type NumberRangeTestCases struct {
	Name   string
	Min    int
	Max    int
	Given  int
	Expect error
}

type SearchTypeTestCases struct {
	Name   string
	Given  string
	Expect error
}

func TestValidator(t *testing.T) {

	numberRangeTestCases := []NumberRangeTestCases{
		{
			Name:   "Valid number within range",
			Min:    1,
			Max:    45,
			Given:  9,
			Expect: nil,
		},
		{
			Name:   "Valid number below range",
			Min:    1,
			Max:    45,
			Given:  -1,
			Expect: ErrExceedRange,
		},
		{
			Name:   "Valid number above range",
			Min:    1,
			Max:    9,
			Given:  10,
			Expect: ErrExceedRange,
		},
	}

	for _, test := range numberRangeTestCases {
		t.Run(test.Name, func(t *testing.T) {
			got := ValidateNumberRange(test.Min, test.Max, test.Given)
			assertError(t, got, test.Expect)
		})
	}

	searchTypeTestCases := []SearchTypeTestCases{
		{
			Name:   "valid search type",
			Given:  "FULL",
			Expect: nil,
		},
		{
			Name:   "invalid search type",
			Given:  "testing",
			Expect: ErrInvalidSearchType,
		},
	}
	for _, test := range searchTypeTestCases {
		t.Run(test.Name, func(t *testing.T) {
			got := ValidateSearchType(test.Given)
			assertError(t, got, test.Expect)
		})
	}

}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
