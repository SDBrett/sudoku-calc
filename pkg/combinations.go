package sudokucalc

import (
	"slices"
	"sort"
	"strings"
)

const (
	ErrNotFound          = ValueErr("could not find the value you were looking for")
	ErrValueExists       = ValueErr("cannot add value because it already exists")
	ErrValueDoesNotExist = ValueErr("cannot update value because it does not exist")
	ErrUpdating          = ValueErr("error updating")
)

var all = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

type ValueErr string
type ValidCombinations map[string]bool
type DataSet map[int]map[int][]string

// Generates a list of valid number combinations
func generateNumberLists(idx int, nl []string) []string {

	if nl == nil {
		nl = all
	}

	outlist := []string{}
	sorted := ""

	for i := idx; i < len(all); i++ {
		for x := 0; x < len(nl); x++ {
			if !strings.Contains(nl[x], all[i]) {
				sorted = SortString(nl[x] + all[i])
				if !slices.Contains(outlist, sorted) {
					outlist = append(outlist, sorted)
				}
			}
		}
	}
	sort.Strings(outlist)
	return outlist
}

func (e ValueErr) Error() string {
	return string(e)
}

func GetValidCombinations(combinations, exclude, include []string) []string {

	valid := ValidCombinations{}
	nl := []string{}

	for _, item := range combinations {
		valid[item] = true
	}

	// All combinations start as valid and then checked to determine if invalid
	for k := range valid {
		// Invalidate if combination contains an excluded number
		if valid[k] {
			valid[k] = ConfirmNoExcludedNumbersInCombinations(k, exclude)
		}

		// Invalidate if combination does not contain an included number
		if valid[k] {
			valid[k] = ConfirmCombinationContainsIncludedNumbers(k, include)
		}
		// if still valid add to the returned number list
		if valid[k] {
			nl = append(nl, k)
		}
	}
	return nl
}

func ConfirmCombinationContainsIncludedNumbers(candidateCombination string, includeNumberList []string) bool {
	for _, item := range includeNumberList {
		if !strings.Contains(candidateCombination, item) {
			return false
		}
	}
	return true
}

func ConfirmNoExcludedNumbersInCombinations(candidateCombination string, excludeNumberList []string) bool {
	for _, item := range excludeNumberList {
		if strings.Contains(candidateCombination, item) {
			return false
		}
	}
	return true
}
