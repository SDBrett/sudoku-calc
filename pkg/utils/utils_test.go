package utils

import (
	"testing"
)

func TestIntToString(t *testing.T) {
	got := IntToString(1234)
	want := "1234"

	assertStrings(t, got, want)
}

func TestStringToInt(t *testing.T) {

	t.Run("no error", func(t *testing.T) {
		got, err := StringToInt("1234")
		want := 1234

		if err != nil {
			t.Errorf("expected no error, go %s", err)
		}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("invalid input", func(t *testing.T) {
		_, err := StringToInt("12abc34")

		if err == nil {
			t.Errorf("expected error but did not get one")
		}
	})
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
