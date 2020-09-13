package leetcode_1143

/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 *
 * https://leetcode-cn.com/problems/longest-common-subsequence/description/
 *
 * algorithms
 * Medium (60.33%)
 * Likes:    230
 * Dislikes: 0
 * Total Accepted:    43.8K
 * Total Submissions: 72.6K
 * Testcase Example:  '"abcde"\n"ace"'
 *
 * 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。
 *
 * 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
 * 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde"
 * 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。
 *
 * 若这两个字符串没有公共子序列，则返回 0。
 *
 *
 *
 * 示例 1:
 *
 * 输入：text1 = "abcde", text2 = "ace"
 * 输出：3
 * 解释：最长公共子序列是 "ace"，它的长度为 3。
 *
 *
 * 示例 2:
 *
 * 输入：text1 = "abc", text2 = "abc"
 * 输出：3
 * 解释：最长公共子序列是 "abc"，它的长度为 3。
 *
 *
 * 示例 3:
 *
 * 输入：text1 = "abc", text2 = "def"
 * 输出：0
 * 解释：两个字符串没有公共子序列，返回 0。
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= text1.length <= 1000
 * 1 <= text2.length <= 1000
 * 输入的字符串只含有小写英文字符。
 *
 *
 */

// @lc code=start
// 压缩空间 https://leetcode-cn.com/problems/longest-common-subsequence/solution/a-fei-xue-suan-fa-zhi-by-a-fei-8/
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([]int, len(text2)+1)

	for i, l := 1, len(text1); i <= l; i++ {
		// 此时dp中存放的是i-1对应的值
		prev := 0
		for j, l2 := 1, len(text2); j <= l2; j++ {

			if text1[i-1] == text2[j-1] {
				tmp := dp[j]
				// prev 对应i-1,j,  下一次j循环, j +1,  pre正是所需要的值
				dp[j] = prev + 1
				prev = tmp
				// dp[j],prev = prev + 1, dp[j]
			} else {
				prev = dp[j]
				dp[j] = maxInt(dp[j], dp[j-1])
			}
		}
	}

	return dp[len(text2)]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

// dp 无压缩空间
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for row := range dp {
		dp[row] = make([]int, len(text2)+1)
		for col := range dp[row] {
			if col == 0 || row == 0 {
				dp[row][col] = 0
			}
		}
	}
	for i, l := 1, len(text1); i <= l; i++ {
		for j, l2 := 1, len(text2); j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

// dp 递归+ memo
func longestCommonSubsequence(text1 string, text2 string) int {
	memo := make([][]int, len(text1))
	for row := range memo {
		memo[row] = make([]int, len(text2))
		for col := range memo[row] {
			memo[row][col] = -1
		}
	}
	var dp func(l1, l2 int) int
	dp = func(l1, l2 int) int {
		if l1 == -1 {
			return 0
		}
		if l2 == -1 {
			return 0
		}
		if memo[l1][l2] != -1 {
			return memo[l1][l2]
		}
		if text1[l1] == text2[l2] {
			res := dp(l1-1, l2-1) + 1
			memo[l1][l2] = res
			return res
		} else {
			// 有重叠子问题, dp(l1-1, l2) 可能会递归到dp(l1-1,l2-1)
			// dp(l1,l2-1)也可能会递归到dp(l1-1,l2-1)
			res := maxInt(dp(l1-1, l2), dp(l1, l2-1))
			memo[l1][l2] = res
			return res
		}
	}

	return dp(len(text1)-1, len(text2)-1)
}

// 暴力回溯    内存超出
func longestCommonSubsequence(text1 string, text2 string) int {
	subText1 := make([]string, 0)
	subText2 := make([]string, 0)
	sub := make([]byte, 0)
	backtracking(text1, 0, sub, &subText1)
	backtracking(text2, 0, sub, &subText2)
	maxCommon := 0

	for _, item := range subText1 {
		for _, item2 := range subText2 {
			if item == item2 {
				maxCommon = maxInt(maxCommon, len(item))
			}
		}
	}
	return maxCommon
}
func backtracking(text string, start int, sub []byte, result *[]string) {
	for i, l := start, len(text); i < l; i++ {
		sub = append(sub, text[i])
		*result = append(*result, string(sub))
		backtracking(text, i+1, sub, result)
		sub = sub[:len(sub)-1]
	}
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
