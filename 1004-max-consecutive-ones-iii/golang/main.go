func longestOnes(nums []int, k int) int {
    i, j := 0, 0
    ans := 0
    zr := k // zeros remain
    
    for j = 0; j < len(nums); j++ {
        if nums[j] == 0 {
            if zr == 0 {
                for ;i < j && nums[i] != 0;i++ {}
                i++
            } else {
                zr--
            }
        }
        
        if ans < j-i+1 {
            ans = j-i+1
        }
    }

    return ans
}