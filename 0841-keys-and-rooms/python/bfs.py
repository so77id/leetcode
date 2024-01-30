class Solution:
    def canVisitAllRooms(self, rooms: List[List[int]]) -> bool:
        visited = [False for _ in range(len(rooms))]

        queue = []
        queue.append(0)

        while len(queue) >  0:
            src = queue[0]
            queue = queue[1:]

            visited[src] = True

            for room in rooms[src]:
                if visited[room] == False:
                    queue.append(room)

        for v in visited:
            if v == False:
                return False

        return True