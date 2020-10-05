package leetcode_15

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (28.93%)
 * Likes:    2462
 * Dislikes: 0
 * Total Accepted:    296.2K
 * Total Submissions: 1M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？请你找出所有满足条件且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 * 示例：
 *
 * 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 *
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */

// @lc code=start
import (
	"sort"
)

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0, 3)
	if len(nums) < 3 {
		return ans
	}
	sort.Ints(nums)

	for i, l := 0, len(nums); i < l-1; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]

		left := i + 1
		right := l - 1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < l && nums[left] == nums[left-1] {
					left++
				}

				for right >= 0 && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < target {
				// 这里不需要判断是否nums[left]重复, 因为如果重复, 下次循环sum依旧小于target
				left++
			} else {
				// 同left
				right--

			}
		}

	}

	return ans
}

// @lc code=end
