package Internal

import (
	"fmt"
	"strconv"
)

func FizzBuzz(array []string) {
	arr := make([]int, 0, len(array))
	for _, v := range array {
		i, _ := strconv.Atoi(v)
		arr = append(arr, i)
	}

	for i, n := range arr {
		if n%3 == 0 && n%5 == 0 {
			array[i] = "FizzBuzz"
		} else if n%3 == 0 {
			array[i] = "Fizz"

		} else if n%5 == 0 {
			array[i] = "Buzz"
		}
	}

	fmt.Println(array)
}
