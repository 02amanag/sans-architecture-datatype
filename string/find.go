package string

import (
	"errors"
)

func Find(value string, list []string) error {
	for _, i := range list {
		if i == value {
			return nil
		}
	}
	return errors.New("not found")
}

func IfFindRemove(value string, list []string) []string {
	var newlist []string
	for _, i := range list {
		if i != value {
			newlist = append(newlist, i)
		}
	}
	return newlist
}

func IfNotFindAdd(value string, list []string) []string {
	var newlist []string
	for _, i := range list {
		if i == value {
			return list
		}
	}
	newlist = append(list, value)
	return newlist
}
