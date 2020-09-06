package main

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 *
 * https://leetcode-cn.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (64.47%)
 * Likes:    727
 * Dislikes: 0
 * Total Accepted:    209.9K
 * Total Submissions: 325.6K
 * Testcase Example:  '[3,2,3]'
 *
 * 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
 *
 * 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [3,2,3]
 * 输出: 3
 *
 * 示例 2:
 *
 * 输入: [2,2,1,1,1,2,2]
 * 输出: 2
 *
 *
 */

// @lc code=start

func majorityElement(nums []int) int {
	count := func(nums []int, n int) int {
		c := 0
		for _, item := range nums {
			if item == n {
				c++
			}
		}
		return c
	}
	if len(nums) == 1 {
		return nums[0]
	}

	left := nums[:len(nums)/2]
	right := nums[len(nums)/2:]

	leftMajority := majorityElement(left)
	rightMajority := majorityElement(right)

	if leftMajority == rightMajority {
		return leftMajority
	}

	lfCount := count(nums, leftMajority)
	rgCount := count(nums, rightMajority)

	if lfCount > rgCount {
		return leftMajority
	} else {
		return rightMajority
	}
}

// @lc code=end

func majorityElement(nums []int) int {
	sort.Ints(nums)

	result := math.MinInt64

	nums = append([]int{math.MinInt64}, nums...)
	nums = append(nums, math.MinInt64)

	count := 1
	var curNum int
	for i, l := 1, len(nums); i < l; i++ {

		if nums[i] == nums[i-1] {
			count++
			continue
		}
		max := int(math.Max(float64(result), float64(count)))
		if max != result {
			if nums[i-1] != math.MinInt64 {
				curNum = nums[i-1]
				result = max
			}
		}
		count = 1
	}
	// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
	return curNum

}

// hashmap
func majorityElement(nums []int) int {
	m := make(map[int]int)
	l := len(nums) / 2
	if l == 0 {
		return nums[1]
	}

	for _, n := range nums {
		if _, ok := m[n]; ok {
			m[n] += 1
			if m[n] > l {
				return n
			}
		} else {
			m[n] = 1
		}
	}
	return 0

	// max := math.MinInt64
	// var result int
	// for key, item := range m {
	// 	if item > max {
	// 		max = item
	// 		result = key
	// 	}
	// }

}
