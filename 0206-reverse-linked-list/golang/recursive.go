/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    if head.Next == nil {
        return head
    }
    return helper(nil, head)
}

func helper(prev *ListNode, curr *ListNode) *ListNode {
    if curr == nil {
        return prev
    }
    // general case 
    // prev -> curr -> ...
    
    next := curr.Next
    curr.Next = prev
    
    head := helper(curr, next)
    
    return head
}