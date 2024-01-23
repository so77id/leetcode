func uniqueOccurrences(arr []int) bool {
    counter := make(map[int]int)
    inv_counter := make(map[int]struct{})
    for _, n := range arr {
        counter[n]++
    }

    for _,v := range counter {
        _, ok := inv_counter[v]
        if ok {
            return false
        }
        inv_counter[v] = struct{}{}
    }

    return true
}