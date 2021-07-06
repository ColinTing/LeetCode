# [121. 买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)


## 题目

给定一个数组 `prices` ，它的第 `i` 个元素 `prices[i]` 表示一支给定股票第 `i` 天的价格。

你只能选择 **某一天** 买入这只股票，并选择在 **未来的某一个不同的日子** 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 `0` 。

**Example 1:**

```
输入：[7,1,5,3,6,4]

输出：5

解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
```

**Example 2:**

```
输入：prices = [7,6,4,3,1]

输出：0

解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
```    

## 题目大意

给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。注意你不能在买入股票前卖出股票。

## 解题思路

- 题目要求找出股票中能赚的钱最多的差价
- 经典的动态规划题

## 代码


**Java版**

```java
class Solution {
    public int maxProfit(int[] prices) {

        int res = 0;
        int[] dp = new int[prices.length];

        for (int i = 1; i < prices.length; i++) {
            dp[i] = prices[i] - prices[i - 1];
        }

        for (int i = 1; i < dp.length; i++) {
            if (dp[i - 1] > 0) {
                dp[i] += dp[i - 1];
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

func maxProfit(prices []int) int {

	res, dp := 0, make([]int, len(prices))

	for i := 1; i < len(prices); i++ {
		dp[i] = prices[i] - prices[i-1]
	}

	for i := 1; i < len(dp); i++ {
		if dp[i-1] > 0 {
			dp[i] += dp[i-1]
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