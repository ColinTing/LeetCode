package leetcode

func updateBoard(board [][]byte, click []int) [][]byte {
	m, n := len(board), len(board[0])

	if m == 0 || n == 0 {
		return nil
	}

	row, col := click[0], click[1]
	ct := 0
	if board[row][col] == 'M' {
		board[row][col] = 'X'
	} else {
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				x := row + i
				y := col + j
				if x < 0 || x >= m || y < 0 || y >= n ||
					(i == 0 && j == 0) {
					continue
				}
				if board[x][y] == 'M' {
					ct++
				}
			}
		}

		if ct > 0 {
			board[row][col] = byte('0' + ct)
		} else {
			board[row][col] = 'B'
			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {
					x := row + i
					y := col + j
					if x < 0 || x >= m || y < 0 || y >= n ||
						(i == 0 && j == 0) {
						continue
					}
					if board[x][y] == 'E' {
						newClick := []int{x, y}
						updateBoard(board, newClick)
					}
				}
			}
		}
	}
	return board
}
