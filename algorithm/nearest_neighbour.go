package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
)

// simple wrapper
func (this Graph) NearestNeighbourHamiltonCircleLength(start graphLib.VertexInterface) float64 {
    return NearestNeighbourHamiltonCircleLength(this, start)
}

// returns the length of the hamilton circle calculated by the nearest neighbour algorithm
func NearestNeighbourHamiltonCircleLength(graph Graph, start graphLib.VertexInterface) float64 {

    // keep track of which vertices we already visited
    visited := map[graphLib.VertexInterface]bool{start: true}

    // keep track of the total length
    length := 0.0

    // the current vertex
    vertex := start

    // (number of vertices) - 1 iterations are needed
    for n := graph.GetVertices().Count(); n > 1; n-- {

        // find the nearest neighbour, that is not visited yet
        weight, next := 0.0, graphLib.VertexInterface(nil)
        for _, edge := range vertex.GetEdges().All() {
            w, v := edge.GetWeight(), edge.GetOtherVertex(vertex)
            if (next == nil || w < weight) && !visited[v] {
                weight, next = w, v
            }
        }

        // add the distance to the nearest neighbour
        length += weight

        // mark the nearest neighbour visited
        visited[next] = true

        // continue with nearest neighbour
        vertex = next

    }

    // add the return path
    if w := graph.getWeightBetween(vertex, start); w >= 0.0 {
        return length + w
    }

    // if none was found
    return -1
}