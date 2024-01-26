/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func getLeafs(root *TreeNode) []int {
    //using dfs
    leafs := []int{}

    stack := []*TreeNode{root}


    for len(stack) > 0 {
        // top and pop
        curr := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        if curr.Left == nil && curr.Right == nil {
            // is leaf
            leafs = append(leafs, curr.Val)
        } else {
            if curr.Left != nil {
                stack = append(stack, curr.Left)
            }
            if curr.Right != nil {
                stack = append(stack, curr.Right)
            }
        }
    }

    return leafs
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    l1 := getLeafs(root1)
    l2 := getLeafs(root2)

    if len(l1) != len(l2) {
        return false
    }
    ans := true
    for i := 0; i < len(l1); i++ {
        if l1[i] != l2[i] {
            ans = false
        }
    }

    return ans
}