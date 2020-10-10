package main

/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 *
 * https://leetcode-cn.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (45.52%)
 * Likes:    554
 * Dislikes: 0
 * Total Accepted:    210.1K
 * Total Submissions: 461.5K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
 *
 * 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
 *
 * 你可以假设除了整数 0 之外，这个整数不会以零开头。
 *
 * 示例 1:
 *
 * 输入: [1,2,3]
 * 输出: [1,2,4]
 * 解释: 输入数组表示数字 123。
 *
 *
 * 示例 2:
 *
 * 输入: [4,3,2,1]
 * 输出: [4,3,2,2]
 * 解释: 输入数组表示数字 4321。
 *
 *
 */

// @lc code=start
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + 1
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			return digits
		}
	}
	// digits[0] = 1
	// digits = append(digits, 0) // 会扩容
	digits = append([]int{1}, digits...)
	// 利用初始化的都是0
	// 也可以直接 newDigits := make([]int, len(digits) + 1)
	// newDigits[0] = 1
	return digits
}

// @lc code=end

func plusOne(digits []int) []int {
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		digits[i] = digits[i] + 1
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			break
		}
	}
	// 这里有优化的余地
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
