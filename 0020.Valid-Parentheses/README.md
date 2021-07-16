# [20. 有效的括号](https://leetcode-cn.com/problems/valid-parentheses/)

## 题目

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

1.左括号必须用相同类型的右括号闭合。
2.左括号必须以正确的顺序闭合。

Example 1:

```
输入: "()"

输出: true
```


Example 2:

```
输入: "()[]{}"

输出: true
```

Example 3:

```
输入：s = "(]"

输出：false
```

Example 4:

```
输入：s = "([)]"

输出：false
```

Example 5:

```
输入：s = "{[]}"

输出：true
```

## 题目大意

括号匹配问题。

## 解题思路


- 与 `32. 最长有效括号` 类似，需要比对三种不同的括号，求出其中的最长值，只不过这里求出最长有效长度后还要比对是不是和整体字符串长度相等，相等则表示是有效括号
- 需要注意，空字符串是满足括号匹配的，即输出 true。

## 代码

**Java版**

```java
class Solution {
    public boolean isValid(String s) {
        int n = s.length();
        int[] dp = new int[n];
        int max = 0;
        for (int i = n - 2; i >= 0; i--) {
            int a = validFunc(s, dp, i, '(', ')');
            int b = validFunc(s, dp, i, '[', ']');
            int c = validFunc(s, dp, i, '{', '}');
            max = Math.max(a, Math.max(b, c));
        }
        return max == s.length();
    }

    private int validFunc(String s, int[] dp, int i, char openChar, char closeChar) {
        if (s.charAt(i) == openChar) {
            int j = i + dp[i + 1] + 1;
            if (j < dp.length && s.charAt(j) == closeChar) {
                dp[i] = dp[i + 1] + 2;
                if (j + 1 < dp.length) {
                    dp[i] += dp[j + 1];
                }
            }
        }
        return dp[i];
    }
}
```

**Golang版**

```go
package leetcode

func isValid(s string) bool {
	n := len(s)
	dp := make([]int, n)
	maxValidLen := 0
	for i := n - 2; i >= 0; i-- {
		a := validFunc(&dp, s, i, '(', ')')
		b := validFunc(&dp, s, i, '[', ']')
		c := validFunc(&dp, s, i, '{', '}')
		maxValidLen = max(a, max(b, c))
	}
	return maxValidLen == len(s)
}

func validFunc(dp *[]int, s string, i int, openByte byte, closeByte byte) int {
	if s[i] == openByte {
		j := i + (*dp)[i+1] + 1
		if j < len(s) && s[j] == closeByte {
			(*dp)[i] = (*dp)[i+1] + 2
			if j+1 < len(s) {
				(*dp)[i] += (*dp)[j+1]
			}
		}
	}
	return (*dp)[i]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
```