# [164. 最大间距](https://leetcode-cn.com/problems/maximum-gap/)

## 题目

给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。

如果数组元素个数小于 2，则返回 0。

Example 1:

```c
输入: [3,6,9,1]

输出: 3

解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。
```

Example 2:

```c
输入: [10]

输出: 0

解释: 数组元素个数小于 2，因此返回 0。
```


Note:

- 你可以假设数组中所有元素都是非负整数，且数值在 32 位有符号整数范围内。
- 请尝试在线性时间复杂度和空间复杂度的条件下解决此问题。


## 题目大意

在数组中找到 2 个数字之间最大的间隔。要求尽量用 O(1) 的时间复杂度和空间复杂度。

## 解题思路


这道题满足要求的做法是基数排序,但我使用的是快排。


## 代码

**Java版**

```java
class Solution {
    public int maximumGap(int[] nums) {
        int res = 0;
        quickSort(nums, 0, nums.length - 1);
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] - nums[i - 1] > res) {
                res = nums[i] - nums[i - 1];
            }

        }
        return res;
    }

    private void quickSort(int[] nums, int left, int right) {
        if (left >= right) {
            return;
        }
        int p = partition(nums, left, right);
        quickSort(nums, left, p - 1);
        quickSort(nums, p + 1, right);
    }

    private int partition(int[] nums, int left, int right) {

        int pivot = nums[left];
        int i = left;
        for (int j = left + 1; j <= right; j++) {
            if (nums[j] < pivot) {
                i++;
                int temp = nums[i];
                nums[i] = nums[j];
                nums[j] = temp;
            }
        }
        int temp = nums[i];
        nums[i] = nums[left];
        nums[left] = temp;
        return i;
    }
}
```

**Golang版**

```go
package leetcode

func maximumGap(nums []int) int {
	res := 0
	quickSort(nums, 0, len(nums)-1)
	for i := 1; i < len(nums); i++ {
		if res < nums[i]-nums[i-1] {
			res = nums[i] - nums[i-1]
		}
	}
	return res
}

func quickSort(nums []int, start int, end int) {
	if start >= end {
		return
	}
	p := partition(nums, start, end)
	quickSort(nums, start, p-1)
	quickSort(nums, p+1, end)
}

func partition(nums []int, start int, end int) int {

	pivot := nums[start]
	i := start
	for j := i + 1; j <= end; j++ {
		if pivot > nums[j] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[start] = nums[start], nums[i]
	return i
}
```

