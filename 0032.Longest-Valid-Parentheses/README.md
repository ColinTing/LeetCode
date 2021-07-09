# [32. 最长有效括号](https://leetcode-cn.com/problems/longest-valid-parentheses/)


## 题目

给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

**Example 1:**

```
输入：s = "(()"

输出：2

解释：最长有效括号子串是 "()"
```

**Example 2:**

```
输入：s = ")()())"

输出：4

解释：最长有效括号子串是 "()()"
```

**Example 3:**

```
输入：s = ""

输出：0
```

**提示:**

- `0 <= s.length <= 3 * 104`
- `s[i]` 为 `'('`, 或 `')'`.

## 题目大意

给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

## 解题思路

- dp[i]数组表示在此位置（i），最长合法子串长度是多少
- 从后往前遍历，因为必须保证是‘（’起始，所以，位置定位在字符串长度的倒数第二个位置（n-2）
- si = '('，则需要找到右括号和它匹配，可以跳过以i + 1开头的合法子串。i+1开头的合法子串有多长，有dp[i+1]个长度
- 取这些长度之后的一位，即j = i + dp[i+1] + 1;
- 在j没有越界的情况下，如果j位是')'，dp[i] = dp[i+1] + 2
- 如果j位没有到达末尾，还要加上j+1位的dp值

## 代码

**Java版**

```java
class Solution {
    public int longestValidParentheses(String s) {
        int n = s.length();
        int[] dp = new int[n];
        int res = 0;
        for (int i = n - 2; i >= 0; i--) {
            if (s.charAt(i) == '(') {
                int j = i + dp[i + 1] + 1;
                if (j < n && s.charAt(j) == ')') {
                    dp[i] += dp[i + 1] + 2;
                    if (j + 1 < n) {
                        dp[i] += dp[j + 1];
                    }
                }
            }
            res = Math.max(res, dp[i]);
        }
        return res;
    }
}
```

**Golang版**


```go
package leetcode

func longestValidParentheses(s string) int {
	n := len(s)
	dp := make([]int, n)
	res := 0
	for i := n - 2; i >= 0; i-- {
		if s[i] == '(' {
			j := i + dp[i+1] + 1
			if j < n && s[j] == ')' {
				dp[i] += dp[i+1] + 2
				if j+1 < n {
					dp[i] += dp[j+1]
				}
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
```