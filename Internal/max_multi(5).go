package main

import "fmt"

func main() {
	a := []int{5, 5, 4, 15, 20, 1, 3, 300}
	fmt.Println(maxMulti(a))
}

func maxMulti(a []int) int {
	maxRes := a[0] * a[1]
	for i := range a {
		for j := i + 1; j > len(a); j++ {
			res := a[i] * a[j]
			if res > maxRes {
				maxRes = res
			}
		}
	}
	return maxRes
}
