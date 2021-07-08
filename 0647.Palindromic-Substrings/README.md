# [647. 回文子串](https://leetcode-cn.com/problems/palindromic-substrings/)


## 题目

给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

**Example 1:**

```
输入："abc"

输出：3

解释：三个回文子串: "a", "b", "c"
```

**Example 2:**

```
输入："aaa"

输出：6

解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
```

**Note:**

- 输入的字符串长度不会超过 1000 。

## 题目大意

给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

## 解题思路

- 将 dp[i][j] 定义成子字符串在 [i, j] 区间是否是回文串的bool变量，然后i从 n-1 往0遍历，j从i往 n-1 遍历，然后看 s[i] 和 s[j] 是否相等，这时候需要留意一下，有了 s[i] 和 s[j] 相等这个条件后，i和j的位置关系很重要，如果i和j相等了，则 dp[i][j] 肯定是 true；如果i和j是相邻的，那么 dp[i][j] 也是 true；如果i和j中间只有一个字符，那么 dp[i][j] 还是 true；如果中间有多余一个字符存在，即中间有两个字符（j - i > 2的情况），则需要看 dp[i+1][j-1] 是否为 true，若为 true，那么 dp[i][j] 就是 true。赋值 dp[i][j] 后，如果其为 true，结果 res 自增1

## 代码

**Java版**

```java
class Solution {
    public int countSubstrings(String s) {
        int n = s.length();
        boolean[][] dp = new boolean[n][n];
        int res = 0;
        for (int i = n - 1; i >= 0; i--) {
            for (int j = i; j < n; j++) {
                if (s.charAt(i) == s.charAt(j)) {
                    if (j - i <= 2 || dp[i + 1][j - 1]) {
                        dp[i][j] = true;
                        res++;
                    }
                }
            }
        }
        return res;
    }
}
```



**Golang版**

```go
package leetcode

func countSubstrings(s string) int {
	res := 0
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 2 || dp[i+1][j-1] {
					res++
					dp[i][j] = true
				}
			}
		}
	}
	return res
}
```