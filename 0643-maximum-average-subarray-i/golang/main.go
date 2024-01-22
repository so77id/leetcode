func findMaxAverage(nums []int, k int) float64 {
    i, j := 0, 0
    max := 0
    sum := 0
    for ; j < k; j++ {
        max += nums[j]
    }
    sum = max
    for ; j<len(nums); j++ {
        sum += nums[j] - nums[i]
        if sum > max {
            max = sum
        }
        i++
    }

    return float64(max)/float64(k)
}