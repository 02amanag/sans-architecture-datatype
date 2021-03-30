package int64

func CheckMin(list []int64) int64 {
	min := list[0]
	for _, value := range list {
		if value < min {
			min = value
		}
	}
	return min
}
