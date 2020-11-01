package leetcode297test1

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	ans := make([]string, 0, 10)
	stack := []*TreeNode{}
	cur := root

	// 前序遍历, 入栈时记录
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			ans = append(ans, strconv.Itoa(cur.Val))
			cur = cur.Left
		}
		ans = append(ans, "#")
		cur = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	ans = append(ans, "#")
	return strings.Join(ans, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	ans := strings.Split(data, ",")

	i := 0
	v, _ := strconv.Atoi(ans[i])
	root := &TreeNode{v, nil, nil}
	stack := make([]*TreeNode, 0, 10)

	cur := root
	for cur != nil || len(stack) > 0 {
		i++
		if cur != nil {
			stack = append(stack, cur)
			cur.Left = getTreeNode(ans[i])
			cur = cur.Left
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			top.Right = getTreeNode(ans[i])
			cur = top.Right
		}
	}

	return root
}

func getTreeNode(s string) *TreeNode {
	if s == "#" {
		return nil
	}
	v, _ := strconv.Atoi(s)
	return &TreeNode{v, nil, nil}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
