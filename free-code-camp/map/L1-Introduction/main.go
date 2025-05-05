package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	nameToUserMap := make(map[string]user)
	for i, name := range names {
		nameToUserMap[name] = user{
			name:        name,
			phoneNumber: phoneNumbers[i],
		}
	}
	return nameToUserMap, nil
}

type user struct {
	name        string
	phoneNumber int
}
