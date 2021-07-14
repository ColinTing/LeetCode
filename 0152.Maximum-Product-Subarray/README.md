# [152. 乘积最大子数组](https://leetcode-cn.com/problems/maximum-product-subarray/)


## 题目

给你一个整数数组 `nums` ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

**Example 1:**

```
输入: [2,3,-2,4]

输出: 6

解释: 子数组 [2,3] 有最大乘积 6。
```

**Example 2:**
```
输入: [-2,0,-1]

输出: 0

解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
```

## 题目大意

给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。


## 解题思路

- 给出一个数组，要求找出这个数组中连续元素乘积最大的值。
- 最大值存在于
1. 当前值（最大值为负数的情况下，当前值为正数）
2. 最小值*当前值（最小值为负数,当前值为负数）
3. 最大值*当前值（最大值为正数，当前值为正数）
- 因为三个条件都包含了当前值，所以保证了连续性,所需要的就是维护两个dp方程，一个最大值的dp方程、一个最小值的dp方程，最后结果取历史中的最大值

## 代码

**Java版**

```java
class Solution {
    public int maxProduct(int[] nums) {
        int res = nums[0];
        int minNum = nums[0], maxNum = nums[0];

        for (int i = 1; i < nums.length; i++) {
            if (nums[i] < 0) {
                int temp = minNum;
                minNum = maxNum;
                maxNum = temp;
            }
            maxNum = Math.max(nums[i], nums[i] * maxNum);
            minNum = Math.min(nums[i], nums[i] * minNum);
            res = Math.max(res, maxNum);
        }
        return res;
    }
}
```

**Golang版**

```go
package leetcode

func maxProduct(nums []int) int {
	maxNum := nums[0]
	minNum := nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxNum, minNum = minNum, maxNum
		}
		maxNum = max(nums[i], nums[i]*maxNum)
		minNum = min(nums[i], nums[i]*minNum)
		res = max(res, maxNum)
	}
	return res
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