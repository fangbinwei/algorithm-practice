package leetcode_120

/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 *
 * https://leetcode-cn.com/problems/triangle/description/
 *
 * algorithms
 * Medium (66.65%)
 * Likes:    592
 * Dislikes: 0
 * Total Accepted:    106.6K
 * Total Submissions: 159.9K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
 *
 * 相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
 *
 *
 *
 * 例如，给定三角形：
 *
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 *
 *
 * 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
 *
 *
 *
 * 说明：
 *
 * 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
 *
 */

// @lc code=start
// 压缩空间
// 另一种优化方式, 目前是从上往下遍历, 可以从下往上走
// https://leetcode-cn.com/problems/triangle/solution/di-gui-ji-yi-hua-dp-bi-xu-miao-dong-by-sweetiee/
func minimumTotal(triangle [][]int) int {
	row := len(triangle)
	col := len(triangle[row-1])

	dp := make([]int, col+1, col+1)

	// 一共row+1行
	for i := 1; i <= row; i++ {
		for j := len(triangle[i-1]); j >= 1; j-- {
			tr := triangle[i-1][j-1]
			// 第一列数据
			if j == 1 {
				dp[j] = dp[j] + tr
			} else if j == len(triangle[i-1]) {
				// 最后一列数据
				// 因为需要用到j-1的数据, 要倒着遍历j,
				// 保证拿到的dp[j-1], 是dp[i-1][j-1]的值
				dp[j] = dp[j-1] + tr
			} else {
				// 中间数据
				dp[j] = minInt(dp[j], dp[j-1]) + tr
			}
		}
	}

	min := dp[1]

	// 最后一行总共有col+1列, 所以是<=col
	for index := 2; index <= col; index++ {
		min = minInt(min, dp[index])
	}
	return min

}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end

// 无压缩空间
func minimumTotal(triangle [][]int) int {
	row := len(triangle)
	col := len(triangle[row-1])

	dp := make([][]int, row+1)
	for index := range dp {
		dp[index] = make([]int, col+1)
	}

	// 一共row+1行
	for i := 1; i <= row; i++ {
		for j := 1; j <= len(triangle[i-1]); j++ {
			tr := triangle[i-1][j-1]
			// 第一列数据
			if j == 1 {
				dp[i][j] = dp[i-1][j] + tr
			} else if j == len(triangle[i-1]) {
				// 最后一列数据
				dp[i][j] = dp[i-1][j-1] + tr
			} else {
				// 中间数据
				dp[i][j] = minInt(dp[i-1][j]+tr, dp[i-1][j-1]+tr)
			}
		}
	}

	min := dp[row][1]

	// 最后一行总共有col+1列, 所以是<=col
	for index := 2; index <= col; index++ {
		min = minInt(min, dp[row][index])
	}
	return min

}
