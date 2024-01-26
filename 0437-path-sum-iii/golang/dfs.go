/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func nSuffixpath(nums []int, targetSum int) int {
    n := 0
    suffixSum := 0

    for i := 0; i < len(nums); i++ {
        suffixSum += nums[i]
        if suffixSum == targetSum {
            n++
        }
    }

    return n
}

type dfsItem struct {
    path []int
    root *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
    if root == nil {
        return 0
    }

    stack := []dfsItem{{
        path: []int{root.Val},
        root: root,
    }}

    ans := 0

    for len(stack) > 0 {
        // top and pop
        curr := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        // get n suffix paths with targetSum value
        ans += nSuffixpath(curr.path, targetSum)

        // append childs
        if curr.root.Left != nil {
            stack = append(stack, dfsItem{
                path: append([]int{curr.root.Left.Val}, curr.path...),
                root: curr.root.Left,
            })
        }
        if curr.root.Right != nil {
            stack = append(stack, dfsItem{
                path: append([]int{curr.root.Right.Val}, curr.path...),
                root: curr.root.Right,
            })
        }
    }

    return ans
}