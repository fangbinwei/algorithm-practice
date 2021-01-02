package leetcode367t2

/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 *
 * https://leetcode-cn.com/problems/valid-perfect-square/description/
 *
 * algorithms
 * Easy (43.51%)
 * Likes:    168
 * Dislikes: 0
 * Total Accepted:    44.2K
 * Total Submissions: 101.7K
 * Testcase Example:  '16'
 *
 * 给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。
 *
 * 说明：不要使用任何内置的库函数，如  sqrt。
 *
 * 示例 1：
 *
 * 输入：16
 * 输出：True
 *
 * 示例 2：
 *
 * 输入：14
 * 输出：False
 *
 *
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	a := sqrt(num)
	if a*a == num {
		return true
	}
	return false
}

func sqrt(x int) int {
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
