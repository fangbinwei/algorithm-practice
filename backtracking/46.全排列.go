package leetcode46

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (77.09%)
 * Likes:    969
 * Dislikes: 0
 * Total Accepted:    213.1K
 * Total Submissions: 276.5K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
 *
 * 示例:
 *
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 */

// @lc code=start
func permute(nums []int) [][]int {
	ans := make([][]int, 0, 10)
	stop := len(nums)
	if stop == 0 {
		return ans
	}
	collection := make([]int, 0)

	var _permute func([]int, []int)
	_permute = func(nums []int, collection []int) {
		// 每次执行collection都是新构造的
		if len(collection) == stop {
			ans = append(ans, collection)
			return
		}
		for i, v := range nums {
			nums[0],nums[i] = nums[i], nums[0]
			_permute(nums[1:], append(collection[:len(collection):len(collection)], v))
			nums[0],nums[i] = nums[i], nums[0]
			// left := nums[:i:i]
			// right := nums[i+1:]
			// tmp := collection[:len(collection):len(collection)]
			// tmp = append(tmp, v)
			// remain := append(left, right...)
			// _permute(remain, tmp)
		}
	}
	_permute(nums, collection)

	return ans
}

// @lc code=end

func permute2(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) == 0 {
		return ans
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	// 全程使用同一个q
	q := []int{}

	var backtracking func(n []int)
	backtracking = func(n []int) {
		if len(n) == 0 {
			ans = append(ans, append([]int{}, q...))
			return
		}
		for i := 0; i < len(n); i++ {
			q = append(q, n[i])
			res := append(n[:i:i], n[i+1:]...)
			backtracking(res)
			q = q[:len(q)-1]
		}

	}
	backtracking(nums)

	return ans
}
