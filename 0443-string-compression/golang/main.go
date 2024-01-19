import "strconv"
func compress(chars []byte) int {
    if len(chars) == 0 {
        return 0
    }
    if len(chars) == 1 {
        return 1
    }

    prev := chars[0]
    count := 1
    i := 1

    for _, c := range(chars[1:]) {
        if prev == c {
            count++
        } else {
            if count > 1 {
                nums := strconv.Itoa(count)
                for j := 0; j<len(nums); j++ {
                    chars[i] = nums[j]
                    i++
                }
            }
            chars[i] = c
            i++
            prev = c
            count = 1
        }
    }

    if count > 1 {
        nums := strconv.Itoa(count)
        for j := 0; j<len(nums); j++ {
            chars[i] = nums[j]
            i++
        }
    }

    return i
}