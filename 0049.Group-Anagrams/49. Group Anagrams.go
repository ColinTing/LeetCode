package leetcode

func groupAnagrams(strs []string) [][]string {

	m := make(map[string][]string)
	for _, str := range strs {
		cs := make([]byte, 26)
		for _, c := range str {
			cs[c-'a']++
		}
		if v, ok := m[string(cs)]; ok {
			m[string(cs)] = append(v, str)
		} else {
			m[string(cs)] = []string{str}
		}
	}
	res := make([][]string, len(m))
	i := 0
	for _, strings := range m {
		res[i] = strings
		i++
	}
	return res
}
