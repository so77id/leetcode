
func longestSubarray(nums []int) int {
    i := 0
    ans := 0
    zeros := 1

    for j:= 0; j < len(nums); j++ {
        if nums[j] == 0 {
            zeros--
        }
        if zeros < 0 {
            for ;i<j && nums[i] != 0; i++ {}
            i++
            zeros++
        }
        if ans < j - i {
            ans = j-i
        }
    }

    return ans
}