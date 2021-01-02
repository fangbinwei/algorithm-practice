package leetcode69t2

import "fmt"

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 *
 * https://leetcode-cn.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (38.88%)
 * Likes:    507
 * Dislikes: 0
 * Total Accepted:    206.3K
 * Total Submissions: 530.7K
 * Testcase Example:  '4'
 *
 * 实现 int sqrt(int x) 函数。
 *
 * 计算并返回 x 的平方根，其中 x 是非负整数。
 *
 * 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
 *
 * 示例 1:
 *
 * 输入: 4
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: 8
 * 输出: 2
 * 说明: 8 的平方根是 2.82842...,
 * 由于返回类型是整数，小数部分将被舍去。
 *
 *
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 1 || x == 2 || x == 3 {
		return 1
	}
	if x == 4 {
		return 2
	}
	var ans int

	if x > 4 {
		left := 2
		right := x / 2
		for left < right {
			mid := (left + right) / 2
			s := mid * mid
			if s == x {
				return mid
			}
			if s < x {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		ans = left
		if left*left > x {
			ans = left - 1
		}
	}
	return ans
}

// @lc code=end
