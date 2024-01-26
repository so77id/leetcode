

import "math"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type dfsItem struct {
    max int
    tmp *TreeNode
}

func goodNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }

    stack := []dfsItem{
        dfsItem{
            max: -math.MaxInt,
            tmp: root,
        },
    }

    ans := 0

    for len(stack) > 0 {
        // top and pop
        curr := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        // check if curr > max 
        // ans++ and update max
        if curr.tmp.Val >= curr.max {
            ans++
            curr.max = curr.tmp.Val
        }

        // add all childs
        if curr.tmp.Left != nil {
            stack = append(stack, dfsItem{
                max: curr.max,
                tmp: curr.tmp.Left,
            })
        }
        if curr.tmp.Right != nil {
            stack = append(stack, dfsItem{
                max: curr.max,
                tmp: curr.tmp.Right,
            })
        }
    }

    return ans
}