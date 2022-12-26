package Internal

import (
	"fmt"
)

func Piramid(height int) {
	for i := 1; i <= height; i++ {
		for j := 1; j <= height-i; j++ {
			fmt.Print(" ")
		}
		for x := 0; x != (2*i - 1); x++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
