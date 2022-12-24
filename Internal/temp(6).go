package Internal

func ConvertTemp(temperature float64, convertTo string) float64 {
	if convertTo == "C" {
		return (temperature * 1.8) + 32
	} else {
		return (temperature - 32) / (1.8)
	}
}
