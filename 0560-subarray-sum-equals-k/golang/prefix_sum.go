func subarraySum(nums []int, k int) int {
    n := 0
    
    for j := 0; j < len(nums); j++ {
        suffixSum := 0
        for i := j; i < len(nums); i++ {
            suffixSum += nums[i]
            if suffixSum == k {
                n++
            }
        }
    }
    

    return n
}