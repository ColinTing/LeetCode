# [242. 有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/)

## 题目

给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

**Example 1**:

```
输入: s = "anagram", t = "nagaram"

输出: true
```

**Example 2:**

```
输入: s = "rat", t = "car"

输出: false
```

说明: 

  
你可以假设字符串只包含小写字母。
  
  
    
进阶: 


如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

## 题目大意

给出 2 个字符串 s 和 t，如果 t 中的字母在 s 中都存在，输出 true，否则输出 false。

## 解题思路

这道题可以用打表的方式做。先把 s 中的每个字母都存在一个 map 中, key 为可能出现的unicode/小写字符。s 中每个字母都对应表中一个字母，每出现一次就加 1。然后再扫字符串 t，每出现一个字母就在表里面减一。如果都出现了，最终表里面的值肯定都是 0 。最终判断表里面的值是否都是 0 即可，有非 0 的数都输出 false 。


## 代码

**Java版**

```java
class Solution {
    public boolean isAnagram(String s, String t) {

        int[] alphabet = new int[26];

        for (char c : s.toCharArray()) {
            ++alphabet[c - 'a'];
        }

        for (char c : t.toCharArray()) {
            --alphabet[c - 'a'];
        }

        for (int i : alphabet) {
            if (i != 0) {
                return false;
            }
        }
        return true;
    }
}
```

**Golang版**

```go
package leetcode

func isAnagram(s string, t string) bool {

	alphabet := make(map[rune]int)
	for _, c := range s {
		alphabet[c]++
	}

	for _, c := range t {
		alphabet[c]--
		if alphabet[c] < 0 {
			return false
		} else if alphabet[c] == 0 {
			delete(alphabet, c)
		}
	}

	for _, v := range alphabet {
		if v != 0 {
			return false
		}
	}
	return true
}
```