/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (36.64%)
 * Likes:    5117
 * Dislikes: 0
 * Total Accepted:    877.3K
 * Total Submissions: 2.4M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 * 示例 4:
 *
 *
 * 输入: s = ""
 * 输出: 0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * s 由英文字母、数字、符号和空格组成
 *
 *
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)

	ans := 0

	count := 0
	// 无重复字符左侧起点
	left := 0
	for i := 0; i < len(s); i++ {
		// l>=left 说明当前字符串内存在与s[i]重复的字符,其index为l
		if l, ok := m[s[i]]; ok && l >= left {
			ans = max(ans, count)
			count = i - l
			left = l + 1
		} else {
			count++
		}

		m[s[i]] = i
	}
	// 如果if ok的分支没有执行过, 需要手动赋值一次
	ans = max(ans, count)
	return ans

}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// @lc code=end

