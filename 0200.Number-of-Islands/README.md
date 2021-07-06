# [200. 岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)


## 题目

给你一个由 `'1'`（陆地）和 `'0'`（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

**Example 1:**

```
输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]

输出：1
```

**Example 2:**

```
输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]

输出：3
```    

**提示：**

- m == grid.length
- n == grid[i].length
- 1 <= m, n <= 300
- grid[i][j] 的值为 '0' 或 '1'

## 题目大意

给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。


## 解题思路

- 要求找出地图中的孤岛。孤岛的含义是四周被海水包围的岛。
- 这一题可以按照第 79 题的思路进行搜索，只要找到为 "1" 的岛以后，从这里开始搜索这周连通的陆地，也都标识上访问过。每次遇到新的 "1" 且没有访问过，就相当于遇到了新的岛屿了。

## 代码

**Java版**

```java
class Solution {
    public int numIslands(char[][] grid) {

        int m = grid.length;
        int n = grid[0].length;
        int[][] visited = new int[m][n];
        int res = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (visited[i][j] == 1 || grid[i][j] == '0') {
                    continue;
                }
                numIslandsDFS(grid, visited, i, j);
                res++;
            }
        }
        return res;
    }

    private void numIslandsDFS(char[][] grid, int[][] visited, int i, int j) {

        if (i < 0 || i >= grid.length || j < 0 || j >= grid[0].length ||
                visited[i][j] == 1 || grid[i][j] == '0') {
            return;
        }
        visited[i][j] = 1;
        numIslandsDFS(grid, visited, i + 1, j);
        numIslandsDFS(grid, visited, i - 1, j);
        numIslandsDFS(grid, visited, i, j + 1);
        numIslandsDFS(grid, visited, i, j - 1);
    }
}
```

**Golang版**

```go
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
```