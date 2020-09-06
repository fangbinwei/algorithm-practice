package main

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
	m := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	result := []string{}
	backtracking(&result, []byte{}, 0, digits, m)
	return result

}

func backtracking(result *[]string, sub []byte, start int, digits string, m map[byte][]byte) {
	if start == len(digits) {
		if len(sub) == 0 {
			return
		}
		*result = append(*result, string(sub))
		return
	}

	for _, item := range m[digits[start]] {
		sub = append(sub[:len(sub):len(sub)], item)
		backtracking(result, sub, start+1, digits, m)
		sub = sub[:len(sub)-1]
	}

}

// @lc code=end
