package leetcode

func isAnagram(s string, t string) bool {

	alphabet := make(map[rune]int)
	for _, c := range s {
		alphabet[c]++
	}

	for _, c := range t {
		alphabet[c]--
		if alphabet[c] < 0 {
			return false
		} else if alphabet[c] == 0 {
			delete(alphabet, c)
		}
	}

	for _, v := range alphabet {
		if v != 0 {
			return false
		}
	}
	return true
}
