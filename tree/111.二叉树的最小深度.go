package main

import "container/ring"

/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (44.27%)
 * Likes:    354
 * Dislikes: 0
 * Total Accepted:    129.3K
 * Total Submissions: 292K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 *
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 *
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回它的最小深度  2.
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if root.Left != nil && root.Right != nil {
		if left < right {
			return 1 + left
		} else {
			return 1 + right
		}
	} else {
		// 哪个节点为nil, 哪个的深度就是0
		return 1 + left + right
	}

}

// @lc code=end

// 层序遍历/广度优先  首次遇到的 没有左右子节点的节点, 是最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0, 10)

	deep := 0
	queue = append(queue, root)
	// 可以直接用1个for循环代替
	for len(queue) > 0 {
		for _ = range queue {
			cur := queue[0]
			queue = queue[1:len(queue)]

			if cur.Left == nil && cur.Right == nil {
				return deep + 1
			}

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		deep++
	}

	return deep

}
