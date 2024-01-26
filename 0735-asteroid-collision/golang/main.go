func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func asteroidCollision(asteroids []int) []int {
    stack := make([]int, 0)
    stack = append(stack, asteroids[0])
    
    for i:= 1; i<len(asteroids); i++ {
        if asteroids[i] > 0 {
            stack = append(stack, asteroids[i])
            continue
        }
        flag := true
        for len(stack) > 0 && stack[len(stack)-1] > 0 {
            if abs(stack[len(stack)-1]) == abs(asteroids[i]) {
                stack = stack[:len(stack)-1]
                flag = false
                break
            } 
            if abs(stack[len(stack)-1]) > abs(asteroids[i]) {
                flag = false
                break
            } 
            stack = stack[:len(stack)-1]

        }
        if flag {
            stack = append(stack, asteroids[i])
        }
    }

    return stack
}