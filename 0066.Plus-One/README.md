# [66. 加一](https://leetcode-cn.com/problems/plus-one/)


## 题目

给定一个由 **整数** 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

**Example 1:**

```
输入：digits = [1,2,3]

输出：[1,2,4]

解释：输入数组表示数字 123。
```

**Example 2:**

```
输入：digits = [4,3,2,1]

输出：[4,3,2,2]

解释：输入数组表示数字 4321。
```

**Example 3:**

```
输入：digits = [0]

输出：[1]
```


## 题目大意


给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。你可以假设除了整数 0 之外，这个整数不会以零开头。



## 解题思路

- 从数组尾部开始往前扫，逐位进位即可。最高位如果还有进位需要在数组里面第 0 位再插入一个 1 。

## 代码


**Java版**

```java
class Solution {
    public int[] plusOne(int[] digits) {

        int n = digits.length;

        for (int i = n - 1; i >= 0; i--) {
            digits[i]++;
            if (digits[i] != 10) {
                return digits;
            }
            digits[i] = 0;
        }

        int[] newDigits = new int[n + 1];
        newDigits[0] = 1;
        return newDigits;
    }
}
```

**Golang版**

```go
package leetcode

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 {
			return digits
		}
		digits[i] = 0
	}
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}
```