package main

import "fmt"

func main() {
	a := []int{5, 5, 4, 15, 20, 1, 3, -1}
	maxMulti(a)
}

func maxMulti(a []int) {
	maxRes := a[0] * a[1]
	for i := range a {
		res := a[i] * a[i+1]
		if res > maxRes {
			maxRes = res
		}
	}
	fmt.Println(maxRes)
}
