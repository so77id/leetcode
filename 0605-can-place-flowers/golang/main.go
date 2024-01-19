func canPlaceFlowers(flowerbed []int, n int) bool {
    count := 0
    var left bool
    var right bool

    for i := 0 ; i < len(flowerbed); i++ {
        if flowerbed[i] == 0 {
            if i == 0 || flowerbed[i-1] == 0 {
                left = true
            } else {
                left = false
            }

            if i == len(flowerbed) -1 || flowerbed[i+1] == 0 {
                right = true
            } else {
                right = false
            }

            if left && right {
                count++
                flowerbed[i] = 1
                if count >= n {
                    return true
                }
            }
        }
    }
    return count >= n
}