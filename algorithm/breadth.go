package algorithm

import (
    "github.com/teelevision/fhac-mmi/graph"
)

// performs a breadth-first search on the graph and returns a slice of vertices
func (this Graph) BreadthFirstSearch(start graph.VertexInterface) []graph.VertexInterface {

    // 1. add every discovered vertex to the result slice
    // 2. move along the result slice and check for undiscovered neighbours
    // 3. stop when end of result slice is reached

    // the result slice, already containing the start vertex
    result := []graph.VertexInterface{start}

    // to keep track of discovered vertices
    // the start vertex is already discovered
    discovered := map[graph.VertexInterface]bool{start: true}

    // go through the result
    for i := 0; i < len(result); i++ {
        vertex := result[i]

        // add all neighbours that were not discovered yet to the result
        for _, v := range this.getNeighboursOfVertex(vertex).All() {
            if !discovered[v] {
                discovered[v] = true
                result = append(result, v)
            }
        }
    }

    return result
}
