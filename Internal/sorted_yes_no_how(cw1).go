package Internal

import (
	"reflect"
	"sort"
)

func IsSortedAndHow(array []int) string {
	arrAscending := SortAscending(array)
	arrDescending := SortDescending(array)

	if reflect.DeepEqual(arrDescending, array) {
		return "yes, descending"
	} else if reflect.DeepEqual(arrAscending, array) {
		return "yes, ascending"
	} else {
		return "no"
	}
}

func SortAscending(array []int) []int {
	var arr []int
	for i := 0; i < len(array); i++ {
		arr = append(arr, array[i])
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}

func SortDescending(array []int) []int {
	var arr []int
	for i := 0; i < len(array); i++ {
		arr = append(arr, array[i])
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	return arr
}
