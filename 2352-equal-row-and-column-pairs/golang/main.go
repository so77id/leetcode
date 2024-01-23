import (
    "strings"
    "strconv"
)

func row2string(grid [][]int, i int) string {
    var builder strings.Builder
    for j := 0; j < len(grid); j++ {
        if j != 0 {
            builder.WriteString("-")    
        }
        builder.WriteString(strconv.Itoa(grid[i][j]))
    }

    return builder.String()
}

func cols2string(grid [][]int, i int) string {
    var builder strings.Builder
    for j := 0; j < len(grid); j++ {
        if j != 0 {
            builder.WriteString("-")    
        }
        builder.WriteString(strconv.Itoa(grid[j][i]))
    }

    return builder.String()
}


func equalPairs(grid [][]int) int {
    rowsMap := make(map[string]int)
    count := 0

    for i:=0; i<len(grid); i++ {
        rowsMap[row2string(grid, i)]++
    }
    for i:=0; i<len(grid[0]); i++ {
        v, ok := rowsMap[cols2string(grid, i)]
        if ok {
            count += v
        }
    }

    return count
}
