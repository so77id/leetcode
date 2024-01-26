
import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type bfsItem struct {
    root *TreeNode
    level int
}

func maxLevelSum(root *TreeNode) int {
    if root == nil {
        return 0
    }

    queue := []bfsItem{{
        root: root,
        level: 1,
    }}


    max := math.MinInt
    max_l := 1
    currLevel := 1
    sum := root.Val

    for len(queue) > 0 {
        curr, l := queue[0].root, queue[0].level
        queue = queue[1:]
        if currLevel != l {
            if  sum > max {
                max = sum
                max_l = currLevel
            }
            currLevel++
            sum = 0
        }

        sum += curr.Val

        if curr.Left != nil {
            queue = append(queue, bfsItem{
                root: curr.Left,
                level: l+1,
            })
        }
        if curr.Right != nil {
            queue = append(queue, bfsItem{
                root: curr.Right,
                level: l+1,
            })
        }
    }

    if sum > max {
        return currLevel
    }

    return max_l
}