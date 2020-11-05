package leetcode77

/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 *
 * https://leetcode-cn.com/problems/combinations/description/
 *
 * algorithms
 * Medium (75.87%)
 * Likes:    431
 * Dislikes: 0
 * Total Accepted:    112.9K
 * Total Submissions: 148.7K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
 *
 * 示例:
 *
 * 输入: n = 4, k = 2
 * 输出:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 */

// @lc code=start
// 回溯+剪枝
func combine(n int, k int) [][]int {
	ans := make([][]int, 0, 10)

	var backtracking func(int, []int)
	backtracking = func(start int, collection []int) {
		if len(collection)+n-start < k {
			return
		}
		if len(collection) == k {
			ans = append(ans, collection)
			return
		}
		for i := start; i <= n; i++ {
			temp := collection[:len(collection):len(collection)]

			temp = append(temp, i)
			backtracking(i+1, temp)
		}
	}

	backtracking(1, []int{})

	return ans
}

// @lc code=end
