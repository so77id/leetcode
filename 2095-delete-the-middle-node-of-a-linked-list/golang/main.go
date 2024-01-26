/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteMiddle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }

    count, tmp := 0, head
    for ; tmp != nil; tmp = tmp.Next {
        count++
    }
    
    half := count/2
    tmp = head
    for i := 0; i < half-1; i++ {
        tmp = tmp.Next
    }
    
    tmp.Next = tmp.Next.Next

    return head
}