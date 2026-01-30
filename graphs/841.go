package main

func dfs(rooms [][]int, visited []bool, index int){
	visited[index] = true
	for i:=0;i<len(rooms[index]);i++{
		next:=rooms[index][i]
		if !visited[next]{
			dfs(rooms,visited,next)
		}
	}
}

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
    visited := make([]bool, n)
    dfs(rooms, visited, 0)
	for i := 0; i < n; i++ {
        if !visited[i] {
            return false
        }
    }
    return true
}