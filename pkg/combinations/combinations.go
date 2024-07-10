package combinations

import (
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
)

var all = NumberList{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

type NumberList []string
type ValueErr string
type DigitCombinations map[int]NumberList
type Combinations map[int]DigitCombinations

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
	//var strValue string
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
			//strValue = utils.IntToString(value)
			err := dc.Update(value, combination)
			if err == ErrValueDoesNotExist {
				dc.Add(value, combination)
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

func (dc DigitCombinations) Add(value int, combination string) error {
	_, err := dc.Search(value)
	switch err {
	case ErrNotFound:
		nl := NumberList{combination}
		dc[value] = nl
	case nil:
		return ErrValueExists
	default:
		return err
	}
	return nil
}

func (dc DigitCombinations) Update(value int, combination string) error {

	nl, err := dc.Search(value)
	switch err {
	case ErrNotFound:
		return ErrValueDoesNotExist
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

func (combo Combinations) Add(value int, dc DigitCombinations) error {
	_, err := combo.Search(value)
	switch err {
	case ErrNotFound:
		combo[value] = dc
	case nil:
		return ErrValueExists
	default:
		return err
	}
	return nil
}
