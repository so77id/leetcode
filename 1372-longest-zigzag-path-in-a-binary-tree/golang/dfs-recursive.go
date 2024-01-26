/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func dfs(root *TreeNode, prev string, count int) int {
    if root == nil {
        return count
    }
    l, r := 0, 0 
    if prev == "l" {
        l = dfs(root.Left, "l", 0)
        r = dfs(root.Right, "r", count+1)
    } else {
        l = dfs(root.Left, "l", count+1)
        r = dfs(root.Right, "r", 0)
    }

    if l > r {
        return l
    }
    return r
}

func longestZigZag(root *TreeNode) int {
    l := dfs(root.Left, "l", 0)
    r := dfs(root.Right, "r", 0)

    if l > r {
        return l
    }

    return r
}