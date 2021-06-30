# [189. 旋转数组](https://leetcode-cn.com/problems/rotate-array/)

## 题目

给定一个数组，将数组中的元素向右移动 `k` 个位置，其中 `k` 是非负数。

**进阶：**

- 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
- 你可以使用空间复杂度为 O(1) 的 **原地** 算法解决这个问题吗？

**Example 1:**

```
输入: nums = [1,2,3,4,5,6,7], k = 3

输出: [5,6,7,1,2,3,4]

解释:

向右旋转 1 步: [7,1,2,3,4,5,6]

向右旋转 2 步: [6,7,1,2,3,4,5]

向右旋转 3 步: [5,6,7,1,2,3,4]
```

**Example 2:**

```
输入：nums = [-1,-100,3,99], k = 2

输出：[3,99,-1,-100]

解释: 

向右旋转 1 步: [99,-1,-100,3]

向右旋转 2 步: [3,99,-1,-100]
```

**提示：**

- `1 <= nums.length <= 2 * 10^4`
- `-2^31 <= nums[i] <= 2^31 - 1`
- `0 <= k <= 10^5`

## 题目大意

给定一个数组，将数组中的元素向右移动 `k` 个位置，其中 `k` 是非负数。

## 解题思路

- 解法一：三次旋转法
由于题目要求不能使用额外的空间。翻转最终态是，末尾 `k mod n` 个元素移动至了数组头部，剩下的元素右移 `k mod n` 个位置至最尾部。确定了最终态以后再变换就很容易。先将数组中所有元素从头到尾翻转一次，尾部的所有元素都到了头部，然后再将 `[0,(k mod n) − 1]` 区间内的元素翻转一次，最后再将 `[k mod n, n − 1]` 区间内的元素翻转一次，即可满足题目要求。

- 解法二：迭代取模赋值法，缺点是空间复杂度是O(n)

## 代码


**Java版**

```java

class Solution {

    //解法一:三次旋转法, 时间复杂度 O(n)，空间复杂度 O(1)
    public void rotate(int[] nums, int k) {
        int n = nums.length;
        k %= n;
        reverseArray(nums, 0, n - 1);
        reverseArray(nums, 0, k - 1);
        reverseArray(nums, k, n - 1);
    }

    private void reverseArray(int[] nums, int start, int end) {

        while (start < end) {
            int temp = nums[start];
            nums[start++] = nums[end];
            nums[end--] = temp;
        }

    }

    //解法二：迭代取模赋值法, 时间复杂度 O(n)，空间复杂度 O(n)
    public void rotate1(int[] nums, int k) {
        int n = nums.length;
        k %= n;
        int[] newNums = new int[n];
        for (int i = 0; i < nums.length; i++) {
            newNums[(i + k) % n] = nums[i];
        }

        System.arraycopy(newNums, 0, nums, 0, n);
    }    
}

```


**Golang版**

```go
package leetcode

//解法一:三次旋转法, 时间复杂度 O(n)，空间复杂度 O(1)
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}

//解法二：迭代取模赋值法, 时间复杂度 O(n)，空间复杂度 O(n)
func rotate1(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}

	copy(nums, newNums)
}
```