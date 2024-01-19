func twoSum(nums []int, target int) []int {
    

    m := make(map[int]int)

    for i, n := range nums {
        m[n] = i
    }

    ans := make([]int, 0)
    for i, n := range nums {
        val, ok := m[target-n]

        if ok && val != i {
            ans = append(ans, i)
            ans = append(ans, val)
            break
        }
    }

    return ans
}