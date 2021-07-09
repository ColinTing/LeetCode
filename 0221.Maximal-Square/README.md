# [221. 最大正方形](https://leetcode-cn.com/problems/maximal-square/)


## 题目

在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/11/26/max1grid.jpg)

```
输入：matrix = [
  ["1","0","1","0","0"],
  ["1","0","1","1","1"],
  ["1","1","1","1","1"],
  ["1","0","0","1","0"]
]

输出：4
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2020/11/26/max2grid.jpg)

```
输入：matrix = [
  ["0","1"],
  ["1","0"]
]

输出：1
```    

**提示：**

- m == matrix.length
- n == matrix[i].length
- 1 <= m, n <= 300
- matrix[i][j] 为 '0' 或 '1'

## 题目大意

在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。


## 解题思路

- 正方形的形成条件需要两边都相等，可转化成上、左、左上位置的边长都准备好了，右下角的边长就可以边长+1了，这时右下角的边长就是其中边长最小值+1。

## 代码

**Java版**

```java
class Solution {
    public int maximalSquare(char[][] matrix) {
        int m = matrix.length;
        int n = matrix[0].length;
        int[][] dp = new int[m][n];
        int maxSideLen = 0;
        for (int i = 0; i < m; i++) {
            if (matrix[i][0] == '1') {
                dp[i][0] = 1;
            }
            maxSideLen = Math.max(maxSideLen, dp[i][0]);
        }
        for (int j = 0; j < n; j++) {
            if (matrix[0][j] == '1') {
                dp[0][j] = 1;
            }
            maxSideLen = Math.max(maxSideLen, dp[0][j]);
        }
        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                if (matrix[i][j] == '1') {
                    dp[i][j] = Math.min(Math.min(dp[i - 1][j], dp[i][j - 1]), dp[i - 1][j - 1]) + 1;
                }
                maxSideLen = Math.max(maxSideLen, dp[i][j]);
            }
        }
        return maxSideLen * maxSideLen;
    }
}
```

**Golang版**

```go
package leetcode

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	maxSideLen := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
		}
		maxSideLen = max(maxSideLen, dp[i][0])
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
		}
		maxSideLen = max(maxSideLen, dp[0][j])
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
			maxSideLen = max(maxSideLen, dp[i][j])
		}
	}
	return maxSideLen * maxSideLen
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
```