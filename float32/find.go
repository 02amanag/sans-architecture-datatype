package float32

import (
	"errors"
)

func Find(value float32, list []float32) error {
	for _, i := range list {
		if i == value {
			return nil
		}
	}
	return errors.New("not found")
}

func IfFindRemove(value float32, list []float32) []float32 {
	var newlist []float32
	for _, i := range list {
		if i != value {
			newlist = append(newlist, i)
		}
	}
	return newlist
}

func IfNotFindAdd(value float32, list []float32) []float32 {
	var newlist []float32
	for _, i := range list {
		if i == value {
			return list
		}
	}
	newlist = append(list, value)
	return newlist
}
