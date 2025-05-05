package main

func getNameCounts(names []string) map[rune]map[string]int {
	nameCountMap := make(map[rune]map[string]int)

	for _, name := range names {
		if name == "" {
			continue
		}
		firstLetter := []rune(name)[0]

		if _, exists := nameCountMap[firstLetter]; !exists {
			nameCountMap[firstLetter] = make(map[string]int)
		}

		nameCountMap[firstLetter][name]++
	}

	return nameCountMap
}
