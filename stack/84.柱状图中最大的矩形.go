package leetcode_84

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 *
 * https://leetcode-cn.com/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (41.37%)
 * Likes:    847
 * Dislikes: 0
 * Total Accepted:    79.2K
 * Total Submissions: 191.3K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
 *
 * 求在该柱状图中，能够勾勒出来的矩形的最大面积。
 *
 *
 *
 *
 *
 * 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
 *
 *
 *
 *
 *
 * 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
 *
 *
 *
 * 示例:
 *
 * 输入: [2,1,5,6,2,3]
 * 输出: 10
 *
 */

// @lc code=start
// 单调栈常数优化, 一次循环计算left, right, 并压缩空间
func largestRectangleArea(heights []int) int {
	area := 0
	if len(heights) == 0 {
		return area
	}
	// 增加两个哨兵
	tmp := make([]int, len(heights)+2)
	copy(tmp[1:], heights)
	heightsWithGuard := tmp
	// 单调栈
	monoStack := make([]int, 0, len(heights))

	for i := 0; i < len(heightsWithGuard); i++ {
		// 维护单调栈
		// 和未压缩空间的方法比, 这边必需使用 >, 而不是>=
		for len(monoStack) > 0 && heightsWithGuard[monoStack[len(monoStack)-1]] > heightsWithGuard[i] {
			area = maxInt(area, heightsWithGuard[monoStack[len(monoStack)-1]]*(i-monoStack[len(monoStack)-2]-1))
			monoStack = monoStack[:len(monoStack)-1]
		}

		monoStack = append(monoStack, i)

	}

	return area
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

// 单调栈常数优化, 一次循环计算left, right
func largestRectangleArea2_1(heights []int) int {
	area := 0
	if len(heights) == 0 {
		return area
	}
	// 增加两个哨兵
	tmp := make([]int, len(heights)+2)
	copy(tmp[1:], heights)
	heightsWithGuard := tmp

	left := make([]int, len(heights)+2)
	right := make([]int, len(heights)+2)
	// 单调栈
	monoStack := make([]int, 0, len(heights))

	for i := 0; i < len(heightsWithGuard); i++ {
		// 维护单调栈
		for len(monoStack) > 0 && heightsWithGuard[monoStack[len(monoStack)-1]] >= heightsWithGuard[i] {
			right[monoStack[len(monoStack)-1]] = i
			monoStack = monoStack[:len(monoStack)-1]
		}

		if len(monoStack) == 0 {
			left[i] = 0
		} else {
			left[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)

	}

	for i := 0; i < len(heightsWithGuard); i++ {
		area = maxInt(area, heightsWithGuard[i]*(right[i]-left[i]-1))
	}

	return area
}

// 使用单调栈优化暴力算法
func largestRectangleArea2(heights []int) int {
	area := 0
	if len(heights) == 0 {
		return area
	}

	left := make([]int, len(heights))
	right := make([]int, len(heights))
	// 单调栈
	monoStack := make([]int, 0, len(heights))

	for i := 0; i < len(heights); i++ {
		// 维护单调栈
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}

		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			left[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)

	}

	monoStack = make([]int, 0, len(heights))

	for i := len(heights) - 1; i >= 0; i-- {
		// 维护单调栈
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}

		if len(monoStack) == 0 {
			right[i] = len(heights)
		} else {
			right[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}

	for i := 0; i < len(heights); i++ {
		area = maxInt(area, heights[i]*(right[i]-left[i]-1))
	}

	return area
}

// 两种暴力法
// 1. 固定宽找对应高
// 2. 固定高找对应宽
// 这里采用2
func largestRectangleArea1(heights []int) int {
	area := 0
	if len(heights) == 0 {
		return area
	}
	// 增加两个哨兵
	tmp := make([]int, len(heights)+2)
	copy(tmp[1:], heights)
	heightsWithGuard := tmp

	left := make([]int, 0, len(heights))
	right := make([]int, 0, len(heights))

	// 固定高, 找宽
	for i := 1; i < len(heightsWithGuard)-1; i++ {
		if heightsWithGuard[i] == 0 {
			left = append(left, 0)
			right = append(right, 0)
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if heightsWithGuard[j] < heightsWithGuard[i] {
				left = append(left, j)
				break
			}
		}
		for j := i + 1; j < len(heightsWithGuard); j++ {
			if heightsWithGuard[j] < heightsWithGuard[i] {
				right = append(right, j)
				break
			}
		}
	}

	for i := 0; i < len(heights); i++ {
		area = maxInt(area, heights[i]*(right[i]-left[i]-1))
	}

	return area
}
