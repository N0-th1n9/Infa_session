package Internal

import "fmt"

func MaxCount(a []int) {
	mCount := a[0]
	for _, l := range a { // _ - индекс(игнорирует, затычка) l - элемент
		if mCount < l {
			mCount = l
		}
	}
	fmt.Println(mCount)
}
