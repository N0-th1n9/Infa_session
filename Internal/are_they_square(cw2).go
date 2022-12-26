package Internal

import "math"

func Square(array []int) bool {
	var count int
	if len(array) == 0 {
		return false
	}
	for _, v := range array {
		if math.Sqrt(float64(v)) == math.Round(math.Sqrt(float64(v))) {
			count++
		}
	}
	if count == len(array) {
		return true
	} else {
		return false
	}
}
