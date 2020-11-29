package leetcode17t2

// package main

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (55.43%)
 * Likes:    903
 * Dislikes: 0
 * Total Accepted:    172.6K
 * Total Submissions: 311.4K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 * 示例:
 *
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 *
 */

// @lc code=start
func letterCombinations(digits string) []string {
	ans := make([]string, 0)
	m := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}

	l := len(digits)
	if l == 0 {
		return ans
	}
	var backtracking func([]byte, int)
	backtracking = func(collection []byte, start int) {
		if start == l {
			ans = append(ans, string(collection))
			return
		}

		alternative := m[digits[start]]

		for _, item := range alternative {
			collection = append(collection, item)
			backtracking(collection, start+1)
			collection = collection[:len(collection)-1]
		}
	}
	backtracking([]byte{}, 0)

	return ans
}

// @lc code=end
