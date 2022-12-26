package Internal

import (
	"math"
)

func FuncSqrs(a, b, x, n float64) []float64 {
	var resArr []float64
	for x = 0; x <= 30; x += n {
		var res float64 = math.Pow(x, a*1.0/3.0) + math.Pow(b, 1.0/3.0)
		resArr = append(resArr, res)
	}
	return resArr
}
