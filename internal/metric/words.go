package metric

import "strings"

func WordsCount(text string) int {
	words := strings.Fields(text)

	return len(words)
}
