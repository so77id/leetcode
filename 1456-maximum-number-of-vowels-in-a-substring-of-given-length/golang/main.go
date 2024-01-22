
func isVowel(v byte) bool {
    if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
        return true
    }

    return false
}

func maxVowels(s string, k int) int {
    i, j := 0, 0
    max, actual := 0, 0
    

    for ; j < k; j++{
        if isVowel(s[j]) {
            max++
        }
    }    
    actual = max
    for ; j < len(s); j++ {
        if isVowel(s[j]){
            actual++
        }
        if isVowel(s[i]) {
            actual--
        }
        if actual > max {
            max = actual
        }
        i++
    }

    return max
}