package leetcode297test3

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

func serialize(root *TreeNode, s *[]string) {

	if root == nil {
		*s = append(*s, "#")
		return
	} else {
		*s = append(*s, strconv.Itoa(root.Val))
		serialize(root.Left, s)
		serialize(root.Right, s)
	}

}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	ans := make([]string, 0, 10)
	serialize(root, &ans)

	return strings.Join(ans, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	ans := strings.Split(data, ",")

	var i = new(int)
	*i = 0
	return deserialize(&ans, i)
}

func deserialize( ans *[]string, i *int) *TreeNode {
	root := getTreeNode((*ans)[*i])
	if root != nil {
		*i++
		root.Left = deserialize(ans, i)
		*i++
		root.Right = deserialize(ans, i)
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
