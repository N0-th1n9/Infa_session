package Internal

import "fmt"

func Matrix(d, m int) {
	var arr [6][5]int
	n := 1
	// Нумеруем все элементы массива
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			arr[i][j] = i
			n++
		}
	}
	// Выводим значения
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], "\t")
		}
		fmt.Println()
	}
}
