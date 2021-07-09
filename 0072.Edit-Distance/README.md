# [72. 编辑距离](https://leetcode-cn.com/problems/edit-distance/)


## 题目

给你两个单词 `word1` 和 `word2`，请你计算出将 `word1` 转换成 `word2` 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

- 插入一个字符
- 删除一个字符
- 替换一个字符

**Example 1:**

```
输入：word1 = "horse", word2 = "ros"

输出：3

解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
```

**Example 2:**

```
输入：word1 = "intention", word2 = "execution"

输出：5

解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
```

**提示：**

- 0 <= word1.length, word2.length <= 500
- word1 和 word2 由小写英文字母组成

## 题目大意


已知两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。每次操作可选择插入或者删除或者替换



## 解题思路

- dp[i][j] 表示从 word1 的前i个字符转换到 word2 的前j个字符所需要的步骤
- word1[i] 和 word2[j] 相等时，dp[i][j] = dp[i - 1][j - 1]，因为字符相同，所以不用转换和dp[i-1][j-1]处相同
- 不等时,结果是下面三种操作中的最小值+1
1. 左边dp[i][j-1]对应增加操作
2. 上边dp[i-1][j]对应删除操作
3. 左上dp[i-1][j-1]对应修改操作


## 代码


**Java版**

```java
class Solution {
    public int minDistance(String word1, String word2) {
        int m = word1.length();
        int n = word2.length();
        int[][] dp = new int[m + 1][n + 1];
        dp[0][0] = 0;
        for (int i = 1; i <= m; i++) {
            dp[i][0] = i;
        }
        for (int j = 1; j <= n; j++) {
            dp[0][j] = j;
        }

        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                if (word1.charAt(i - 1) == word2.charAt(j - 1)) {
                    dp[i][j] = dp[i - 1][j - 1];
                } else {
                    dp[i][j] = 1 +
                            Math.min(Math.min(dp[i][j - 1], dp[i - 1][j]), dp[i - 1][j - 1]);
                }
            }
        }
        return dp[m][n];
    }
}
```

**Golang版**

```go
package leetcode

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1])
			}
		}
	}
	return dp[m][n]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
```