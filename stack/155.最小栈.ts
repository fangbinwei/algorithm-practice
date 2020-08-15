/*
 * @lc app=leetcode.cn id=155 lang=typescript
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
type InnerStructure = {
    val: number
    min: number
}
class MinStack {
    // 也可以用链表来实现
    stack: InnerStructure[]
    constructor() {
        this.stack = []
    }

    push(x: number): void {
        if (this.stack.length) {
            // min存储当前以及之前节点的最小值
            this.stack.push({
                val: x,
                min: Math.min(x, this.stack[this.stack.length-1].min)
            })
            return
        } 
        this.stack.push({
            val: x,
            min: x
        })
    }

    pop(): void {
        this.stack.pop()
    }

    top(): number {
        return this.stack[this.stack.length -1].val
    }

    getMin(): number {
        return this.stack.length
        ? this.stack[this.stack.length -1].min
        : 0
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * var obj = new MinStack()
 * obj.push(x)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */
// @lc code=end

