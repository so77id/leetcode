
import "fmt"
type coords struct {
    x,y int 
}

func orangesRotting(grid [][]int) int {
    /*
        Haremos un bfs por turnos, cada nivel sera recorrido completo antes de comenzaar el otro.
        Contaremos la cantidad de niveles que se necesitan para que esten todos visitados o que no podamos visitar mas.
        si todas las narajas fueron infectadas retornamos el numero.
        si no todas fueron infectadas retornamos -1
    */
    visited := make(map[string]struct{})

    moves := []coords{
        {
            x: -1, y: 0,
        },
        {
            x: 0, y: -1,
        },
        {
            x: 1, y: 0,
        },
        {
            x: 0, y: 1,
        },
    }

    currQ, nextQ := []coords{}, []coords{};

    oranges := 0
    for x, row := range grid {
        for y, orange := range row {
            if orange == 1 {
                oranges++
            }
            if orange == 2 {
                oranges++
                currQ = append(currQ, coords{
                    x: x,
                    y: y,
                })
                key := fmt.Sprintf("%d-%d", x, y)

                visited[key] = struct{}{}
            }
        }
    }


    counter := 0
    for len(currQ) > 0 {
        for len(currQ) > 0 {
            curr := currQ[0]
            currQ = currQ[1:]

        
            for _, move := range moves {
                nextX := curr.x + move.x
                nextY := curr.y + move.y

                if  nextX < 0 || nextX >= len(grid) ||
                    nextY < 0 || nextY >= len(grid[0]) || 
                    grid[nextX][nextY] == 0 || 
                    grid[nextX][nextY] == 2  {
                        continue
                }
                key := fmt.Sprintf("%d-%d", nextX, nextY)
                _, ok := visited[key]
                if ok {
                    continue
                }

                visited[key] = struct{}{}

                nextQ = append(nextQ, coords{
                    x: nextX,
                    y: nextY,
                })
                
            }
        }
        if len(nextQ) > 0 {
            counter++;
        }   

        currQ = nextQ
        nextQ = []coords{}
    } 

    if len(visited) == oranges {
        return counter
    }

    return -1;
}