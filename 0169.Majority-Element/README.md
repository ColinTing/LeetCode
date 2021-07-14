# [169. 多数元素](https://leetcode-cn.com/problems/majority-element/)


## 题目

给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 `⌊ n/2 ⌋` 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

**Example 1:**

```
输入：[3,2,3]

输出：3
```

**Example 2:**

```
输入：[2,2,1,1,1,2,2]

输出：2
```

## 题目大意


给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。你可以假设数组是非空的，并且给定的数组总是存在众数。


## 解题思路

- 题目要求找出数组中出现次数大于 `⌊ n/2 ⌋` 次的数。要求空间复杂度为 O(1)。简单题。
- 这一题利用的算法是 Boyer-Moore Majority Vote Algorithm(多数投票算法)。[如何理解摩尔投票算法？](https://www.zhihu.com/question/49973163/answer/235921864)
- 简单来说，迭代选择不一样的数进行抵消，最后剩下的那个数就是出现次数大于总数一半的那个
- 一开始我却想到的是快速排序，然后取中间位置值。。。
## 代码

**Java版**

```java
class Solution {
    public int majorityElement(int[] nums) {
        int res = nums[0];
        int count = 1;
        for (int i = 1; i < nums.length; i++) {
            if (count == 0) {
                res = nums[i];
                count++;
            } else if (res == nums[i]) {
                count++;
            } else {
                count--;
            }
        }
        return res;
    }
}
```

**Golang版**

```go
package leetcode

func majorityElement(nums []int) int {
	res, count := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			res = nums[i]
			count++
		} else if nums[i] == res {
			count++
		} else {
			count--
		}
	}
	return res
}
```