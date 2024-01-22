func moveZeroes(nums []int)  {
    i := 0
    j := 0

    for j < len(nums) {
        for ; i < len(nums) && nums[i] != 0; i++ {}
        for j=i+1; j < len(nums) && nums[j] == 0; j++ {}

        if i >= len(nums) || j >= len(nums) {
            break
        }

        tmp := nums[i]
        nums[i] = nums[j]
        nums[j] = tmp
    }

    return
}