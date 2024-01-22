import "math"

func min(a , b int) int {
    if a > b {
        return a
    }
    return b
}

func maxArea(height []int) int {
    if len(height) == 1 {
        return 0
    }
    
    l, r := 0, len(height) - 1 
    max := -math.MaxInt

    for l < r {
        // get index of min value
        var i int
        var isLeft bool
        if height[l] > height[r] {
            i = r
            isLeft = false
        } else {
            i = l
            isLeft = true
        }
        area := height[i] * (r - l)

        if max < area {
            max = area
        }

        if isLeft {
            l++
        } else {
            r--
        }
    }

    return max
}