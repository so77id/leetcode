func canVisitAllRooms(rooms [][]int) bool {
    visited := make([]bool, len(rooms))
    for i := range visited {
        visited[i] = false
    }

    var queue[] int;
    queue = append(queue, 0)

    for len(queue) > 0 {
        src := queue[0]
        queue = queue[1:]

        visited[src] = true

        for _, room := range rooms[src] {
            if visited[room] == false {
                queue = append(queue, room)
            }
        }
    }

    for _, room := range visited {
        if room == false {
            return false
        }
    }

    return true

}