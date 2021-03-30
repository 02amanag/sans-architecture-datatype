package int32

import (
	"errors"
)

func Find(value int32, list []int32) error {
	for _, i := range list {
		if i == value {
			return nil
		}
	}
	return errors.New("not found")
}

func IfFindRemove(value int32, list []int32) []int32 {
	var newlist []int32
	for _, i := range list {
		if i != value {
			newlist = append(newlist, i)
		}
	}
	return newlist
}

func IfNotFindAdd(value int32, list []int32) []int32 {
	var newlist []int32
	for _, i := range list {
		if i == value {
			return list
		}
	}
	newlist = append(list, value)
	return newlist
}
