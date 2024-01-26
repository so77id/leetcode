/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func pairSum(head *ListNode) int {
    stack := []*ListNode{}
    tmp := head

    for ;tmp != nil; tmp = tmp.Next {
        stack = append(stack, tmp)
    }
    maxSum := 0
    curSum := 0
    tmp = head
    for i := 0; i < len(stack)/2; i++ {
        curSum = stack[len(stack)-i-1].Val + tmp.Val
        if curSum > maxSum {
            maxSum = curSum
        }
        tmp = tmp.Next
    }

    return maxSum
}