package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
)

// simple wrapper
func (this Graph) DoubleTreeHamiltonCircleLength(mst MinimalSpanningTreeFunction, traverse TraverseFunction, start graphLib.VertexInterface) float64 {
    return DoubleTreeHamiltonCircleLength(this, mst, traverse, start)
}

// returns the length of the hamilton circle calculated by the double tree algorithm
func DoubleTreeHamiltonCircleLength(graph Graph, mst MinimalSpanningTreeFunction, traverse TraverseFunction, start graphLib.VertexInterface) float64 {

    // get the minimal spanning tree
    _, mstGraph, vMap := mst(graph, start)

    // traverse
    result := traverse(mstGraph, mstGraph.GetVertices().Get(0))

    // reverse mapping
    mapping := make(map[graphLib.VertexInterface]graphLib.VertexInterface, len(vMap))
    for a, b := range vMap {
        mapping[b] = a
    }

    // sum up the weight
    length := 0.0
    for i := len(result) - 1; i > 0; i-- {
        if w := graph.getWeightBetween(mapping[result[i]], mapping[result[i - 1]]); w >= 0.0 {
            length += w
        } else {
            panic("weight not found")
        }
    }

    // return to start
    length += graph.getWeightBetween(mapping[result[0]], mapping[result[len(result)-1]])

    return length
}