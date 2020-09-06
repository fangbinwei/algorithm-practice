package main

/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 *
 * https://leetcode-cn.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (73.06%)
 * Likes:    598
 * Dislikes: 0
 * Total Accepted:    75.1K
 * Total Submissions: 102.8K
 * Testcase Example:  '4'
 *
 * n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 *
 *
 * 上图为 8 皇后问题的一种解法。
 *
 * 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
 *
 * 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
 *
 *
 *
 * 示例：
 *
 * 输入：4
 * 输出：[
 * ⁠[".Q..",  // 解法 1
 * ⁠ "...Q",
 * ⁠ "Q...",
 * ⁠ "..Q."],
 *
 * ⁠["..Q.",  // 解法 2
 * ⁠ "Q...",
 * ⁠ "...Q",
 * ⁠ ".Q.."]
 * ]
 * 解释: 4 皇后问题存在两个不同的解法。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。
 *
 *
 */

// @lc code=start

func solveNQueens(n int) [][]string {

	// leetcode里不要把result放全局作用域,会有问题, 执行不同的测试, 会公用
	// 如果放全局作用域, 需要    result = make([][]string, 0)来初始化下

	var result [][]string
	columnMap := make(map[int]bool)
	diagonal1Map := make(map[int]bool)
	diagonal2Map := make(map[int]bool)
	queens := make([]int, n)
	for i := range queens {
		queens[i] = -1
	}

	backtracking(queens, 0, columnMap, diagonal1Map, diagonal2Map, &result)
	return result
}

func backtracking(queens []int, row int, columnMap, diagonal1Map, diagonal2Map map[int]bool, result *[][]string) {
	col := len(queens)
	if row == col {
		*result = append(*result, generateBoard(queens))
		return
	}

	for c := 0; c < col; c++ {
		// 列
		if columnMap[c] {
			continue
		}

		if diagonal1Map[c-row] {
			continue
		}

		if diagonal2Map[c+row] {
			continue
		}

		queens[row] = c
		columnMap[c] = true
		// 对角线
		// 反对角线
		diagonal1Map[c-row] = true
		diagonal2Map[c+row] = true

		backtracking(queens, row+1, columnMap, diagonal1Map, diagonal2Map, result)

		queens[row] = -1
		delete(columnMap, c)
		delete(diagonal1Map, c-row)
		delete(diagonal2Map, c+row)
	}

}

// queens 第1行queen的列序, 第2行queen的列序
func generateBoard(queens []int) []string {
	row := len(queens)
	res := make([]string, 0, row)

	str := []byte{}
	for r := 0; r < row; r++ {
		for c, l := 0, len(queens); c < l; c++ {
			if queens[r] >= 0 && c == queens[r] {
				str = append(str, 'Q')
			} else {
				str = append(str, '.')
			}
		}
		res = append(res, string(str))
		str = []byte{}
	}
	return res
}

// @lc code=end
