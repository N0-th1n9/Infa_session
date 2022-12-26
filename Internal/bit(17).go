package Internal

import (
	"fmt"
	"strings"
)

func Bit(arr []string) {
	var resArr []string
	if arr[0] == "0" {
		resArr = append(resArr, "0")
	} else {
		resArr = append(resArr, "1")
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			resArr = append(resArr, "0")
		} else {
			resArr = append(resArr, "1")
		}
	}
	fmt.Println(strings.Join(resArr, ""))
}
