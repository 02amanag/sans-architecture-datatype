package float32

func CheckMin(list []float32) float32 {
	min := list[0]
	for _, value := range list {
		if value < min {
			min = value
		}
	}
	return min
}
