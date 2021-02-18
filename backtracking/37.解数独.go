package leetcode37

/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 *
 * https://leetcode-cn.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (67.01%)
 * Likes:    754
 * Dislikes: 0
 * Total Accepted:    70K
 * Total Submissions: 104.5K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * 编写一个程序，通过填充空格来解决数独问题。
 *
 * 一个数独的解法需遵循如下规则：
 *
 *
 * 数字 1-9 在每一行只能出现一次。
 * 数字 1-9 在每一列只能出现一次。
 * 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
 *
 *
 * 空白格用 '.' 表示。
 *
 *
 *
 * 一个数独。
 *
 *
 *
 * 答案被标成红色。
 *
 * 提示：
 *
 *
 * 给定的数独序列只包含数字 1-9 和字符 '.' 。
 * 你可以假设给定的数独只有唯一解。
 * 给定数独永远是 9x9 形式的。
 *
 *
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	solve(board)

}
// 回溯
func solve(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				var v byte

				for v = '1'; v <= '9'; v++ {
					if isValid(board, i, j, v) {
						board[i][j] = v
						if solve(board) {
							return true
						}
						// 当前值无效 则撤回选择
						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

// 判断假如row行col列填入v是否合法
func isValid(board [][]byte, row int, col int, v byte) bool {
	for i := 0; i < 9; i++ {
		// 检查列
		if board[i][col] == v {
			return false
		}
		// 检查行
		if board[row][i] == v {
			return false
		}
		// 检查box
		if board[row/3*3+i/3][col/3*3+i%3] == v {
			return false
		}
	}
	return true
}

// @lc code=end
