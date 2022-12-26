package Internal

func Average(arr []float64, lenArr int) float64 {
	var sum float64
	for i := 0; i < lenArr; i++ {
		sum += arr[i]
	}
	return sum / float64(lenArr)
}
