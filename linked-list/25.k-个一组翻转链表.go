package main

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 *
 * https://leetcode-cn.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (62.45%)
 * Likes:    678
 * Dislikes: 0
 * Total Accepted:    92.3K
 * Total Submissions: 147.7K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
 *
 * k 是一个正整数，它的值小于或等于链表的长度。
 *
 * 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
 *
 *
 *
 * 示例：
 *
 * 给你这个链表：1->2->3->4->5
 *
 * 当 k = 2 时，应当返回: 2->1->4->3->5
 *
 * 当 k = 3 时，应当返回: 3->2->1->4->5
 *
 *
 *
 * 说明：
 *
 *
 * 你的算法只能使用常数的额外空间。
 * 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
 *
 *
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	reverseGroup := func(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
		current := head
		pre := tail.Next
		// 反转链表, head指向下一group的head, head成为tail
		for pre != tail {
			next := current.Next
			current.Next = pre
			pre = current
			current = next
		}
		return tail, head
	}

	guard := &ListNode{Next: head}
	tail := guard
	pre := guard
	for head != nil {
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return guard.Next
			}
		}
		// tail.Next已经正确, 需要将head和前group的tail链接
		head, tail = reverseGroup(head, tail)
		pre.Next = head // 第一次执行的时候改变了guard.Next
		pre = tail
		head = tail.Next
	}
	return guard.Next

}

// @lc code=end

// 递归 和24类似
func reverseKGroup(head *ListNode, k int) *ListNode {
	p := head
	for i := 0; i < k; i++ {
		if p != nil {
			p = p.Next
		} else {
			return head
		}
	}

	tmp := head.Next
	head.Next = reverseKGroup(p, k)

	prev := head
	for (tmp != nil) && (tmp != p) {
		next := tmp.Next
		tmp.Next = prev
		prev = tmp
		tmp = next
	}
	return prev
}
