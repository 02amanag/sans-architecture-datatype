package string

import "errors"

func DuplicacyCheck(list []string) error {
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

func DuplicacyRemove(list []string) []string {
	var newlist []string
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
