package leetcode74

/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 *
 * https://leetcode-cn.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (39.51%)
 * Likes:    294
 * Dislikes: 0
 * Total Accepted:    76.3K
 * Total Submissions: 193.1K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3'
 *
 * 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
 *
 *
 * 每行中的整数从左到右按升序排列。
 * 每行的第一个整数大于前一行的最后一个整数。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,50]], target = 3
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,50]], target = 13
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：matrix = [], target = 0
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 0
 * -10^4
 *
 *
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	top := 0
	bottom := len(matrix) - 1
	col := len(matrix[0]) - 1

	for top <= bottom {
		mid := top + (bottom-top)/2

		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][col] == target {
			return true
		}
		// 在mid行
		if matrix[mid][0] < target && target < matrix[mid][col] {
			return binarySearch(matrix[mid], target)
		} else if matrix[mid][0] > target {
			bottom = mid - 1
		} else {
			top = mid + 1
		}
	}

	return false
}
func binarySearch(line []int, target int) bool {
	left := 0
	right := len(line) - 1

	for left <= right {
		mid := left + (right-left)/2
		if line[mid] == target {
			return true
		}
		if line[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// @lc code=end
