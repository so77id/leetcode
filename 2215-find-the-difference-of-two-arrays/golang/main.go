func findDifference(nums1 []int, nums2 []int) [][]int {
    s1 := make(map[int]struct{})
    s2 := make(map[int]struct{})
    ans1, ans2 := []int{}, []int{}

    for _, n := range nums1 {
        s1[n] = struct{}{}
    }
    
    for _, n := range nums2 {
        s2[n] = struct{}{}
    }

    for k,_ := range s1 {
        _, ok := s2[k]
        if !ok {
            ans1 = append(ans1, k)
        }
    }

    for k,_ := range s2 {
        _, ok := s1[k]
        if !ok {
            ans2 = append(ans2, k)
        }
    }

    return [][]int{ans1, ans2}
}