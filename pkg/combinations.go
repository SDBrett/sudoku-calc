package combinations

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const (
	ErrNotFound          = ValueErr("could not find the value you were looking for")
	ErrValueExists       = ValueErr("cannot add value because it already exists")
	ErrValueDoesNotExist = ValueErr("cannot update value because it does not exist")
)

type NumberList []string
type ValueErr string
type DigitCombinations map[string]NumberList
type Combinations map[string]DigitCombinations

func GetValidCombinations() {

	twoDigitList := GetCombinations(1, nil)
	threeDigitList := GetCombinations(2, twoDigitList)
	fourDigitList := GetCombinations(3, threeDigitList)
	fiveDigitList := GetCombinations(4, fourDigitList)
	sixDigitList := GetCombinations(5, fiveDigitList)
	sevenDigitList := GetCombinations(6, sixDigitList)
	eightDigitList := GetCombinations(7, sevenDigitList)
	nineDigitList := GetCombinations(8, eightDigitList)
	fmt.Println(nineDigitList)
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func GetCombinations(idx int, nl NumberList) NumberList {

	all := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	if nl == nil {
		nl = all
	}

	outlist := NumberList{}
	sorted := ""

	for i := idx; i < len(all); i++ {
		for x := 0; x < len(nl); x++ {
			if !strings.Contains(nl[x], all[i]) {
				sorted = sortString(nl[x] + all[i])
				if !slices.Contains(outlist, sorted) {
					outlist = append(outlist, sorted)
				}
			}
		}
	}
	sort.Strings(outlist)
	return outlist
}

func GetValues(nl NumberList) DigitCombinations {

	dc := DigitCombinations{}
	value := 0
	asInt := 0
	var strValue string
	for _, combination := range nl {
		value = 0
		idx := 1
		for i := 0; i < len(combination); i++ {
			asInt, _ = strconv.Atoi(combination[i:idx])
			value += asInt
			idx++
		}
		strValue = intToString(value)
		err := dc.Update(strValue, combination)
		if err == ErrValueDoesNotExist {
			dc.Add(strValue, combination)
		}

	}
	return dc

}

func (dc DigitCombinations) Search(value string) (NumberList, error) {

	nl, ok := dc[value]

	if !ok {
		return nil, ErrNotFound
	}
	return nl, nil
}

func (dc DigitCombinations) Add(value string, combination string) error {
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

func (dc DigitCombinations) Update(value string, combination string) error {

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

func (combo Combinations) Search(value string) (DigitCombinations, error) {

	nl, ok := combo[value]

	if !ok {
		return nil, ErrNotFound
	}
	return nl, nil
}

func (combo Combinations) Add(value string, dc DigitCombinations) error {
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

func intToString(i int) string {
	return strconv.Itoa(i)
}
