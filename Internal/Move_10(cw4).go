package Internal

import (
	"fmt"
	"strings"
)

func Move(str string) {
	var strArr []string = strings.Split(str, "")

	var newArray []string
	for i := 0; i < len(strArr); i++ {
		newLetter := Switch(strArr[i])
		newArray = append(newArray, newLetter)
	}
	fmt.Println(strings.Join(newArray, ""))
}

func Switch(i string) string {
	alf := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for s := 0; s < len(alf); s++ {
		if alf[s] == i {
			if s+10 < len(alf) {
				return alf[s+10]
			} else {
				return alf[10-(len(alf)-s)]
			}
		}
	}
	return ""
}
