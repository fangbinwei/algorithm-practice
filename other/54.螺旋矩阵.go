package leetcode54

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode-cn.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (46.69%)
 * Likes:    764
 * Dislikes: 0
 * Total Accepted:    149K
 * Total Submissions: 319K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[1,2,3,6,9,8,7,4,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
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
 * -100
 *
 *
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	col := len(matrix[0])

	ans := make([]int, 0, row*col)

	direction := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	// -100<=matrix[i][j] <=100

	ans = append(ans, matrix[0][0])
	matrix[0][0] = -200

	x, y := 0, 0
	start := 0

	for len(ans) < col*row {

		d := direction[start]
		tmpX := x + d[0]
		tmpY := y + d[1]

		if tmpX == row || tmpY == col || tmpX < 0 || tmpY < 0 {
			start = (start + 1) % 4
			continue
		}

		done := matrix[tmpX][tmpY] == -200
		if done {
			start = (start + 1) % 4
			continue
		}

		x = tmpX
		y = tmpY
		ans = append(ans, matrix[x][y])
		matrix[x][y] = -200

	}

	return ans

}

// @lc code=end

func spiralOrder2(matrix [][]int) []int {

	row := len(matrix)
	col := len(matrix[0])

	ans := make([]int, 0, row*col)

	direction := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	// 优化方向, 使用matrix当mem, -100<=matrix[i][j] <=100
	mem := make(map[[2]int]struct{})

	ans = append(ans, matrix[0][0])
	mem[[2]int{0, 0}] = struct{}{}

	x, y := 0, 0
	start := 0

	for len(ans) < col*row {

		d := direction[start]
		tmpX := x + d[0]
		tmpY := y + d[1]
		_, ok := mem[[2]int{tmpX, tmpY}]

		if tmpX == row || tmpY == col || tmpX < 0 || tmpY < 0 || ok {
			start = (start + 1) % 4
			continue
		}

		x = tmpX
		y = tmpY
		ans = append(ans, matrix[x][y])
		mem[[2]int{x, y}] = struct{}{}

	}

	return ans

}
