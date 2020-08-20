package main

/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N叉树的层序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (66.23%)
 * Likes:    106
 * Dislikes: 0
 * Total Accepted:    27.1K
 * Total Submissions: 40.9K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
 *
 * 例如，给定一个 3叉树 :
 *
 *
 *
 *
 *
 *
 *
 * 返回其层序遍历:
 *
 * [
 * ⁠    [1],
 * ⁠    [3,2,4],
 * ⁠    [5,6]
 * ]
 *
 *
 *
 *
 * 说明:
 *
 *
 * 树的深度不会超过 1000。
 * 树的节点总数不会超过 5000。
 *
 */

type Node struct {
	Val      int
	Children []*Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	res := make([][]int, 0, 10)
	if root == nil {
		return res
	}
	// init
	queue := make([]*Node, 0)
	queue = append(queue, root)
	// 用一个nil做标记, 或者使用queue的长度
	queue = append(queue, nil)

	level := 0
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		if front != nil {
			if len(res) < level+1 {
				res = append(res, []int{})
			}
			res[level] = append(res[level], front.Val)

			for i, l := 0, len(front.Children); i < l; i++ {
				queue = append(queue, front.Children[i])
			}

			// 读到nil, 说明一行读完了
		} else {
			level++
			if len(queue) > 0 {
				queue = append(queue, nil)
			}
		}
	}
	return res

}

// @lc code=end

func levelOrder(root *Node) [][]int {
	res := make([][]int, 0, 10)
	if root == nil {
		return res
	}
	res = append(res, []int{root.Val})

	for _, c := range root.Children {
		for index, nodes := range levelOrder(c) {
			if len(res) < index+2 {
				res = append(res, []int{})
			}
			res[index+1] = append(res[index+1], nodes...)
		}
	}
	return res

}
