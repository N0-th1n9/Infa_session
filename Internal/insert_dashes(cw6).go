package Internal

import (
	"fmt"
	"strconv"
	"strings"
)

func Dashes(num []int) {
	var arr []string
	for i := 0; i < len(num)-1; i++ {
		arr = append(arr, strconv.Itoa(num[i]))
		if num[i+1]%2 != 0 && num[i]%2 != 0 {
			arr = append(arr, "-")
		}
		if i == len(num)-2 {
			arr = append(arr, strconv.Itoa(num[i+1]))
		}
	}
	fmt.Println(strings.Join(arr, ""))
}
