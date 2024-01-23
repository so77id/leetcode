import "strings"
type Stack []rune

func removeStars(s string) string {
    var st Stack

    for _, c := range s {
        if c == '*' {
            st = st[:len(st)-1]
        } else {
            st = append(st, c)
        }
    }

    var builder strings.Builder
    for _, c := range st {
        builder.WriteRune(c)
    }

    return builder.String()
}