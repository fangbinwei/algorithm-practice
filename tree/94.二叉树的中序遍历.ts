/*
 * @lc app=leetcode.cn id=94 lang=typescript
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (72.52%)
 * Likes:    632
 * Dislikes: 0
 * Total Accepted:    216.7K
 * Total Submissions: 298.9K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的中序 遍历。
 *
 * 示例:
 *
 * 输入: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * 输出: [1,3,2]
 *
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 *
 */

class TreeNode {
  val: number
  left: TreeNode | null
  right: TreeNode | null
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = val === undefined ? 0 : val
    this.left = left === undefined ? null : left
    this.right = right === undefined ? null : right
  }
}
// @lc code=start
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */
// 遍历
// 中序遍历是左->根->右
// 按照右->根->左的顺序入栈
function inorderTraversal(root: TreeNode | null): number[] {
  let ans: number[] = []
  if (root == null) {
    return ans
  }
  let stack: (TreeNode | null)[] = [root]
  while (stack.length > 0) {
    let cur = stack.pop()
    if (cur == null) {
      ans.push((stack.pop() as TreeNode).val)
    } else {
      if (cur.right) {
        stack.push(cur.right)
      }
      stack.push(cur)
      stack.push(null)
      if (cur.left) {
        stack.push(cur.left)
      }
    }
  }
  return ans
}
// @lc code=end

// 递归
function inorderTraversal(root: TreeNode | null): number[] {
  if (root == null) {
    return []
  }
  // let left = inorderTraversal(root.left)
  // let right = inorderTraversal(root.right)
  return [
    ...inorderTraversal(root.left),
    root.val,
    ...inorderTraversal(root.right),
  ]
}
