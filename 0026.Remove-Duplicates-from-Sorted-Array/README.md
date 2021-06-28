# [26. 删除有序数组中的重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)

## 题目

给你一个有序数组 **nums** ，请你 **原地** 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 **原地** 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

Example 1:

```
输入：nums = [1,1,2]

输出：2, nums = [1,2]

解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
```

Example 2:

```
输入：nums = [0,0,1,1,1,2,2,3,3,4]

输出：5, nums = [0,1,2,3,4]

解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
```

说明:

为什么返回数值是整数，但输出的答案是数组呢?

请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下:

```
// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
```

## 题目大意

给定一个有序数组 nums，对数组中的元素进行去重，使得原数组中的每个元素只有一个。最后返回去重以后数组的长度值。

## 解题思路


这里数组的删除并不是真的删除，只是将删除的元素移动到数组后面的空间内，然后返回数组实际剩余的元素个数，

可有两种解法，1.计数法  2.双指针法

## 代码
解法一：双指针法

Java版

```java
    /**
     *
     * i: 重复元素的起始位置
     * j: 当前数组的迭代位置
     * 如果i、j位置元素不重复，i往后挪动一位，并把j位元素放置到i位
     * i从0开始计数，所以返回i+1个元素
     */
    class Solution {
        public int removeDuplicates(int[] nums) {
            int i = 0;
            for (int j = 1; j < nums.length; j++) {
                if (nums[i] != nums[j]) {
                    nums[++i] = nums[j];
                }
            }
            return i + 1;
        }
    }

```

Golang版


```go
package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last := 0
	for finder := 1; finder < len(nums); finder++ {
		if nums[finder] != nums[last] {
			last++
			nums[last] = nums[finder]
		}
	}
	return last + 1
}
```