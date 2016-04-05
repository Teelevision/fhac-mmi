package algorithm

import (
    "github.com/teelevision/fhac-mmi/graph"
    "fmt"
)

// calculates the number of connected components
func (this Graph) GetNumConnectedComponents() uint {

    // 1. pick a vertex that was not visited yet
    // 2. visit all vertices that are somehow connected
    // 3. if some vertices are not visited, continue with 1.

    numVertices := this.GetVertices().Count()
    numComponents := uint(0)

    // to keep track of which vertices were visited yet
    visited := map[graph.VertexInterface]bool{}

    // loop all vertices
    for _, vertex := range this.GetVertices().All() {

        // every vertex that was not visited yet is a good starting point
        if !visited[vertex] {

            // there is at least the one vertex, so it is a new connected component
            numComponents++

            // actual search
            result := this.Breitensuche(vertex)

            // early result
            // the number of newly visited vertices completes the search
            if uint(len(visited) + len(result)) == numVertices {
                fmt.Println("early")
                return numComponents
            }

            // keep track of which vertices were visited
            for _, v := range result {
                visited[v] = true
            }

        }
    }

    return numComponents
}