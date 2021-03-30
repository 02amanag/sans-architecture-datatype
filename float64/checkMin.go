package float64

func CheckMin(list []float64) float64 {
	min := list[0]
	for _, value := range list {
		if value < min {
			min = value
		}
	}
	return min
}