# [49. 字母异位词分组](https://leetcode-cn.com/problems/group-anagrams/)

## 题目

给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。


**Example:**

```
输入: ["eat", "tea", "tan", "ate", "nat", "bat"]

输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]

```

Note:

- 所有输入均为小写字母。
- 不考虑答案输出的顺序。

## 题目大意

给出一个字符串数组，要求对字符串数组里面有 字母异位词 关系的字符串进行分组。字母异位词 关系是指两个字符串的字符完全相同，顺序不同。

## 解题思路

建立一个map， key通过一个26有序字母数组获得（把字符串拆分成有序字符数组再转string），相同的key存入一个list中，最后返回map。


## 代码

**Java版**

```java
class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> map = new HashMap<>();
        for (String str : strs) {
            char[] cs = new char[26];
            for (char c : str.toCharArray()) {
                cs[c - 'a']++;
            }

            String key = String.valueOf(cs);
            map.putIfAbsent(key, new ArrayList<>());
            map.get(key).add(str);
        }
        return new ArrayList<>(map.values());
    }
}
```

**Golang版**

```go
package leetcode

func groupAnagrams(strs []string) [][]string {

	m := make(map[string][]string)
	for _, str := range strs {
		cs := make([]byte, 26)
		for _, c := range str {
			cs[c-'a']++
		}
		if v, ok := m[string(cs)]; ok {
			m[string(cs)] = append(v, str)
		} else {
			m[string(cs)] = []string{str}
		}
	}
	res := make([][]string, len(m))
	i := 0
	for _, strings := range m {
		res[i] = strings
		i++
	}
	return res
}
```
