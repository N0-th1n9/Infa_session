package Internal

func KmFromMils(mils float64) float64 {
	return mils / 1.609
}

func MilsFromKm(km float64) float64 {
	return km * 1.609
}
