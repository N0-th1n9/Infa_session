package main

import "fmt"

func main() {
	fmt.Println("\nV1")
	multiplicationTableV1()
	fmt.Println("\nV2")
	multiplicationTableV2()
}

func multiplicationTableV1() {
	for i := 1; i < 10; i++ {
		fmt.Println(fmt.Sprintf("Умножение на %d", i))
		for j := 1; j < 10; j++ {
			res := i * j
			fmt.Println(fmt.Sprintf("%d * %d = %d", i, j, res))
		}
	}
}

func multiplicationTableV2() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			res := i * j
			fmt.Print(res, "\t")
		}
		fmt.Println()
	}
}
