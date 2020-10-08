/*
 * @lc app=leetcode.cn id=25 lang=javascript
 *
 * [25] K 个一组翻转链表
 *
 * https://leetcode-cn.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (63.21%)
 * Likes:    761
 * Dislikes: 0
 * Total Accepted:    106.9K
 * Total Submissions: 169.1K
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

 function ListNode(val, next) {
     this.val = (val===undefined ? 0 : val)
     this.next = (next===undefined ? null : next)
 }
// @lc code=start
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} k
 * @return {ListNode}
 */
var reverseKGroup = function(head, k) {
    
  if (head == null || head.next == null) return head

  let dummy = new ListNode(-1, head)

  let current = head
  let prev = dummy

  while (current!= null) {
    let tmp = current

    let enough = true
    for (let i=0; i < k-1; i++) {
      tmp = tmp.next
      if (tmp == null) {
        enough = false
        break
      }
    }

    if (enough) {
      let first = current
      let c = current
      let nextNode = c.next
      for (let i=0; i < k -1; i++) {
        let tmp = nextNode.next
        nextNode.next = c
        c = nextNode
        nextNode = tmp
      }
      prev.next = c

      first.next = nextNode
      prev = first
      current = nextNode
    } else {
      break
    }
  }
  return dummy.next
};
// @lc code=end

