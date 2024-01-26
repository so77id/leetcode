
import (
	"fmt"
	"strconv"
	"strings"
)

func isDigit(n string) bool {
    _, err := strconv.Atoi(n)
    if err != nil {
        return false
    }
    return true
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func isLetter(l rune) bool {
    if l >= 'a' && l <= 'z' {
        return true
    }
    return false
}

func multiplyStr(n int, s []string) string {
    var builder strings.Builder
    for _, w := range s {
        builder.WriteString(w)
    }
    str := builder.String()
    builder.Reset()

    for i:=0; i<n; i++ {
        builder.WriteString(str)
    }
    return builder.String()
}

func decodeString(s string) string {
    stack := []string{}

    for _, c := range s {
        if c == ']' {
            str_l := []string{}
            for stack[len(stack)-1] != "[" {
                str_l = append([]string{stack[len(stack)-1]}, str_l...)
                stack = stack[:len(stack)-1]
            }
            // delete [
            stack = stack[:len(stack)-1]
            // now is number
            num_l := []int{}
            for len(stack) > 0 && isDigit(stack[len(stack)-1]) {
                v, _ := strconv.Atoi(stack[len(stack)-1])
                num_l = append([]int{v}, num_l...)
                stack = stack[:len(stack)-1]
            }
            
            num := 0
            for _, n := range num_l {
                num = num * 10 + n
            }
            //create new string
            str := multiplyStr(num, str_l)
            stack = append(stack, str)
        } else {
            stack = append(stack, string(c))
        }
    }

    var builder strings.Builder
    for _, w := range stack{
        builder.WriteString(w)
    }

    return builder.String()
}