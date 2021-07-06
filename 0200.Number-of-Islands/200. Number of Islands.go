package leetcode

func numIslands(grid [][]byte) int {
	res, m, n := 0, len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] == true || grid[i][j] == '0' {
				continue
			}

			numIslandsDFS(grid, &visited, i, j)
			res++
		}
	}
	return res
}

func numIslandsDFS(grid [][]byte, visited *[][]bool, i int, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) ||
		(*visited)[i][j] == true || grid[i][j] == '0' {
		return
	}

	(*visited)[i][j] = true
	numIslandsDFS(grid, visited, i+1, j)
	numIslandsDFS(grid, visited, i, j+1)
	numIslandsDFS(grid, visited, i-1, j)
	numIslandsDFS(grid, visited, i, j-1)
}
