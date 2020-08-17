package main

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 *
 * https://leetcode-cn.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (49.23%)
 * Likes:    8918
 * Dislikes: 0
 * Total Accepted:    1.3M
 * Total Submissions: 2.6M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 *
 * 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
 *
 *
 *
 * 示例:
 *
 * 给定 nums = [2, 7, 11, 15], target = 9
 *
 * 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 *
 *
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	for i := 0; i < len(nums)-1; i++ {
		t := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if t == nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil

}

// @lc code=end

import (
	"sort"
)

func findIndex(s []int, target int, index int) (int, bool) {
	for i := index; i < len(s); i++ {
		if s[i] == target {
			return i, true
		}
	}
	return -1, false
}

func twoSum(nums []int, target int) []int {
	l := len(nums)
	if l < 2 {
		return nil
	}
	tmp := make([]int, l)
	copy(tmp, nums)
	sort.Ints(nums)
	for i, j := 0, l-1; i < j; {
		sum := nums[i] + nums[j]
		if sum == target {
			i1, _ := findIndex(tmp, nums[i], 0)
			var i2 int
			if nums[i] == nums[j] {
				i2, _ = findIndex(tmp, nums[i], i1+1)
			} else {
				i2, _ = findIndex(tmp, nums[j], 0)
			}

			return []int{i1, i2}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return nil

}