/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    if head == nil  {
        return head
    }

    var prev, tmp, next *ListNode
    prev = nil
    
    for tmp = head ; tmp.Next != nil; {
        next = tmp.Next 
        tmp.Next = prev
        prev = tmp
        tmp = next
    }

    tmp.Next = prev

    return tmp
}