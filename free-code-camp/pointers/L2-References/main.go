package main

import (
	"strings"
)

func removeProfanity(message *string) {
	safeMessage := *message
	safeMessage = strings.ReplaceAll(safeMessage, "fubb", "****")
	safeMessage = strings.ReplaceAll(safeMessage, "shiz", "****")
	safeMessage = strings.ReplaceAll(safeMessage, "witch", "*****")

	*message = safeMessage
}
