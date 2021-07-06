# [127. 单词接龙](https://leetcode-cn.com/problems/word-ladder/)


## 题目

字典 `wordList` 中从单词 `beginWord` 和 `endWord` 的 转换序列 是一个按下述规格形成的序列：

- 序列中第一个单词是 beginWord 。
- 序列中最后一个单词是 endWord 。
- 每次转换只能改变一个字母。
- 转换过程中的中间单词必须是字典 wordList 中的单词。

给你两个单词 `beginWord` 和 `endWord` 和一个字典 `wordList` ，找到从 `beginWord` 到 `endWord` 的 **最短转换序列** 中的 **单词数目** 。如果不存在这样的转换序列，返回 0。


**Example 1:**

```
输入：
beginWord = "hit", 
endWord = "cog", 
wordList = ["hot","dot","dog","lot","log","cog"]

输出：5

解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。
```

**Example 2:**

```
输入：
beginWord = "hit", 
endWord = "cog", 
wordList = ["hot","dot","dog","lot","log"]

输出：0

解释：endWord "cog" 不在字典中，所以无法进行转换。
```    

**提示：**

- 1 <= beginWord.length <= 10
- endWord.length == beginWord.length
- 1 <= wordList.length <= 5000
- wordList[i].length == beginWord.length
- beginWord、endWord 和 wordList[i] 由小写英文字母组成
- beginWord != endWord
- wordList 中的所有字符串 互不相同
  

## 题目大意

给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：

1. 每次转换只能改变一个字母。
2. 转换过程中的中间单词必须是字典中的单词。

说明:

- 如果不存在这样的转换序列，返回 0。
- 所有单词具有相同的长度。
- 所有单词只由小写字母组成。
- 字典中不存在重复的单词。
- 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。


## 解题思路

- 这一题要求输出从 `beginWord` 变换到 `endWord` 最短变换次数。可以用 BFS，从 `beginWord` 开始变换，把该单词的每个字母都用 `'a'~'z'` 变换一次，生成的数组到 `wordList` 中查找，这里用 Map 来记录查找。找得到就入队列，找不到就输出 0 。入队以后按照 BFS 的算法依次遍历完，当所有单词都 `len(queue)<=0` 出队以后，整个程序结束。
- 这一题题目中虽然说了要求找到一条最短的路径，但是实际上最短的路径的寻找方法已经告诉你了：
	1. 每次只变换一个字母
	2. 每次变换都必须在 `wordList` 中  
所以不需要单独考虑何种方式是最短的。 

- 类似题目有 `429. N 叉树的层序遍历`

## 代码

**Java版**

```java
class Solution {
    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        Set<String> wordSet = new HashSet<>(wordList);
        Queue<String> queue = new LinkedList<>();
        queue.offer(beginWord);
        int res = 0;
        while (!queue.isEmpty()) {
            int n = queue.size();
            for (int i = 0; i < n; i++) {
                String word = queue.poll();
                if (word.equals(endWord)) {
                    return res + 1;
                }
                for (int j = 0; j < word.length(); j++) {
                    char[] cw = word.toCharArray();
                    for (char c = 'a'; c <= 'z'; c++) {
                        cw[j] = c;
                        String newWord = new String(cw);
                        if (!newWord.equals(word) && wordSet.contains(newWord)) {
                            wordSet.remove(newWord);
                            queue.offer(newWord);
                        }
                    }
                }
            }
            res++;
        }
        return 0;
    }
}
```

**Golang版**

```go
package leetcode

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap, queue, res := buildWordMap(wordList), []string{beginWord}, 0

	for l := len(queue); len(queue) != 0; l = len(queue) {
		for i := 0; i < l; i++ {
			word := queue[i]
			if word == endWord {
				return res + 1
			}
			for j := 0; j < len(word); j++ {
				cw := []byte(word)
				for c := 'a'; c <= 'z'; c++ {
					cw[j] = byte(c)
					newWord := string(cw)
					if newWord == word {
						continue
					}
					if _, ok := wordMap[newWord]; ok {
						delete(wordMap, newWord)
						queue = append(queue, newWord)
					}
				}

			}
		}
		queue = queue[l:]
		res++
	}
	return 0
}

func buildWordMap(wordList []string) map[string]int {
	wordMap := make(map[string]int, len(wordList))
	for i, word := range wordList {
		wordMap[word] = i
	}
	return wordMap
}
```
