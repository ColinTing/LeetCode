package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	j := 0
	for i := 0; i < len(s); i++ {
		if j >= len(g) {
			break
		}
		if s[i] >= g[j] {
			j++
		}
	}
	return j
}
