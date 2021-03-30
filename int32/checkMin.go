package int32

func CheckMin(list []int32) int32 {
	min := list[0]
	for _, value := range list {
		if value < min {
			min = value
		}
	}
	return min
}
