package Internal

func Bingo(array []int) bool {
	count := 0
	arr := []int{2, 7, 9, 14, 15}

	array = SortAscending(array)
	for i := 0; i < len(array)-1; i++ {
		if array[i] == array[i+1] {
			array[i] = 0
		}
	}

	for _, n := range array {
		for _, v := range arr {
			if n == v {
				count += 1
			}
		}
	}
	if count == 5 {
		return true
	} else {
		return false
	}
}
