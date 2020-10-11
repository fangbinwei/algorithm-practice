package leetcode_155

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 *
 * https://leetcode-cn.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (55.06%)
 * Likes:    640
 * Dislikes: 0
 * Total Accepted:    151.1K
 * Total Submissions: 274.4K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
 *
 *
 * push(x) —— 将元素 x 推入栈中。
 * pop() —— 删除栈顶的元素。
 * top() —— 获取栈顶元素。
 * getMin() —— 检索栈中的最小元素。
 *
 *
 *
 *
 * 示例:
 *
 * 输入：
 * ["MinStack","push","push","push","getMin","pop","top","getMin"]
 * [[],[-2],[0],[-3],[],[],[],[]]
 *
 * 输出：
 * [null,null,null,null,-3,null,0,-2]
 *
 * 解释：
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 *
 *
 *
 *
 * 提示：
 *
 *
 * pop、top 和 getMin 操作总是在 非空栈 上调用。
 *
 *
*/

// @lc code=start
type MinStack struct {
	s      []node
	length int
}

type node struct {
	val int
	min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		s:      make([]node, 0, 20),
		length: 0,
	}
}

func (this *MinStack) Push(x int) {
	if this.length > 0 {
		this.s = append(this.s, node{x, minInt(this.s[this.length-1].min, x)})
	} else {
		this.s = append(this.s, node{x, x})
	}
	this.length++
}

func (this *MinStack) Pop() {
	this.s = this.s[:this.length-1]
	this.length--
}

func (this *MinStack) Top() int {
	return this.s[this.length-1].val
}

func (this *MinStack) GetMin() int {
	return this.s[this.length-1].min
}
func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
