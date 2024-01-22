import "slices"

func maxOperations(nums []int, k int) int {
    count := 0
    l, r := 0, len(nums)-1
    slices.Sort(nums)

    for l < r {
        if nums[l] + nums[r] < k {
            l++
        } else if nums[l] + nums[r] > k {
            r--
        } else {
            count++
            l++
            r--
        }
    }
    return count
}