package main

/*
 * @lc app=leetcode.cn id=94 lang=golang
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
//遍历, 莫里斯, morris traversal
func inorderTraversal(root *TreeNode) []int {
	// 如果current节点有左子树, 令节点成为最右侧节点的子节点
	// (结合中序遍历来理解 左子树遍历完需要轮到current节点, 左子树最后遍历的节点是最右侧的节点 )
	// current节点没有左子树, 则输出current节点, 然后进入右子树
	res := make([]int, 0, 10)

	cur := root
	for cur != nil {
		if left := cur.Left; left != nil {
			tmp := left
			for left.Right != nil {
				left = left.Right
			}
			left.Right = cur
			// 避免死循环
			cur.Left = nil
			cur = tmp
		} else {
			// 没有左子树了, 输出
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}

// @lc code=end

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 10)
	stack := make([]*TreeNode, 0, 10)

	cur := root
	for cur != nil || len(stack) > 0 {
		// 对于一个节点, 先找左子节点
		for ; cur != nil; cur = cur.Left {
			stack = append(stack, cur)
		}
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, pop.Val)
		// 左子节点处理好后, 处理右子节点 (然后用外层循环处理这个节点的左子节点...)
		// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/er-cha-shu-de-zhong-xu-bian-li-by-leetcode/
		if pop.Right != nil {
			cur = pop.Right
		} else {
			cur = nil
		}
	}
	return res
}

// 递归
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 10)
	if root == nil {
		return nil
	}

	if l := inorderTraversal(root.Left); l != nil {
		res = append(res, l...)
	}
	res = append(res, root.Val)
	if r := inorderTraversal(root.Right); r != nil {
		res = append(res, r...)
	}
	return res

}
