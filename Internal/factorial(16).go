package Internal

import "fmt"

func Factorial(num int) {
	for i := 1; i <= num; i++ {
		res := 1
		fmt.Println("\nfac", i)
		for j := 1; j <= i; j++ {
			if j != i {
				fmt.Print(j, "*")
				res *= j
			} else {
				fmt.Print(j)
				res *= j
			}
		}
		fmt.Println(" =", res)
	}
}
