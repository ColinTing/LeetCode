# [88. 合并两个有序数组](https://leetcode-cn.com/problems/merge-sorted-array/)

## 题目

给你两个有序整数数组 `nums1` 和 `nums2`，请你将 `nums2` 合并到 `nums1` 中，使 `nums1` 成为一个有序数组。

Note:  

初始化 `nums1` 和 `nums2` 的元素数量分别为 `m` 和 `n` 。你可以假设 `nums1` 的空间大小等于 `m + n`，这样它就有足够的空间保存来自 `nums2` 的元素。

**Example 1:**

```
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3

输出：[1,2,2,3,5,6]
```

**Example 2:**
```
输入：nums1 = [1], m = 1, nums2 = [], n = 0

输出：[1]
```

## 题目大意

合并两个已经有序的数组，结果放在第一个数组中，第一个数组假设空间足够大。要求算法时间复杂度足够低。

## 解题思路

为了不大量移动元素，就要从2个数组长度之和的最后一个位置开始，依次选取两个数组中大的数，从第一个数组的尾巴开始往头放，只要循环一次以后，就生成了合并以后的数组了。

## 代码

**Java版**

```java
class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int mergeNumsIdx = m + n - 1;
        int nums1Idx = m - 1;
        int nums2Idx = n - 1;
        while (nums1Idx >= 0 && nums2Idx >= 0) {
            if (nums1[nums1Idx] > nums2[nums2Idx]) {
                nums1[mergeNumsIdx--] = nums1[nums1Idx--];
            } else {
                nums1[mergeNumsIdx--] = nums2[nums2Idx--];
            }
        }
        while (nums2Idx >= 0) {
            nums1[mergeNumsIdx--] = nums2[nums2Idx--];
        }
    }
}
```

**Golang版**

```go
package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	for p := m + n; m > 0 && n > 0; p-- {
		if nums1[m-1] > nums2[n-1] {
			nums1[p-1] = nums1[m-1]
			m--
		} else {
			nums1[p-1] = nums2[n-1]
			n--
		}
	}

	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}
```