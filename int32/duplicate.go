package int32

import "errors"

func DuplicacyCheck(list []int32) error {
	for _, i := range list {
		counter := 0
		for _, j := range list {
			if i == j {
				counter++
			}
			if counter > 1 {
				return errors.New("Data Duplicacy")
			}
		}
	}
	return nil
}

func DuplicacyRemove(list []int32) []int32 {
	var newlist []int32
	for _, i := range list {
		counter := 0
		for _, j := range newlist {
			if i == j {
				counter++
			}
		}
		if counter == 0 {
			newlist = append(newlist, i)
		}
	}
	return newlist
}
