/*
 * @lc app=leetcode.cn id=11 lang=typescript
 *
 * [11] 盛最多水的容器
 *
 * https://leetcode-cn.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (63.91%)
 * Likes:    1710
 * Dislikes: 0
 * Total Accepted:    258.7K
 * Total Submissions: 404.7K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
 * (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 * 
 * 说明：你不能倾斜容器，且 n 的值至少为 2。
 * 
 * 
 * 
 * 
 * 
 * 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：[1,8,6,2,5,4,8,3,7]
 * 输出：49
 * 
 */

 // 双指针
// @lc code=start
function maxArea(height: number[]): number {
  let maxWater = 0
  let i = 0
  let j = height.length - 1

  while (i < j) {
    maxWater = Math.max(maxWater, Math.min(height[i], height[j]) * (j-i))
    if (height[i] < height[j]) {
      // 此时右侧边如果左移, 宽减小, 且高度是min(左侧边高, 新的高) <= 左侧边高, 容量一定变小
      i++
    } else {
      j--
    }
  }
  return maxWater
};
// @lc code=end

