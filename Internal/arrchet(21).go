package Internal

import (
	"fmt"
	"math/rand"
	"time"
)

func ArrChet(n int) {
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, n)
	resArr := []int{}
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(10)
		if arr[i]%2 != 0 {
			resArr = append(resArr, arr[i])
		}
	}
	fmt.Println(arr)
	fmt.Println(resArr)
}
