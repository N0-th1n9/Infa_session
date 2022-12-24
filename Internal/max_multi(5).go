package Internal

func MaxMulti(a []int) int {
	maxRes := a[0] * a[1]
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			res := a[i] * a[j]
			if res > maxRes {
				maxRes = res
			}
		}
	}
	return maxRes
}
