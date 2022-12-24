package Internal

import "fmt"

func Step(a int) {
	fmt.Println("Step + 1")
	for i := 1; i <= a; i++ {
		fmt.Println(i)
	}

	fmt.Println("Step - 2")
	for i := a; i >= 1; i -= 2 {
		fmt.Println(i)
	}
}
