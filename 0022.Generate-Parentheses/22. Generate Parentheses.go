package leetcode

func generateParenthesis(n int) []string {

	res := []string{}
	generateParenthesisDFS(&res, "", n, n)
	return res
}

func generateParenthesisDFS(res *[]string, s string, open int, closed int) {
	if open > closed {
		return
	}

	if open == 0 && closed == 0 {
		*res = append(*res, s)
	} else {
		if open > 0 {
			generateParenthesisDFS(res, s+"(", open-1, closed)
		}

		if closed > 0 {
			generateParenthesisDFS(res, s+")", open, closed-1)
		}
	}
}
