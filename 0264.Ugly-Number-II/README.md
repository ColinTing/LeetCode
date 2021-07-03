# [264. 丑数 II](https://leetcode-cn.com/problems/ugly-number-ii/)


## 题目

给你一个整数 `n` ，请你找出并返回第 `n` 个 丑数 。

**丑数** 就是只包含质因数 2、3、5 的正整数。

**Example 1:**

```
输入：n = 10

输出：12

解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
```

**Example 2:**

```
输入：n = 1

输出：1

解释：1 通常被视为丑数。
```

**提示：**

- `1 <= n <= 1690`

## 题目大意

给你一个整数 `n` ，请你找出并返回第 `n` 个 **丑数** 。**丑数** 就是只包含质因数 `2`、`3`、 `5` 的正整数。

## 解题思路

- 设置三个变量，分别从零开始索引丑数数组，分别*2、*3、*5，如果当前等式值最小，取丑数值的索引位+1 ，公式表达为：Min(L1 * 2, L2 * 3, L3 * 5)。注意：如果三个等式中有同时是最小值的情况，索引位都要加一

## 代码

**Java版**

```java
class Solution {
    public int nthUglyNumber(int n) {
        int[] uglys = new int[n];

        int i = 0, j = 0, k = 0, idx = 0;
        uglys[idx++] = 1;
        while (idx < n) {
            int min = Math.min((Math.min(uglys[i] * 2, uglys[j] * 3)), uglys[k] * 5);
            uglys[idx++] = min;
            if (min == uglys[i] * 2) {
                i++;
            }
            if (min == uglys[j] * 3) {
                j++;
            }
            if (min == uglys[k] * 5) {
                k++;
            }

        }
        return uglys[n - 1];
    }
}
```

**Golang版**

```go
package leetcode

func nthUglyNumber(n int) int {
	dp, p2, p3, p5 := make([]int, n), 0, 0, 0
	dp[0] = 1
	for i := 1; i < n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(x2, x3), x5)
		if x2 == dp[i] {
			p2++
		}

		if x3 == dp[i] {
			p3++
		}

		if x5 == dp[i] {
			p5++
		}
	}
	return dp[n-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
```