package float32

func CheckMax(list []float32) float32 {
	max := list[0]
	for _, value := range list {
		if value > max {
			max = value
		}
	}
	return max
}
