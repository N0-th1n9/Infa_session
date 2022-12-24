package Internal

import "fmt"

func StonePaperScissors(h1, h2 string) {
	if (h1 == "1" && h2 == "2") || (h1 == "2" && h2 == "3") || (h1 == "3" && h2 == "1") {
		fmt.Println("Player 1 won")
	} else if h1 == h2 {
		fmt.Println("Nobody won, try again")
	} else {
		fmt.Println("Player 2 won")
	}
}
