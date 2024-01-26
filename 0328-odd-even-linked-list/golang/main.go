/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    var eListHead, eListTail  *ListNode
    eListHead, eListTail = nil, nil
    oListHead, oListTail := head, head
    tmp := head

    for i := 2;tmp.Next != nil; i++ {
        if i % 2 == 0 {
            //even
            if eListHead == nil {
                eListHead, eListTail = tmp.Next, tmp.Next 
            } else {
                eListTail.Next = tmp.Next
                eListTail = eListTail.Next
            }  
        } else {
            //odd
            oListTail.Next = tmp.Next
            oListTail = oListTail.Next
        }
        tmp = tmp.Next
    }
    
    oListTail.Next = eListHead
    eListTail.Next = nil
    
    
    return oListHead
}