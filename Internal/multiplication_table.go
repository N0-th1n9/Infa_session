package main

import "fmt"

func main() {
	a := 15
	fmt.Println("\nV1")
	multiplicationTableV1(a)
	fmt.Println("\nV2")
	multiplicationTableV2()
}

func multiplicationTableV1(a int) {
	for i := 1; i <= a; i++ {
		fmt.Println(fmt.Sprintf("Умножение на %d", i))
		for j := 1; j <= a; j++ {
			res := i * j
			fmt.Println(fmt.Sprintf("%d * %d = %d", i, j, res))
		}
	}
}

func multiplicationTableV2() {
	for i := 1; i < 21; i++ {
		for j := 1; j < 21; j++ {
			res := i * j
			fmt.Print(res, "\t")
		}
		fmt.Println()
	}
}
