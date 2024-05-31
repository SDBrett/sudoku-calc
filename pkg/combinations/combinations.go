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
type ValueCombinations map[int]NumberList
type DataSet map[int]ValueCombinations
type ValidateCombinations map[string]bool

type DataSetQuery struct {
	NumberOfDigits   int        //The number of digits
	Value            int        //The Value for the digits to sum to
	NumbersToExclude NumberList //Numbers which cannot be part of the solution
	NumbersToInclude NumberList //Numbers which must be part of the solution
}

// Generate the DataSet of all possible combinations
func GenerateDataSet() DataSet {
	combinationsByValue := getCombinationsForValue()

	// Initialize dataSet
	dataSet := newDataSet()

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
func (dc ValueCombinations) Search(value int) (NumberList, error) {

	nl, ok := dc[value]

	if !ok {
		return nil, ErrNotFound
	}
	return nl, nil
}

// Update the combinations to a list of combinations for a a given value
func (dc ValueCombinations) Update(value int, combination string) error {

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

// Returns possible values and their combinations with a given
// number of digits.
func (dc DataSet) Search(value int) (ValueCombinations, error) {

	nl, ok := dc[value]

	if !ok {
		return nil, ErrNotFound
	}
	return nl, nil
}

func (dc DataSet) Query(dsq DataSetQuery) (NumberList, error) {

	nl := NumberList{}

	err := dsq.Validate()
	if err != nil {
		return nil, err
	}

	//combinations := dc[dsq.NumberOfDigits][dsq.Value]
	// valid := ValidateCombinations{}

	// for _, item := range combinations{
	// 	valid[item] = true
	// }

	// for k,_ := range valid{
	// 	for _, item := range dsq.NumbersToExclude{
	// 		if strings.Contains(k,item){
	// 			valid[k] = false
	// 		}
	// 	}
	// 	if valid[k] == true {
	// 	for _, item := range dsq.NumbersToInclude{
	// 		if !strings.Contains(k,item){
	// 			valid[k] = false
	// 		}
	// 	}
	// 	if valid[k] == true {
	// 		nl = append(nl, valid[k])
	// 		}
	// }

	return nl, nil
}

func GetValidCombinations(combinations, exclude, include NumberList) NumberList {

	valid := ValidateCombinations{}
	nl := NumberList{}

	for _, item := range combinations {
		valid[item] = true
	}

	// All combinations start as valid and then checked to determine if invalid
	for k := range valid {
		// Invalidate if combination contains an excluded number
		for _, item := range exclude {
			if strings.Contains(k, item) {
				valid[k] = false
			}
		}
		// Invalidate if combination does not contain an included number
		if valid[k] {
			for _, item := range include {
				if !strings.Contains(k, item) {
					valid[k] = false
				}
			}
		}
		// if still valid add to the returned number list
		if valid[k] {
			nl = append(nl, k)
		}
	}
	return nl
}

func (dsq DataSetQuery) Validate() error {

	var err error
	var x int

	err = utils.ValidateNumberRange(1, 9, dsq.NumberOfDigits)
	if err != nil {
		return err
	}

	err = utils.ValidateNumberRange(3, 45, dsq.Value)
	if err != nil {
		return err
	}

	for _, v := range dsq.NumbersToInclude {
		x, err = utils.StringToInt(v)
		if err != nil {
			return err
		}
		err = utils.ValidateNumberRange(1, 9, x)
		if err != nil {
			return err
		}
	}

	for _, v := range dsq.NumbersToExclude {
		x, err = utils.StringToInt(v)
		if err != nil {
			return err
		}
		err = utils.ValidateNumberRange(1, 9, x)
		if err != nil {
			return err
		}
	}

	return nil
}

func newDataSet() DataSet {
	dataSet := DataSet{
		2: ValueCombinations{},
		3: ValueCombinations{},
		4: ValueCombinations{},
		5: ValueCombinations{},
		6: ValueCombinations{},
		7: ValueCombinations{},
		8: ValueCombinations{},
		9: ValueCombinations{},
	}

	return dataSet
}
