/*
 * @lc app=leetcode.cn id=239 lang=typescript
 *
 * [239] 滑动窗口最大值
 *
 * https://leetcode-cn.com/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (48.71%)
 * Likes:    503
 * Dislikes: 0
 * Total Accepted:    68K
 * Total Submissions: 139.7K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k
 * 个数字。滑动窗口每次只向右移动一位。
 * 
 * 返回滑动窗口中的最大值。
 * 
 * 
 * 
 * 进阶：
 * 
 * 你能在线性时间复杂度内解决此题吗？
 * 
 * 
 * 
 * 示例:
 * 
 * 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
 * 输出: [3,3,5,5,6,7] 
 * 解释: 
 * 
 * ⁠ 滑动窗口的位置                最大值
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * 1 <= k <= nums.length
 * 
 * 
 */

// @lc code=start
// 单调队列
// https://leetcode.com/problems/sliding-window-maximum/discuss/65885/This-is-a-typical-monotonic-queue-problem
// https://leetcode.com/problems/sliding-window-maximum/discuss/65884/Java-O(n)-solution-using-deque-with-explanation
function maxSlidingWindow(nums: number[], k: number): number[] {
  if (nums.length * k === 0) return []
  if (nums.length === 1) return nums
  // 1. 窗口就是一个队列, 前面的数不断地出队列, 后面的数不断进队列
  // 2. 当后面的数准备进队列的时候, 前面比它小的数, 需要被删除, 因为轮不到它们当最大值, 前面只剩比它大的数
  // 队列内单调递减(不严格)
  // 3. 当后面的数进队列的时候, 因为 (2) 前面比它小的数, 都被剔除了
  // 只可能剩比它大的数, 等待它们出队列就行了
  const outputLength = nums.length - k + 1
  const output: number[] = []
  const queue: number[] = []

  // 要比较方便知道哪些数轮到出队列, 我们的队列中保存数字的index
  // 删除出队的数, 保证(2)条件, i 为入队数字的index
  function cleanQueue(i: number, k: number) {
    if (queue.length && queue[0] === i - k) queue.shift()

    while(queue.length && nums[i] > nums[queue[queue.length - 1]]) {
      // pop之后, length会减1, 所以while循环会继续判断
      queue.pop()
    }
  }

  // init queue
  for (let i = 0; i < k; i++) {
    cleanQueue(i, k)
    queue.push(i)
  }
  output[0] = nums[queue[0]]

  for (let i = k, l = nums.length; i < l; i++) {
    cleanQueue(i, k)
    queue.push(i)
    output[output.length] = nums[queue[0]]
  }
  return output

};
// @lc code=end

