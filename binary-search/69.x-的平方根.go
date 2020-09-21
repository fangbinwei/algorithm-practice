package leetcode_69

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
// 二分法, ans^2 <= x, [0,x] 进行二分
func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	left := 0
	right := x

	var mid int
	for left <= right {
		mid = (left + right) / 2
		a := mid * mid

		if a == x {
			return mid
		}
		if a > x {
			// (1)大于right的值, 肯定不会为解
			// 因为它们使得mid*mid >x
			right = mid - 1
		} else if a < x {
			// left不会为解
			// 循环结束的时候, left>right
			// 而大于right的值, 肯定不会为解, 见(1)
			left = mid + 1
		}
	}
	return right
}

// @lc code=end
