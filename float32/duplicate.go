package float32

import "errors"

func DuplicacyCheck(list []float32) error {
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

func DuplicacyRemove(list []float32) []float32 {
	var newlist []float32
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
