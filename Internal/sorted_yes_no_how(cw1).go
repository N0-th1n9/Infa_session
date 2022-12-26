package Internal

import (
	"sort"
)

func IsSortedAndHow(array []int) string {
	arrAscending := SortAscending(array)
	arrDescending := SortDescending(array)
	for j := 0; j < len(array); j++ {
		if arrAscending[j] != array[j] {
			if arrDescending[j] != array[j] {
				return "no"
			}
			return "yes, descending"
		}
	}
	return "yes, ascending"
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
