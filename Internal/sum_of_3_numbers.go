package Internal

import "fmt"

// {1, 2, 3, 4, 5, 6, 100, 54, 32, 2}
func SumOfNumbers(arr []int) {
	arrRes := []int{}
	for i := 2; i < len(arr); i++ {
		res := arr[i] + arr[i-1] + arr[i-2]
		if res < 100 {
			arrRes = append(arrRes, res)
		} else {
			break
		}
	}
	fmt.Print(arrRes, " ")
}
