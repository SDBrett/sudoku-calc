package utils

import (
	"sort"
	"strconv"
	"strings"
)

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
