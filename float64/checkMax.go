package float64

func CheckMax(list []float64) float64 {
	max := list[0]
	for _, value := range list {
		if value > max {
			max = value
		}
	}
	return max
}
