package leetcode773

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=773 lang=golang
 *
 * [773] 滑动谜题
 *
 * https://leetcode-cn.com/problems/sliding-puzzle/description/
 *
 * algorithms
 * Hard (63.05%)
 * Likes:    117
 * Dislikes: 0
 * Total Accepted:    6.7K
 * Total Submissions: 10.7K
 * Testcase Example:  '[[1,2,3],[4,0,5]]'
 *
 * 在一个 2 x 3 的板上（board）有 5 块砖瓦，用数字 1~5 来表示, 以及一块空缺用 0 来表示.
 *
 * 一次移动定义为选择 0 与一个相邻的数字（上下左右）进行交换.
 *
 * 最终当板 board 的结果是 [[1,2,3],[4,5,0]] 谜板被解开。
 *
 * 给出一个谜板的初始状态，返回最少可以通过多少次移动解开谜板，如果不能解开谜板，则返回 -1 。
 *
 * 示例：
 *
 *
 * 输入：board = [[1,2,3],[4,0,5]]
 * 输出：1
 * 解释：交换 0 和 5 ，1 步完成
 *
 *
 *
 * 输入：board = [[1,2,3],[5,4,0]]
 * 输出：-1
 * 解释：没有办法完成谜板
 *
 *
 *
 * 输入：board = [[4,1,2],[5,0,3]]
 * 输出：5
 * 解释：
 * 最少完成谜板的最少移动次数是 5 ，
 * 一种移动路径:
 * 尚未移动: [[4,1,2],[5,0,3]]
 * 移动 1 次: [[4,1,2],[0,5,3]]
 * 移动 2 次: [[0,1,2],[4,5,3]]
 * 移动 3 次: [[1,0,2],[4,5,3]]
 * 移动 4 次: [[1,2,0],[4,5,3]]
 * 移动 5 次: [[1,2,3],[4,5,0]]
 *
 *
 *
 * 输入：board = [[3,2,4],[1,5,0]]
 * 输出：14
 *
 *
 * 提示：
 *
 *
 * board 是一个如上所述的 2 x 3 的数组.
 * board[i][j] 是一个 [0, 1, 2, 3, 4, 5] 的排列.
 *
 *
 */

// @lc code=start
func slidingPuzzle(board [][]int) int {

	ans := 0

	start := slice2str(board)

	queue := make([]string, 0)
	queue = append(queue, start)

	visited := make(map[string]bool)

	for len(queue) > 0 {
		s := len(queue)
		for i := 0; i < s; i++ {
			top := queue[i]

			visited[top] = true
			// fmt.Println(top)
			if top == "123450" {
				return ans
			}
			// fmt.Println(getNext(top))
			queue = append(queue, getNext(top, visited)...)
		}
		queue = queue[s:]
		ans++
	}

	return -1
}

func slice2str(s [][]int) string {
	row := len(s)
	col := len(s[0])

	t := make([]string, row*col)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			t = append(t, string(byte(s[i][j])+'0'))
		}
	}
	return strings.Join(t, "")
}

func getNext(s string, visited map[string]bool) []string {
	next := make([]string, 0)
	// 0在各个位置的时候, 可以和哪个位置互换
	m := map[int][]int{
		0: {1, 3},
		1: {0, 2, 4},
		2: {1, 5},
		3: {0, 4},
		4: {1, 3, 5},
		5: {2, 4},
	}

	iPosition := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			iPosition = i
		}
	}

	for _, v := range m[iPosition] {
		org := []byte(s)
		org[v], org[iPosition] = org[iPosition], org[v]
		s := string(org)
		if !visited[s] {
			next = append(next, s)
		}
	}

	return next
}

// @lc code=end
