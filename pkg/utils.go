package sudokucalc

import (
	"sort"
	"strconv"
	"strings"
)

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func StringToInt(i string) (int, error) {
	x, err := strconv.Atoi(i)
	if err != nil {
		return x, err
	}
	return x, nil
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
