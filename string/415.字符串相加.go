package leetcode415

/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 *
 * https://leetcode-cn.com/problems/add-strings/description/
 *
 * algorithms
 * Easy (51.90%)
 * Likes:    330
 * Dislikes: 0
 * Total Accepted:    100.6K
 * Total Submissions: 193.8K
 * Testcase Example:  '"0"\n"0"'
 *
 * 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
 *
 *
 *
 * 提示：
 *
 *
 * num1 和num2 的长度都小于 5100
 * num1 和num2 都只包含数字 0-9
 * num1 和num2 都不包含任何前导零
 * 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式
 *
 *
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)

	maxL := max(l1, l2)

	s := make([]byte, maxL)

	for i := 0; i < maxL; i++ {
		// i 为0 时, maxL - i -1对应 个位
		// i 为1 时, maxL - i -1对应 十位
		s[maxL-i-1] += '0'
		if i < l1 {
			s[maxL-i-1] += num1[l1-i-1] - '0'

		}
		if i < l2 {
			s[maxL-i-1] += num2[l2-i-1] - '0'
		}
		if s[maxL-i-1] >= 10+'0' {
			s[maxL-i-1] -= 10
			if i == maxL-1 {
				return "1" + string(s)
			} else {
				s[maxL-i-2] += 1
			}
		}
	}
	return string(s)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// @lc code=end
