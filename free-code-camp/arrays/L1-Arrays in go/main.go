package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	var accumulatedCost [3]int
	for i := 0; i < len(messages); i++ {
		if i == 0 {
			accumulatedCost[i] = len(messages[i])
			continue
		}
		accumulatedCost[i] = accumulatedCost[i-1] + len(messages[i])
	}

	return messages, accumulatedCost
}
