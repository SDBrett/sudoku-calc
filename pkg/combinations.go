package sudokucalc

import (
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/sdbrett/sudoku-calc/pkg/utils"
)

const (
	ErrNotFound          = ValueErr("could not find the value you were looking for")
	ErrValueExists       = ValueErr("cannot add value because it already exists")
	ErrValueDoesNotExist = ValueErr("cannot update value because it does not exist")
	ErrUpdating          = ValueErr("error updating")
)

var all = NumberList{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

type NumberList []string
type ValueErr string
type ValueCombinations map[int]NumberList
type ValidCombinations map[string]bool

// Generates a list of valid number combinations
func generateNumberLists(idx int, nl NumberList) NumberList {

	if nl == nil {
		nl = all
	}

	outlist := NumberList{}
	sorted := ""

	for i := idx; i < len(all); i++ {
		for x := 0; x < len(nl); x++ {
			if !strings.Contains(nl[x], all[i]) {
				sorted = utils.SortString(nl[x] + all[i])
				if !slices.Contains(outlist, sorted) {
					outlist = append(outlist, sorted)
				}
			}
		}
	}
	sort.Strings(outlist)
	return outlist
}

// Create a map of different values and the possible combinations
// of digits for that value
func getCombinationsForValue() ValueCombinations {

	dc := ValueCombinations{}
	value := 0
	asInt := 0
	nl := all
	numberOfDigits := 2
	idx := 1
	for numberOfDigits < 10 {
		nl = generateNumberLists(idx, nl)

		for _, combination := range nl {
			value = 0
			idx := 1
			for i := 0; i < len(combination); i++ {
				asInt, _ = strconv.Atoi(combination[i:idx])
				value += asInt
				idx++
			}

			err := dc.Update(value, combination)

			if err != ErrUpdating && err != nil {
				log.Fatal(err)
			}

		}

		numberOfDigits++
		idx++

	}

	return dc
}

// Look up number combinations for a given value
// Returns list of combinations for that value
func (vc ValueCombinations) Search(value int) (NumberList, error) {

	nl, ok := vc[value]

	if !ok {
		return nil, ErrNotFound
	}
	return nl, nil
}

// Update the combinations to a list of combinations for a a given value
func (vc ValueCombinations) Update(value int, combination string) error {

	nl, err := vc.Search(value)
	switch err {
	case ErrNotFound:
		nl := NumberList{combination}
		vc[value] = nl
	case nil:
		nl = append(nl, combination)
		vc[value] = nl
	default:
		return err
	}

	return nil
}

func (e ValueErr) Error() string {
	return string(e)
}

func GetValidCombinations(combinations, exclude, include NumberList) NumberList {

	valid := ValidCombinations{}
	nl := NumberList{}

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

func ConfirmCombinationContainsIncludedNumbers(candidateCombination string, includeNumberList NumberList) bool {
	for _, item := range includeNumberList {
		if !strings.Contains(candidateCombination, item) {
			return false
		}
	}
	return true
}

func ConfirmNoExcludedNumbersInCombinations(candidateCombination string, excludeNumberList NumberList) bool {
	for _, item := range excludeNumberList {
		if strings.Contains(candidateCombination, item) {
			return false
		}
	}
	return true
}
