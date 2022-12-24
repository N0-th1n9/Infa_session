package main

import (
	"fmt"
	"math"
)

func main() {
	var b float64 = 2
	var x float64 = 0
	for ; x <= 3*(math.Pi); x += 0.1 * (math.Pi) {
		fmt.Println(funcSin(b, x))
	}
}

func funcSin(b, x float64) float64 {
	return 1 + math.Sin(math.Pow(b, 3)+math.Pow(x, 3))
}
