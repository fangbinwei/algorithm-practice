package leetcode_206

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 *
 * https://leetcode-cn.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (70.15%)
 * Likes:    1155
 * Dislikes: 0
 * Total Accepted:    301.6K
 * Total Submissions: 430K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 反转一个单链表。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL
 * 输出: 5->4->3->2->1->NULL
 *
 * 进阶:
 * 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
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
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		tmp := current.Next
		current.Next = prev
		prev = current

		current = tmp
	}
	return prev
}

// @lc code=end

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	reverseRemain := reverseList(head.Next)

	head.Next.Next = head

	head.Next = nil

	return reverseRemain
}