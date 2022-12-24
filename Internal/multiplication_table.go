package Internal

import "fmt"

func MultiplicationTableV1(a int) {
	for i := 1; i <= a; i++ {
		fmt.Println(fmt.Sprintf("Умножение на %d", i))
		for j := 1; j <= a; j++ {
			res := i * j
			fmt.Println(fmt.Sprintf("%d * %d = %d", i, j, res))
		}
	}
}

func MultiplicationTableV2(a int) {
	for i := 1; i < a; i++ {
		for j := 1; j < a; j++ {
			res := i * j
			fmt.Print(res, "\t")
		}
		fmt.Println()
	}
}
