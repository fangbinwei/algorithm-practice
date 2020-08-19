package main

/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/description/
 *
 * algorithms
 * Easy (74.26%)
 * Likes:    94
 * Dislikes: 0
 * Total Accepted:    32.5K
 * Total Submissions: 43.8K
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
// leetcode 一位老哥超级好用的模板
// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/solution/mo-fang-di-gui-zhi-bian-yi-xing-by-sonp/
func postorder(root *Node) []int {
	res := make([]int, 0, 10)
	if root == nil {
		return res
	}
	stack := make([]*Node, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if t != nil {
			stack = append(stack, t)
			stack = append(stack, nil)

			for i := len(t.Children) - 1; i >= 0; i-- {
				if c := t.Children[i]; c != nil {
					stack = append(stack, c)
				}
			}
		} else {
			res = append(res, stack[len(stack) - 1].Val)
			stack = stack[:len(stack)-1]
		}
	}
	return res
}

// @lc code=end

// 递归
func postorder(root *Node) []int {
	res := make([]int, 0, 10)
	if root == nil {
		return res
	}
	if len(root.Children) > 0 {
		for _, c := range root.Children {
			r := postorder(c)
			res = append(res, r...)
		}
	}
	res = append(res, root.Val)
	return res
}
