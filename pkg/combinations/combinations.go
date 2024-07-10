package combinations

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
type DigitCombinations map[int]NumberList
type Combinations map[int]DigitCombinations

func GenerateDataSet() Combinations {
	combinationsByValue := GetValues()

	dataSet := Combinations{
		2: DigitCombinations{},
		3: DigitCombinations{},
		4: DigitCombinations{},
		5: DigitCombinations{},
		6: DigitCombinations{},
		7: DigitCombinations{},
		8: DigitCombinations{},
		9: DigitCombinations{},
	}

	// k is the sum of the items in the array
	// v is the list of combinations
	for k, v := range combinationsByValue {
		for i := 0; i < len(v); i++ {
			// v[i] is an individual combination for total k
			dataSet[len(v[i])].Update(k, v[i])
		}
	}

	return dataSet
}

func GetCombinations(idx int, nl NumberList) NumberList {

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

func GetValues() DigitCombinations {

	dc := DigitCombinations{}
	value := 0
	asInt := 0
	nl := all
	numberOfDigits := 2
	idx := 1
	for numberOfDigits < 10 {
		nl = GetCombinations(idx, nl)

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

func (dc DigitCombinations) Search(value int) (NumberList, error) {

	nl, ok := dc[value]

	if !ok {
		return nil, ErrNotFound
	}
	return nl, nil
}

func (dc DigitCombinations) Update(value int, combination string) error {

	nl, err := dc.Search(value)
	switch err {
	case ErrNotFound:
		nl := NumberList{combination}
		dc[value] = nl
	case nil:
		nl = append(nl, combination)
		dc[value] = nl
	default:
		return err
	}

	return nil
}

func (e ValueErr) Error() string {
	return string(e)
}

func (combo Combinations) Search(value int) (DigitCombinations, error) {

	nl, ok := combo[value]

	if !ok {
		return nil, ErrNotFound
	}
	return nl, nil
}
