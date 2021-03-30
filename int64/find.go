package int64

import (
	"errors"
)

func Find(value int64, list []int64) error {
	for _, i := range list {
		if i == value {
			return nil
		}
	}
	return errors.New("not found")
}

func IfFindRemove(value int64, list []int64) []int64 {
	var newlist []int64
	for _, i := range list {
		if i != value {
			newlist = append(newlist, i)
		}
	}
	return newlist
}

func IfNotFindAdd(value int64, list []int64) []int64 {
	var newlist []int64
	for _, i := range list {
		if i == value {
			return list
		}
	}
	newlist = append(list, value)
	return newlist
}
