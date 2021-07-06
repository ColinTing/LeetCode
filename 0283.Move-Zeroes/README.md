# [283. 移动零](https://leetcode-cn.com/problems/move-zeroes/)

## 题目

给定一个数组 `nums`，编写一个函数将所有 `0` 移动到数组的末尾，同时保持非零元素的相对顺序。

**Example 1:**

```
输入: [0,1,0,3,12]

输出: [1,3,12,0,0]
```

**Note:**

- 必须在原数组上操作，不能拷贝额外的数组。
- 尽量减少操作次数。



## 题目大意

题目要求不能采用额外的辅助空间，将数组中 0 元素都移动到数组的末尾，并且维持所有非 0 元素的相对位置。

## 解题思路

与 `26. 删除排序数组中的重复项` 题类似

- 解法一：双指针法 （不断的用 i，j 标记 0 和非 0 的元素，然后相互交换，最终到达题目的目的） 
- 解法二：零值个数计算法（当前替换位可通过index减去当前累计零值个数获得）

## 代码


**Java版**

```java
// 解法一：双指针法 （不断的用 i，j 标记 0 和非 0 的元素，然后相互交换，最终到达题目的目的） 
class Solution {
    public void moveZeroes(int[] nums) {

        int i = 0;
        for (int j = 1; j < nums.length; j++) {
            if (nums[i] == 0) {
                if (nums[j] != 0) {
                    int temp = nums[i];
                    nums[i++] = nums[j];
                    nums[j] = temp;
                }
            } else {
                i++;
            }
        }
    }
}

// 解法二：零值个数计算法（当前替换位可通过index减去当前累计零值个数获得）
class Solution {
    public void moveZeroes(int[] nums) {

        int zeroCt = 0;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == 0) {
                zeroCt++;
            } else {
                int temp = nums[i - zeroCt];
                nums[i - zeroCt] = nums[i];
                nums[i] = temp;
            }
        }
    }
}
```

**Golang版**

```go
package leetcode

// 解法一：双指针法 （不断的用 i，j 标记 0 和非 0 的元素，然后相互交换，最终到达题目的目的）
func moveZeroes(nums []int) {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] == 0 {
			if i != j {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					i++
				}
			}
		} else {
			i++
		}
	}
}

// 解法二：零值个数计算法（当前替换位可通过index减去当前累计零值个数获得）
func moveZeroes1(nums []int) {
	zeroCt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCt++
		} else {
			nums[i-zeroCt], nums[i] = nums[i], nums[i-zeroCt]
		}
	}
}

```


