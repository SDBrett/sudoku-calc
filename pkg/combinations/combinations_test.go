package combinations

import (
	"reflect"
	"testing"
)

func TestGetValidCombinations(t *testing.T) {
	t.Run("Test two digit list", func(t *testing.T) {

		got := GetCombinations(1, nil)
		want := NumberList{"12", "13", "14", "15", "16", "17", "18", "19", "23", "24", "25", "26", "27", "28", "29", "34", "35", "36", "37", "38", "39", "45", "46", "47", "48", "49", "56", "57", "58", "59", "67", "68", "69", "78", "79", "89"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

func TestGetValues(t *testing.T) {

	got := GetValues()

	want := DigitCombinations{}
	want[3] = NumberList{"12"}
	want[4] = NumberList{"13"}
	want[5] = NumberList{"14", "23"}
	want[6] = NumberList{"15", "24", "123"}
	want[45] = NumberList{"123456789"}

	for k, v := range want {
		if !reflect.DeepEqual(got[k], v) {
			t.Errorf("got beep %v want %v", got[k], v)
		}
	}

}

func TestCombinations(t *testing.T) {
	dc := DigitCombinations{}
	dc[3] = NumberList{"12"}
	dc[4] = NumberList{"13"}
	dc[5] = NumberList{"14", "23"}
	dc[6] = NumberList{"15", "24"}

	got := Combinations{}

	got.Add(2, dc)

	want := Combinations{}
	want[2] = dc

}