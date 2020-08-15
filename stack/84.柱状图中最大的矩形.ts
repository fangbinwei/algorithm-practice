/*
 * @lc app=leetcode.cn id=84 lang=typescript
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
// 单调栈, 固定高, 找最大的宽
// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/solution/xiang-jie-dan-diao-zhan-bi-xu-miao-dong-by-sweetie/
function largestRectangleArea(heights: number[]): number {
  heights = [0, ...heights, 0]
  let stack : Array<number> = []
  let area = 0
  for (let i = 0; i < heights.length; i++) {
    while(stack.length && heights[i] < heights[stack[stack.length - 1]]) {
      let topIndex = stack.pop() as number
      area = Math.max(area, heights[topIndex] * (i -stack[stack.length - 1] -1)) 
    }
    stack.push(i)
  }
  return area
};
// @lc code=end


// 暴力求解
function largestRectangleArea(heights: number[]): number {
  if (!heights.length) return 0
  let max = heights[0]
  for (let i=0; i< heights.length; i++) {
    let minHeight = heights[i]
    for (let j = i; j< heights.length; j++) {
      let h = Math.min(minHeight, heights[j])
      let area = (j - i + 1)* h
      minHeight = h
      if (area > max) max = area
    }
  }
  return max

};