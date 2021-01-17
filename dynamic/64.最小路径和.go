package leetcode64

/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 *
 * https://leetcode-cn.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (68.01%)
 * Likes:    757
 * Dislikes: 0
 * Total Accepted:    173.3K
 * Total Submissions: 254.9K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 *
 * 说明：每次只能向下或者向右移动一步。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
 * 输出：7
 * 解释：因为路径 1→3→1→1→1 的总和最小。
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[1,2,3],[4,5,6]]
 * 输出：12
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1
 * 0
 *
 *
 */

// @lc code=start
//
func minPathSum(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])

	// 不需要额外的dp数组, 直接用grid
	// dp := make([][]int, row)

	// for i := range dp {
	// 	dp[i] = make([]int, col)
	// }

	// dp[0][0] = grid[0][0]

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i-1 < 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
				continue
			}
			if j-1 < 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
				continue
			}

			grid[i][j] = minInt(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}
	return grid[row-1][col-1]
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

// 递归+ memory
func minPathSum2(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	memo := make(map[[2]int]int)

	var _minPathSum func(int, int) int

	_minPathSum = func(r, c int) int {
		if r == 0 && c == 0 {
			return grid[0][0]
		}
		if r == 0 {
			return _minPathSum(0, c-1) + grid[0][c]
		}
		if c == 0 {
			return _minPathSum(r-1, 0) + grid[r][0]
		}
		if v, ok := memo[[2]int{r, c}]; ok {
			return v
		}
		v := minInt(_minPathSum(r-1, c), _minPathSum(r, c-1)) + grid[r][c]
		memo[[2]int{r, c}] = v
		return v
	}

	return _minPathSum(row-1, col-1)

}
