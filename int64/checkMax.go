package int64

func CheckMax(list []int64) int64 {
	max := list[0]
	for _, value := range list {
		if value > max {
			max = value
		}
	}
	return max
}
