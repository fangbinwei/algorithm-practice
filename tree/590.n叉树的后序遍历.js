/*
 * @lc app=leetcode.cn id=590 lang=javascript
 *
 * [590] N叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/description/
 *
 * algorithms
 * Easy (74.88%)
 * Likes:    104
 * Dislikes: 0
 * Total Accepted:    37.9K
 * Total Submissions: 50.6K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的后序遍历。
 *
 * 例如，给定一个 3叉树 :
 *
 *
 *
 *
 *
 *
 *
 * 返回其后序遍历: [5,6,3,2,4,1].
 *
 *
 *
 * 说明: 递归法很简单，你可以使用迭代法完成此题吗?
 */

function Node(val, children) {
  this.val = val
  this.children = children
}

// @lc code=start
/**
 * // Definition for a Node.
 * function Node(val,children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */

/**
 * @param {Node} root
 * @return {number[]}
 */
// 后序: 左->右->根
// 栈内: 根->右->左
var postorder = function (root) {
  let ans = []
  if (root == null) {
    return ans
  }
  let stack = [root]
  while (stack.length > 0) {
    let cur = stack.pop()
    if (cur == null) {
      ans.push(stack.pop().val)

    } else {
      stack.push(cur)
      stack.push(null)
      if (cur.children) {
        for (let i=cur.children.length-1;i>=0;i--) {
          stack.push(cur.children[i])
        }
      }
    }
  }

  return ans
}
// @lc code=end
