package leetcode127

import "fmt"

/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 *
 * https://leetcode-cn.com/problems/word-ladder/description/
 *
 * algorithms
 * Medium (45.53%)
 * Likes:    661
 * Dislikes: 0
 * Total Accepted:    90.4K
 * Total Submissions: 198.4K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord
 * 的最短转换序列的长度。转换需遵循如下规则：
 *
 *
 * 每次转换只能改变一个字母。
 * 转换过程中的中间单词必须是字典中的单词。
 *
 *
 * 说明:
 *
 *
 * 如果不存在这样的转换序列，返回 0。
 * 所有单词具有相同的长度。
 * 所有单词只由小写字母组成。
 * 字典中不存在重复的单词。
 * 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
 *
 *
 * 示例 1:
 *
 * 输入:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * 输出: 5
 *
 * 解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
 * ⁠    返回它的长度 5。
 *
 *
 * 示例 2:
 *
 * 输入:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * 输出: 0
 *
 * 解释: endWord "cog" 不在字典中，所以无法进行转换。
 *
 */

// @lc code=start
// 双端BFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 0
	}

	count := 1
	wordListMap := make(map[string]bool)

	queue := make([]string, 0, len(wordList))
	queue = append(queue, beginWord)

	beginMap := make(map[string]bool)
	endMap := make(map[string]bool)
	beginMap[beginWord] = true
	endMap[endWord] = true

	listContainEndWord := false
	for _, v := range wordList {
		if v == endWord {
			listContainEndWord = true
		}
		wordListMap[v] = true
	}
	if !listContainEndWord {
		return 0
	}

	for len(beginMap) > 0 && len(endMap) > 0 {
		count++
		if len(beginMap) > len(endMap) {
			beginMap, endMap = endMap, beginMap
		}

		fmt.Println("----", beginMap)
		tmpBeginMap := make(map[string]bool)
		for k, _ := range beginMap {
			delete(beginMap, k)

			for i := 0; i < len(k); i++ {
				h := []byte(k)
				tmp := h[i]
				// 利用word 生成 只有1个字母变化的 新word
				for j := 'a'; j <= 'z'; j++ {
					if tmp != byte(j) {
						h[i] = byte(j)
						s := string(h)

						if endMap[s] {
							return count
						}
						if wordListMap[s] {
							fmt.Println(s)
							tmpBeginMap[s] = true
							// 已经找到的层级节点要删除, 避免死循环
							delete(wordListMap, s)
						}

					}
				}
			}
			// 使用tmp, 否则beginMap会动态变化, range迭代会受影响
			beginMap = tmpBeginMap
		}
	}

	return 0

}

// @lc code=end

// 还能继续优化, 优化方向: 双端BFS
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 0
	}

	count := 1
	queue := make([]string, 0, len(wordList))
	queue = append(queue, beginWord)
	wordListMap := make(map[string]bool)

	listContainEndWord := false

	for _, v := range wordList {
		if v == endWord {
			listContainEndWord = true
		}
		wordListMap[v] = true
	}
	if !listContainEndWord {
		return 0
	}

	for len(queue) > 0 {
		l := len(queue)

		count++
		for i := 0; i < l; i++ {
			head := queue[0]
			queue = queue[1:]
			// O(len(head) * 26)
			for i := 0; i < len(head); i++ {
				h := []byte(head)
				tmp := h[i]
				// 利用word 生成 只有1个字母变化的 新word
				for j := 'a'; j <= 'z'; j++ {
					if tmp != byte(j) {
						h[i] = byte(j)
						s := string(h)

						if s == endWord {
							return count
						}
						if wordListMap[s] {
							queue = append(queue, s)
							// 已经找到的层级节点要删除, 避免死循环
							delete(wordListMap, s)
						}
					}
				}
			}
		}
	}

	return 0

}

// 当wordList数量很多的时候, canChange函数导致时间复杂度很高
func ladderLengthTimeout(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 0
	}
	count := 1
	queue := make([]string, 0, len(wordList))
	queue = append(queue, beginWord)
	pass := make(map[string]bool)

	listContainEndWord := false

	for len(queue) > 0 {
		l := len(queue)

		count++
		for i := 0; i < l; i++ {
			head := queue[0]
			queue = queue[1:]

			for _, v := range wordList {
				if pass[v] {
					continue
				}
				if !listContainEndWord && v == endWord {
					listContainEndWord = true
				}

				if canChange(head, v) {
					if v == endWord {
						return count
					}
					pass[v] = true
					queue = append(queue, v)
				}

			}

		}
		if !listContainEndWord {
			return 0
		}
	}

	return 0

}
func canChange(a, b string) bool {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return true
}
