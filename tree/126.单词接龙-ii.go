package leetcode126

/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 *
 * https://leetcode-cn.com/problems/word-ladder-ii/description/
 *
 * algorithms
 * Hard (38.87%)
 * Likes:    368
 * Dislikes: 0
 * Total Accepted:    26.7K
 * Total Submissions: 68.6K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord
 * 的最短转换序列。转换需遵循如下规则：
 *
 *
 * 每次转换只能改变一个字母。
 * 转换后得到的单词必须是字典中的单词。
 *
 *
 * 说明:
 *
 *
 * 如果不存在这样的转换序列，返回一个空列表。
 * 所有单词具有相同的长度。
 * 所有单词只由小写字母组成。
 * 字典中不存在重复的单词。
 * 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
 *
 *
 * 示例 1:
 *
 * 输入:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * 输出:
 * [
 * ⁠ ["hit","hot","dot","dog","cog"],
 * ["hit","hot","lot","log","cog"]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * 输出: []
 *
 * 解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。
 *
 */

// @lc code=start

type ladder struct {
	current string
	prev    *ladder
}

// 优化方向, 双向BFS
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	ans := make([][]string, 0)

	m := make(map[string]bool)

	containEndWord := false
	for _, v := range wordList {
		if v == endWord {
			containEndWord = true
		}
		m[v] = true
	}
	// wordlist 不包含endword
	if !containEndWord {
		return ans
	}

	var first ladder
	ladderMap := make(map[string]bool)
	ladderMap[beginWord] = true

	first = ladder{
		current: beginWord,
		prev:    nil,
	}

	// BFS队列
	beginWordQueue := []ladder{
		first,
	}

	// 记录层数
	count := 1
	// 找到最短路径
	find := false

	visited := make(map[string]bool)
	visited[beginWord] = true

	for len(beginWordQueue) > 0 {
		count++

		// 记录每层访问过的节点, 每层访问完后同步到visited中
		subVisited := make([]string, 0)
		for z, l := 0, len(beginWordQueue); z < l; z++ {
			word := beginWordQueue[0]
			beginWordQueue = beginWordQueue[1:]

			for i := 0; i < len(word.current); i++ {
				w := []byte(word.current)
				for j := 'a'; j <= 'z'; j++ {
					if w[i] == byte(j) {
						continue
					}
					w[i] = byte(j)

					transform := string(w)
					nextWord := ladder{
						current: transform,
						prev:    &word,
					}

					if endWord == transform {
						find = true
						addToResult(&ans, &nextWord)
					} else if m[transform] && !visited[transform] {
						beginWordQueue = append(beginWordQueue, nextWord)
						// 每一层遍历之后, 再统一更新visited
						subVisited = append(subVisited, transform)
					}
				}
			}
		}

		if find {
			return ans
		}

		//更新visited
		for _, v := range subVisited {
			visited[v] = true
		}

	}

	return ans
}

func addToResult(ans *[][]string, word *ladder) {
	r := make([]string, 0, 10)

	for word != nil {
		r = append(r, word.current)
		word = word.prev
	}
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	*ans = append(*ans, r)
}

// @lc code=end
