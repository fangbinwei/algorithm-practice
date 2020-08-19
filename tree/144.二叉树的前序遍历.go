package main

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Medium (66.51%)
 * Likes:    347
 * Dislikes: 0
 * Total Accepted:    152.1K
 * Total Submissions: 228.6K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 前序 遍历。
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
 * 输出: [1,2,3]
 *
 *
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 *
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 遍历
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 10)
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0, 10)

	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			res = append(res, cur.Val)

			stack = append(stack, cur)
			// 左节点压入栈
			cur = cur.Left
		}
		// 拿出栈中的节点, 开始遍历其右侧
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 下面的代码可要可不要, 可以利用外层循环做下面代码一样的事情
		// for node.Right == nil && len(stack) > 0 {
		// 	node = stack[len(stack)-1]
		// 	stack = stack[:len(stack)-1]
		// }

		cur = node.Right
	}
	return res
}

// @lc code=end

func preorderTraversal(root *TreeNode) []int {

	res := make([]int, 0, 10)
	if root == nil {
		return nil
	}
	res = append(res, root.Val)
	if l := preorderTraversal(root.Left); l != nil {
		res = append(res, l...)
	}
	if r := preorderTraversal(root.Right); r != nil {
		res = append(res, r...)
	}
	return res
}

// 莫里斯遍历
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 10)
	if root == nil {
		return res
	}
	// 存在左子树
	cur := root
	for cur != nil {
		if left := cur.Left; left != nil {
			// 寻找最右侧节点
			for left.Right != nil {
				left = left.Right
			}
			left.Right = cur.Right
		} else {
			// 不存在左子树
			cur.Left = cur.Right
		}
		cur = cur.Left
	}

	for ; root != nil; root = root.Left {
		res = append(res, root.Val)
	}
	return res
}
