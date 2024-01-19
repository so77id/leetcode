import "math"

func increasingTriplet(nums []int) bool {
    f := math.MaxInt
    s := math.MaxInt

    for _, n := range nums {
        if n <= f {
            f = n
        } else if n <= s {
            s = n
        } else {
            return true
        }
    }

    return false
}