
import (
	"fmt"
)

func isBowel(c rune) bool {
    if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
       c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
        return true
    }
    return false
}

func reverseVowels(s string) string {
    l := 0
    r := len(s)-1

    sr := []rune(s)
    
    for l < r {
        for ;isBowel(sr[l]) == false ; l++ {
            if l >= r {
                break
            }
        }
        for ;isBowel(sr[r]) == false && l <= r; r-- {
            if l >= r {
                break
            }
        }
        if l < r {
            tmp := sr[l]
            sr[l] = sr[r]
            sr[r] = tmp 
            l++
            r--
        }
    }
    
    return string(sr)
}