package solution

func MostPopularWord(words []string) string {
	wordsCount := make(map[string]int, len(words))
	mostPopWord := ""
	max := 0
	for _, word := range words {
		wordsCount[word]++
		if wordsCount[word] > max {
			max = wordsCount[word]
		}
	}

	for _, word := range words {
		if wordsCount[word] == max {
			mostPopWord = word
			break
		}
	}

	return mostPopWord
}
