package utils

import "testing"

func TestIntToString(t *testing.T) {
	got := IntToString(1234)
	want := "1234"

	if got != want {
		t.Errorf("got %v with type of %T want %s with type of %T", got, got, want, want)
	}
}
