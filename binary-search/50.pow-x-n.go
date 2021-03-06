package main

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode-cn.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (36.51%)
 * Likes:    483
 * Dislikes: 0
 * Total Accepted:    122.5K
 * Total Submissions: 335.6K
 * Testcase Example:  '2.00000\n10'
 *
 * 实现 pow(x, n) ，即计算 x 的 n 次幂函数。
 *
 * 示例 1:
 *
 * 输入: 2.00000, 10
 * 输出: 1024.00000
 *
 *
 * 示例 2:
 *
 * 输入: 2.10000, 3
 * 输出: 9.26100
 *
 *
 * 示例 3:
 *
 * 输入: 2.00000, -2
 * 输出: 0.25000
 * 解释: 2^-2 = 1/2^2 = 1/4 = 0.25
 *
 * 说明:
 *
 *
 * -100.0 < x < 100.0
 * n 是 32 位有符号整数，其数值范围是 [−2^31, 2^31 − 1] 。
 *
 *
 */

// @lc code=start
// binary-search
// 迭代版本
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	// 利用快速幂, 二进制展开
	// https://leetcode-cn.com/problems/powx-n/solution/powx-n-by-leetcode-solution/

	var result float64 = 1
	flag := false
	if n < 0 {
		n = -n
		flag = true
	}

	for n > 0 {
		// n = bm...b3 b2 b1 
		// x^n = x^(1*b1)x^(2*b2)x^(4*b3)
		// 判断bm是否为1, 为1 要累计到结果中
		if n%2 == 1 {
			result = result * x
		}

		x = x * x
		n = n / 2

	}
	if flag {
		result = 1 / result
	}

	return result

}

// @lc code=end

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n > 0 {
		if n%2 == 1 {
			return myPow(x, n+1) / x
		} else {
			half := myPow(x, n/2)
			return half * half
		}
	} else {
		return 1 / myPow(x, -n)
	}
}
