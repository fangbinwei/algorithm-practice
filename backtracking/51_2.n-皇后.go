// package main

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
	ans := [][]string{}

	queensLoc := make([]int, 0, n)

	var backtracking func(int)

	col := make(map[int]bool)
	leftDiagonal := make(map[int]bool)
	rightDiagonal := make(map[int]bool)

	backtracking = func(start int) {
		if start == n {
			ans = append(ans, print(n, queensLoc))
			return
		}
		for i := 0; i < n; i++ {
			if col[i] {
				continue
			}
			lIndex := i - (len(queensLoc) + 1)
			if leftDiagonal[lIndex] {
				continue
			}
			rIndex := i + len(queensLoc) + 1
			if rightDiagonal[rIndex] {
				continue
			}
			rightDiagonal[rIndex] = true
			leftDiagonal[lIndex] = true
			col[i] = true

			queensLoc = append(queensLoc, i)

			backtracking(start + 1)
			// 回退
			queensLoc = queensLoc[:len(queensLoc)-1]
			col[i] = false
			leftDiagonal[lIndex] = false
			rightDiagonal[rIndex] = false
		}

	}
	backtracking(0)
	return ans
}

func print(n int, queensLoc []int) []string {
	ans := make([]string, 0, n)

	for i := 0; i < n; i++ {
		b := make([]byte, 0, n)
		for j := 0; j < n; j++ {
			if queensLoc[i] == j {
				b = append(b, 'Q')
			} else {
				b = append(b, '.')
			}
		}
		ans = append(ans, string(b))
	}
	return ans
}

// func main() {
// 	fmt.Printf("%v", solveNQueens(4))
// }

// @lc code=end
