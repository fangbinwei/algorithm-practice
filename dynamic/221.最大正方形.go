package leetcode221

/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 *
 * https://leetcode-cn.com/problems/maximal-square/description/
 *
 * algorithms
 * Medium (43.86%)
 * Likes:    652
 * Dislikes: 0
 * Total Accepted:    86.5K
 * Total Submissions: 197.2K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix =
 * [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
 * 输出：4
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [["0","1"],["1","0"]]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：matrix = [["0"]]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1
 * matrix[i][j] 为 '0' 或 '1'
 *
 *
 */

// @lc code=start
// 压缩空间
// https://leetcode-cn.com/problems/maximal-square/solution/dong-tai-gui-hua-kong-jian-you-hua-zhu-xing-jie-2/
func maximalSquare(matrix [][]byte) int {
	row := len(matrix)
	col := len(matrix[0])

	m := 0
	dp := make([]int, col+1)

	pre := 0

	for i := 1; i < row+1; i++ {
		for j := 1; j < col+1; j++ {
			// 记录(i-1, j-1)的数据
			tmp := dp[j]
			if matrix[i-1][j-1] == '0' {
				dp[j] = 0
			} else {
				dp[j] = minInt(dp[j], dp[j-1], pre) + 1
				m = maxInt(m, dp[j])
			}
			pre = tmp
		}
	}

	return m * m
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

// @lc code=end

func maximalSquare2(matrix [][]byte) int {
	row := len(matrix)
	col := len(matrix[0])

	m := 0
	//以(i,j)为右下角的正方形最大边长
	dp := make([][]int, row+1)

	for i := range dp {
		dp[i] = make([]int, col+1)
	}

	for i := 1; i < row+1; i++ {
		for j := 1; j < col+1; j++ {
			if matrix[i-1][j-1] == '0' {
				dp[i][j] = 0
			} else {
				dp[i][j] = minInt(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				m = maxInt(m, dp[i][j])
			}
		}
	}

	return m * m
}
