func string2freq(word string) map[rune]int {
    f :=  make(map[rune]int)
    for _, c := range(word) {
        f[c]++
    }
    return f
}

func invfreq(f map[rune]int) map[int]int {
    invf := make(map[int]int)
    for _, v := range(f) {
        invf[v]++
    }

    return invf
}

func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }

    f1 := string2freq(word1)
    f2 := string2freq(word2)

    for k, _ := range f1 {
        _, ok := f2[k]
        if !ok {
            return false
        }
    }

    for k, _ := range f2 {
        _, ok := f1[k]
        if !ok {
            return false
        }
    }

    invf1 := invfreq(f1)
    invf2 := invfreq(f2)
    
    for k, v := range invf1 {
        _, ok := invf2[k]
        if !ok || invf2[k] != v {
            return false
        }
    }

    return true
}