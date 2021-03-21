package leetcode695

/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 *
 * https://leetcode-cn.com/problems/max-area-of-island/description/
 *
 * algorithms
 * Medium (64.92%)
 * Likes:    457
 * Dislikes: 0
 * Total Accepted:    80.2K
 * Total Submissions: 123.4K
 * Testcase Example:  '[[1,1,0,0,0],[1,1,0,0,0],[0,0,0,1,1],[0,0,0,1,1]]'
 *
 * 给定一个包含了一些 0 和 1 的非空二维数组 grid 。
 *
 * 一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被
 * 0（代表水）包围着。
 *
 * 找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)
 *
 *
 *
 * 示例 1:
 *
 * [[0,0,1,0,0,0,0,1,0,0,0,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,1,0,0,0],
 * ⁠[0,1,1,0,1,0,0,0,0,0,0,0,0],
 * ⁠[0,1,0,0,1,1,0,0,1,0,1,0,0],
 * ⁠[0,1,0,0,1,1,0,0,1,1,1,0,0],
 * ⁠[0,0,0,0,0,0,0,0,0,0,1,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,1,0,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,0,0,0,0]]
 *
 *
 * 对于上面这个给定矩阵应返回 6。注意答案不应该是 11 ，因为岛屿只能包含水平或垂直的四个方向的 1 。
 *
 * 示例 2:
 *
 * [[0,0,0,0,0,0,0,0]]
 *
 * 对于上面这个给定的矩阵, 返回 0。
 *
 *
 *
 * 注意: 给定的矩阵grid 的长度和宽度都不超过 50。
 *
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	ans := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			// 遍历过的点, 把grid[i][j]设置为0, 相当于用grid当memory
			grid[i][j] = 0
			// 计算岛屿数量
			area := countAreaOfIsland3(&grid, [2]int{i, j}, len(grid), len(grid[0]))
			ans = max(ans, area)
		}
	}
	return ans
}

// DFS
func countAreaOfIsland3(grid *[][]int, start [2]int, iLimit, jLimit int) int {

	// 四连通
	d := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	count := 1

	i, j := start[0], start[1]

	for _, v := range d {
		newI := i + v[0]
		newJ := j + v[1]

		if newI >= 0 && newJ >= 0 && newI < iLimit && newJ < jLimit && (*grid)[newI][newJ] == 1 {
			(*grid)[newI][newJ] = 0
			count += countAreaOfIsland3(grid, [2]int{newI, newJ}, iLimit, jLimit)
		}
	}
	return count
}

// DFS by stack
func countAreaOfIsland2(grid *[][]int, start [2]int, iLimit, jLimit int) int {
	stack := [][2]int{start}

	// 四连通
	d := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	count := 0

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		count++
		i, j := top[0], top[1]

		for _, v := range d {
			newI := i + v[0]
			newJ := j + v[1]

			if newI >= 0 && newJ >= 0 && newI < iLimit && newJ < jLimit && (*grid)[newI][newJ] == 1 {
				stack = append(stack, [2]int{newI, newJ})
				(*grid)[newI][newJ] = 0
			}
		}
	}
	return count
}

// BFS
func countAreaOfIsland(grid *[][]int, start [2]int, iLimit, jLimit int) int {
	queue := [][2]int{start}

	// 四连通
	d := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	count := 0

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		count++
		i, j := head[0], head[1]

		for _, v := range d {
			newI := i + v[0]
			newJ := j + v[1]

			if newI >= 0 && newJ >= 0 && newI < iLimit && newJ < jLimit && (*grid)[newI][newJ] == 1 {
				queue = append(queue, [2]int{newI, newJ})
				(*grid)[newI][newJ] = 0
			}
		}
	}
	return count
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// @lc code=end
