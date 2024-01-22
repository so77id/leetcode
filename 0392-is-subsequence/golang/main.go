func isSubsequence(s string, t string) bool {
    i := 0 // index in S
    j := 0 // index in t
    sBound := len(s)
    tBound := len(t)

    for i < sBound && j < tBound {
        if s[i] == t[j] {
            i++
        } 
        j++
    }

    return i == sBound
}