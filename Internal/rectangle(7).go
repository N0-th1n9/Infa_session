package Internal

import "fmt"

func Rectangle(a, b int) {
	for i := 1; i <= a; i++ {
		for j := 1; j <= b; j++ {
			if j == 1 || j == b {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			if (i == 1 || i == a) && j != b {
				fmt.Print("-")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
