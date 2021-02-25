package leetcode1091

/*
 * @lc app=leetcode.cn id=1091 lang=golang
 *
 * [1091] 二进制矩阵中的最短路径
 *
 * https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/description/
 *
 * algorithms
 * Medium (36.55%)
 * Likes:    85
 * Dislikes: 0
 * Total Accepted:    16.3K
 * Total Submissions: 44.6K
 * Testcase Example:  '[[0,1],[1,0]]'
 *
 * 给你一个 n x n 的二进制矩阵 grid 中，返回矩阵中最短 畅通路径 的长度。如果不存在这样的路径，返回 -1 。
 *
 * 二进制矩阵中的 畅通路径 是一条从 左上角 单元格（即，(0, 0)）到 右下角 单元格（即，(n - 1, n -
 * 1)）的路径，该路径同时满足下述要求：
 *
 *
 * 路径途经的所有单元格都的值都是 0 。
 * 路径中所有相邻的单元格应当在 8 个方向之一 上连通（即，相邻两单元之间彼此不同且共享一条边或者一个角）。
 *
 *
 * 畅通路径的长度 是该路径途经的单元格总数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [[0,1],[1,0]]
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[0,0,0],[1,1,0],[1,1,0]]
 * 输出：4
 *
 *
 * 示例 3：
 *
 *
 * 输入：grid = [[1,0,0],[1,1,0],[1,1,0]]
 * 输出：-1
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == grid.length
 * n == grid[i].length
 * 1
 * grid[i][j] 为 0 或 1
 *
 *
 */

// @lc code=start
func shortestPathBinaryMatrix(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	if grid[0][0] != 0 || grid[row-1][col-1] != 0 {
		return -1
	}
	// 特殊情况 [[0]]
	if row == 1 && col == 1 {
		return 1
	}

	queue := make([][2]int, 0)
	queue = append(queue, [2]int{0, 0})
	ans := 0
	grid[0][0] = 1

	// BFS
	for len(queue) > 0 {
		s := len(queue)
		for i := 0; i < s; i++ {
			node := queue[i]

			x := node[0]
			y := node[1]
			nextNode := [8][2]int{
				{x - 1, y - 1},
				{x, y - 1},
				{x + 1, y - 1},
				{x + 1, y},
				{x + 1, y + 1},
				{x, y + 1},
				{x - 1, y + 1},
				{x - 1, y},
			}

			for _, v := range nextNode {
				x := v[0]
				y := v[1]

				if x == row-1 && y == col-1 {
					return ans + 1
				}
				if x < 0 || y < 0 || x >= row || y >= col || grid[x][y] == 1 {
					continue
				}
				queue = append(queue, [2]int{x, y})
				// 很巧妙地使用grid当memory
				grid[x][y] = 1
			}
		}

		queue = queue[s:]
		ans++
	}
	return -1
}

// @lc code=end
