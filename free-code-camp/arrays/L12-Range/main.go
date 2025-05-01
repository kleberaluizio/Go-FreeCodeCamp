package main

func indexOfFirstBadWord(msg []string, badWords []string) int {

	for i, m := range msg {
		for _, badWord := range badWords {
			if m == badWord {
				return i
			}
		}
	}
	return -1
}
