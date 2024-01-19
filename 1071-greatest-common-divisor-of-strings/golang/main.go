
import (
	"fmt"
	"strings"
)

// 0 1 2 3 4 5 -> len 6
// 0 1 2 3 4 -> len 5

func gcdOfStrings(str1 string, str2 string) string {
    if str1 == str2 {
        return str1
    }
    var ans string
    var half int
    
    if len(str1) > len(str2) {
        half = len(str1)/2
    } else {
        half = len(str2)/2
    }
    
    for i:= 0; i<half; i++ {
        if len(str1) % (i+1) != 0 || len(str2) % (i+1) != 0 {
            continue
        }
        
        s1 := strings.Repeat(str1[0:i+1], len(str1)/(i+1))
        s2 := strings.Repeat(str1[0:i+1], len(str2)/(i+1))
        
        if (s1 == str1 && s2 == str2) {
            ans = str1[0:i+1]
        }
    }

    return ans
}