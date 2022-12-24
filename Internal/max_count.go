package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, -5, 10000, 4}
	maxCount(a)
}

func maxCount(a []int) {
	mCount := a[0]
	for _, l := range a { // _ - индекс(игнорирует, затычка) l - элемент
		if mCount < l {
			mCount = l
		}
	}
	fmt.Println(mCount)
}
