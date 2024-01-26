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


func rightSideView(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    queue := []bfsItem{{
        root: root,
        level: 0,
    }}
    curLevel := 0
    ans := []int{}
    last := 0
    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]

        if curLevel != curr.level {
            curLevel++
            ans = append(ans, last)
        }
        last = curr.root.Val

        if curr.root.Left != nil {
            queue = append(queue, bfsItem{
                root: curr.root.Left,
                level: curr.level + 1,
            })
        }
        if curr.root.Right != nil {
            queue = append(queue, bfsItem{
                root: curr.root.Right,
                level: curr.level + 1,
            })
        }
    }

    if len(ans) == curLevel {
        ans = append(ans, last)
    }

    return ans
}