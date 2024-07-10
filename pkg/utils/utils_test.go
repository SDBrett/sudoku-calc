package utils

import (
	"testing"
)

func TestIntToString(t *testing.T) {
	got := IntToString(1234)
	want := "1234"

	assertStrings(t, got, want)
}

func TestSortString(t *testing.T) {

	cases := []struct {
		Name     string
		Input    string
		Expected string
	}{{
		"sort only letters",
		"test",
		"estt",
	},
		{
			"sort only numbers",
			"563478",
			"345678",
		},
		{
			"sort only numbers",
			"563478",
			"345678",
		},
		{
			"sort letters and numbers",
			"568te6422436st",
			"2234456668estt",
		},
		{
			"sort letters and numbers with mixed case",
			"ZR5AD68te6422436st",
			"2234456668ADRZestt",
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := SortString(test.Input)
			assertStrings(t, got, test.Expected)
		})
	}

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}
}
