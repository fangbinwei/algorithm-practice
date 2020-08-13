/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 *
 * https://leetcode-cn.com/problems/linked-list-cycle/description/
 *
 * algorithms
 * Easy (48.91%)
 * Likes:    711
 * Dislikes: 0
 * Total Accepted:    204.8K
 * Total Submissions: 418.5K
 * Testcase Example:  '[3,2,0,-4]\n1'
 *
 * 给定一个链表，判断链表中是否有环。
 *
 * 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
 *
 *
 *
 * 示例 1：
 *
 * 输入：head = [3,2,0,-4], pos = 1
 * 输出：true
 * 解释：链表中有一个环，其尾部连接到第二个节点。
 *
 *
 *
 *
 * 示例 2：
 *
 * 输入：head = [1,2], pos = 0
 * 输出：true
 * 解释：链表中有一个环，其尾部连接到第一个节点。
 *
 *
 *
 *
 * 示例 3：
 *
 * 输入：head = [1], pos = -1
 * 输出：false
 * 解释：链表中没有环。
 *
 *
 *
 *
 *
 *
 * 进阶：
 *
 * 你能用 O(1)（即，常量）内存解决此问题吗？
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
func hasCycle(head *ListNode) bool {
	// 快慢指针, 如果有环, 它们会相遇
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head

	for fast != nil {
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			return true
		}
	}
	return false
}

// @lc code=end

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	guard := head
	for current := head; current != nil; current = current.Next {
		if current.Next == nil {
			return false
		}
		if current.Next == current {
			return true
		}

		// 不断和前面的节点比对
		for ; guard != current; guard = guard.Next {
			if current.Next == guard {
				return true
			}
		}
		guard = head
	}
	return false
}