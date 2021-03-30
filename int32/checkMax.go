package int32

func CheckMax(list []int32) int32 {
	max := list[0]
	for _, value := range list {
		if value > max {
			max = value
		}
	}
	return max
}
