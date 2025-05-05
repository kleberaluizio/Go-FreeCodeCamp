package main

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, userName := range messagedUsers {
		if _, ok := validUsers[userName]; ok {
			validUsers[userName] += 1
		}
	}
}
