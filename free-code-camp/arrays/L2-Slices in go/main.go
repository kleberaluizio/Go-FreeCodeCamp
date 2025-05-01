package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	switch plan {
	case planPro:
		return messages[:], nil
	case planFree:
		return messages[0:2], nil
	default:
		return nil, errors.New("unsupported plan")
	}
}

//A slice is a dynamically-sized, flexible view of the elements of an array.
//arrayname[lowIndex:highIndex]
//arrayname[lowIndex:]
//arrayname[:highIndex]
//arrayname[:]
