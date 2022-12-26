package Internal

import (
	"fmt"
	"strconv"
)

func FizzBuzz(n int) {
	array := []int{}
	for i := 1; i <= n; i++ {
		array = append(array, i)
	}

	arr := make([]string, 0, len(array))
	for _, v := range array {
		i := strconv.Itoa(v)
		arr = append(arr, i)
	}

	for i, n := range array {
		if n%3 == 0 && n%5 == 0 {
			arr[i] = "FizzBuzz"
		} else if n%3 == 0 {
			arr[i] = "Fizz"

		} else if n%5 == 0 {
			arr[i] = "Buzz"
		}
	}

	fmt.Println(arr)
}
