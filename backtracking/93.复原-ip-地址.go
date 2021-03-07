package leetcode93

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 *
 * https://leetcode-cn.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (51.76%)
 * Likes:    516
 * Dislikes: 0
 * Total Accepted:    102.1K
 * Total Submissions: 197.3K
 * Testcase Example:  '"25525511135"'
 *
 * 给定一个只包含数字的字符串，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
 *
 * 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
 *
 * 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
 * 和 "192.168@1.1" 是 无效 IP 地址。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "25525511135"
 * 输出：["255.255.11.135","255.255.111.35"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "0000"
 * 输出：["0.0.0.0"]
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "1111"
 * 输出：["1.1.1.1"]
 *
 *
 * 示例 4：
 *
 *
 * 输入：s = "010010"
 * 输出：["0.10.0.10","0.100.1.0"]
 *
 *
 * 示例 5：
 *
 *
 * 输入：s = "101023"
 * 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * s 仅由数字组成
 *
 *
 */

// @lc code=start
func restoreIpAddresses(s string) []string {

	segments := make([]int, 4)

	ans := make([]string, 0)

	backtracking(s, 0, 0, segments, &ans)

	return ans
}

func backtracking(s string, segId int, segStart int, segments []int, ans *[]string) {
	// 4段集齐了
	if segId == 4 {
		// 刚好把所有数字用完了
		if segStart == len(s) {
			*ans = append(*ans, intToString(segments))
			return
		}
		return
	}
	// 数字用完了, 但是还没集齐4段
	if segStart == len(s) {
		return
	}

	// 数字0 特殊处理
	if s[segStart] == '0' {
		segments[segId] = 0
		backtracking(s, segId+1, segStart+1, segments, ans)
		return
	}

	address := 0

	for segEnd := segStart; segEnd < len(s); segEnd++ {
		address = address*10 + int(s[segEnd]-'0')
		if address > 0 && address <= 255 {
			segments[segId] = address
			backtracking(s, segId+1, segEnd+1, segments, ans)
		} else {
			break
		}
	}

}
func intToString(input []int) string {
	ans := make([]string, len(input))
	for i, v := range input {
		ans[i] = strconv.Itoa(v)
	}
	return strings.Join(ans, ".")
}

// @lc code=end
