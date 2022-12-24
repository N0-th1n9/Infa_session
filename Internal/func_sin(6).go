package Internal

import (
	"math"
)

func FuncSin(b, x float64) float64 {
	return 1 + math.Sin(math.Pow(b, 3)+math.Pow(x, 3))
}
