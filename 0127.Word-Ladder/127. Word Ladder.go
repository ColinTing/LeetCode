package leetcode

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap, queue, res := buildWordMap(wordList), []string{beginWord}, 0

	for l := len(queue); len(queue) != 0; l = len(queue) {
		for i := 0; i < l; i++ {
			word := queue[i]
			if word == endWord {
				return res + 1
			}
			for j := 0; j < len(word); j++ {
				cw := []byte(word)
				for c := 'a'; c <= 'z'; c++ {
					cw[j] = byte(c)
					newWord := string(cw)
					if newWord == word {
						continue
					}
					if _, ok := wordMap[newWord]; ok {
						delete(wordMap, newWord)
						queue = append(queue, newWord)
					}
				}

			}
		}
		queue = queue[l:]
		res++
	}
	return 0
}

func buildWordMap(wordList []string) map[string]int {
	wordMap := make(map[string]int, len(wordList))
	for i, word := range wordList {
		wordMap[word] = i
	}
	return wordMap
}
