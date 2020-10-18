/*
 * @lc app=leetcode.cn id=589 lang=javascript
 *
 * [589] N叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/
 *
 * algorithms
 * Easy (74.10%)
 * Likes:    108
 * Dislikes: 0
 * Total Accepted:    49.8K
 * Total Submissions: 67.2K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的前序遍历。
 * 
 * 例如，给定一个 3叉树 :
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 返回其前序遍历: [1,3,5,6,2,4]。
 * 
 * 
 * 
 * 说明: 递归法很简单，你可以使用迭代法完成此题吗?
 */

// @lc code=start
/**
 * // Definition for a Node.
 * function Node(val, children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */

/**
 * @param {Node} root
 * @return {number[]}
 */
var preorder = function(root) {
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
      if (cur.children) {
        for (let i=cur.children.length-1;i>=0;i--) {
          stack.push(cur.children[i])
        }
      }
      stack.push(cur)
      stack.push(null)
    }
  }

  return ans
};
// @lc code=end

