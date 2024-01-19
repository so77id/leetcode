import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
    max := slices.Max(candies)

    ans := make([]bool, len(candies))

    for i := 0; i<len(candies); i++ {
        if candies[i] + extraCandies >= max {
            ans[i] = true
        }
    }

    return ans
}