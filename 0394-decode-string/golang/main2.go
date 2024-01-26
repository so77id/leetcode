
import (
	"strings"
)

func isDigit(n rune) bool {
    if n >= '0' && n <= '9' {
        return true
    }
    return false
}


func isLetter(l rune) bool {
    if l >= 'a' && l <= 'z' {
        return true
    }
    return false
}

func decodeString(s string) string {
    countStack := []int{}
    strStack := []string{}
    k := 0
    var decodeString strings.Builder

    for _, c := range s {
        if isDigit(c) {
            k = 10 * k + int(c - '0')
        }
        if isLetter(c) {
            decodeString.WriteRune(c)
        }
        if c == '[' {
            countStack = append(countStack, k)
            strStack = append(strStack, decodeString.String())
            k = 0
            decodeString.Reset()
        }
        if c == ']' {
            currentK := countStack[len(countStack)-1]
            countStack = countStack[:len(countStack)-1]

            currentStr := strStack[len(strStack)-1]
            strStack = strStack[:len(strStack)-1]

            dstr := decodeString.String()
            decodeString.Reset()
            decodeString.WriteString(currentStr)
            for i:= 0; i < currentK; i++ {
                decodeString.WriteString(dstr)
            }
        }
    }

    return decodeString.String()
}