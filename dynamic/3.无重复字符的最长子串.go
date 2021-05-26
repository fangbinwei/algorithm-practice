package leetcode3

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
// 动态规划
// dp[i]代表以第i个字符结尾的最长无重复子串
// 若当前字符从未出现过, dp[i] = dp[i-1]+1
// 若当前字符之前出现在j位置, dp[i] = min(i-j, dp[i-1]+1)
func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	ans := 1

	// m用于记录一个ascii字符最后一次出现的位置
	m := make(map[byte]int)
	m[s[0]] = 0

	dp := make([]int, l)

	dp[0] = 1

	for i := 1; i < l; i++ {
		// 当前字符与之前的字符x相同
		if v, ok := m[s[i]]; ok {
			// x包含在dp[i-1]所代表的字符串内, 则为i-v
			// 否则为dp[i-1]+1
			dp[i] = min(i-v, dp[i-1]+1)
		} else {
			dp[i] = dp[i-1] + 1
		}
		m[s[i]] = i

		ans = max(ans, dp[i])
	}

	return ans

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// @lc code=end
