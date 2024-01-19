func productExceptSelf(nums []int) []int {
    prefix := make([]int, len(nums))
    sufix := make([]int, len(nums))
    ans := make([]int, len(nums))

    for i, _ := range nums {
        if i == 0 {
            prefix[i] = nums[i]
            sufix[len(nums)-1-i] = nums[len(nums)-1-i]
        } else {
            prefix[i] = prefix[i-1] * nums[i]
            sufix[len(nums)-1-i] = sufix[len(nums)-i] * nums[len(nums)-1-i]
        }
    }

    for i, _ := range nums {
        if i == 0 {
            // first element
            ans[i] = sufix[i+1]
        } else if i == len(nums)-1 {
            // last element
            ans[i] = prefix[i-1]
        } else {
            // general case
            ans[i] = prefix[i-1] * sufix[i+1]
        }
    }
    return ans
}