import "strings"

func mergeAlternately(word1 string, word2 string) string {
    var builder strings.Builder
    
    i := 0
    for ; i < len(word1) && i < len(word2); i++ {
        builder.WriteByte(word1[i])
        builder.WriteByte(word2[i])
    }

    if i < len(word1) {
        builder.WriteString(word1[i:len(word1)])
    } else if i < len(word2){
        builder.WriteString(word2[i:len(word2)])
    }

    return builder.String()
}