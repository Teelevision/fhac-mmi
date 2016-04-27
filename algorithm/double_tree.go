package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
)

// simple wrapper
func (this Graph) DoubleTreeHamiltonCircleLength(start graphLib.VertexInterface) float64 {
    return DoubleTreeHamiltonCircleLength(this, PrimLength, BreadthFirstSearch, start)
}

// returns the length of the hamilton circle calculated by the double tree algorithm
func DoubleTreeHamiltonCircleLength(graph Graph, mst MinimalSpanningTreeFunction, traverse TraverseFunction, start graphLib.VertexInterface) float64 {

    // get the minimal spanning tree
    _, graph = mst(graph, start)

    // traverse
    result := traverse(graph, start)

    // sum up the weight
    length := 0.0
    for i := len(result) - 1; i > 0; i-- {
        if w := graph.getWeightBetween(result[i], result[i - 1]); w >= 0.0 {
            length += w
        } else {
            panic("weight not found")
        }
    }

    return length
}