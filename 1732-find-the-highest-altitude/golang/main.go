

func largestAltitude(gain []int) int {
    max := 0
    acc := 0
    for _, g := range gain {
        acc += g
        if max < acc {
            max = acc
        }
    }

    return max
}