# [33. 搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)

## 题目

整数数组 `nums` 按升序排列，数组中的值 互不相同 。

在传递给函数之前，`nums` 在预先未知的某个下标 `k`（`0 <= k < nums.length`）上进行了 **旋转**，使数组变为 `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]`（下标 **从 0 开始** 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 **旋转后** 的数组 `nums` 和一个整数 `target` ，如果 `nums` 中存在这个目标值 `target` ，则返回它的下标，否则返回 `-1` 。

**Example 1:**

```
输入：nums = [4,5,6,7,0,1,2], target = 0

输出：4
```

**Example 2:**

```
输入：nums = [4,5,6,7,0,1,2], target = 3

输出：-1
```

**Example 3:**

```
输入：nums = [1], target = 0

输出：-1
```

## 题目大意

假设按照升序排序的数组在预先未知的某个点上进行了旋转。( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。


## 解题思路

- 给出一个数组，数组中本来是从小到大排列的，并且数组中没有重复数字。但是现在把后面随机一段有序的放到数组前面，这样形成了前后两端有序的子序列。在这样的一个数组里面查找一个数，设计一个 O(log n) 的算法。如果找到就输出数组的小标，如果没有找到，就输出 -1 。
- 由于数组基本有序，虽然中间有一个“断开点”，还是可以使用二分搜索的算法来实现。现在数组前面一段是数值比较大的数，后面一段是数值偏小的数。如果 mid 落在了前一段数值比较大的区间内了，那么一定有 `nums[mid] > nums[low]`，如果是落在后面一段数值比较小的区间内，`nums[mid] ≤ nums[low]` 。如果 mid 落在了后一段数值比较小的区间内了，那么一定有 `nums[mid] < nums[high]`，如果是落在前面一段数值比较大的区间内，`nums[mid] ≤ nums[high]` 。还有 `nums[low] == nums[mid]` 和 `nums[high] == nums[mid]` 的情况，单独处理即可。最后找到则输出 mid，没有找到则输出 -1 。


## 代码

**Java版**

```java
class Solution {
    public int search(int[] nums, int target) {
        int left = 0;
        int right = nums.length - 1;
        while (left <= right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] == target) {
                return mid;
            }
            if (nums[mid] < nums[right]) {
                if (target > nums[mid] && target <= nums[right]) {
                    left = mid + 1;
                } else {
                    right = mid - 1;
                }
            } else {
                if (target >= nums[left] && target < nums[mid]) {
                    right = mid - 1;
                } else {
                    left = mid + 1;
                }
            }
        }
        return -1;
    }
}
```


**Golang版**

```go
package leetcode

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < nums[right] {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}
```