package leetcode

func combine(n int, k int) [][]int {
	res, c :=  [][]int{}, []int{}

	combineDFS(&res, c, 1, n, k)
	return res
}

func combineDFS(res *[][]int, c []int, start int, n int, k int) {
	if k == 0 {
		temp := make([]int, len(c))
		copy(temp, c)
		*res = append(*res, temp)
		return
	}

	for i := start; i <= n; i++ {
		c = append(c, i)
		combineDFS(res, c, i + 1, n, k - 1)
		c = c[:len(c) - 1]
	}
	return
 }
