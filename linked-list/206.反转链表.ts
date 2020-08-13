/*
 * @lc app=leetcode.cn id=206 lang=typescript
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

class ListNode {
     val: number
     next: ListNode | null
     constructor(val?: number, next?: ListNode | null) {
         this.val = (val===undefined ? 0 : val)
         this.next = (next===undefined ? null : next)
     }
 }
// @lc code=start
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */
// 递归
function reverseList(head: ListNode | null): ListNode | null {
  if (head == null || head.next == null) return head
  // 1->2->3
  // head.next是2, current是3
  let current : ListNode | null =  reverseList(head.next)
  // 反转后, 2需要-> 1(head)
  head.next.next = head
  // 1需要->null
  head.next = null
  return current
};
// @lc code=end

// 空间复杂度O(1)
function reverseList(head: ListNode | null): ListNode | null {
  if (head == null) return null
  let previous: ListNode | null = null
  let current: ListNode | null  = head
  while (current) {
    let next: ListNode | null = current.next
    current.next = previous
    previous = current
    current = next
  }
  return previous
};



// 利用数组存储链表, 反转, 空间复杂度 O(n)
function reverseList(head: ListNode | null): ListNode | null {
  let bak: ListNode[] = []
  if (head == null) return null
  bak.push(head)
  let h: ListNode | null  = head
  while (h= h.next) {
    bak.push(h)
  }

  let p: ListNode = bak.pop() as ListNode
  let reverseList = p
  let next: ListNode| undefined
  while(next = bak.pop()) {
    p.next = next as ListNode
    p = next
  }
  p.next = null
  return reverseList 
};