package algorithm

import (
    "github.com/teelevision/fhac-mmi/graph"
)

type Graph struct {
    graph.GraphInterface
}

// returns all or only the outgoing neighbours of the vertex depending on the graph
func (this Graph) getNeighboursOfVertex(vertex graph.VertexInterface) graph.VerticesInterface {
    if this.IsDirected() {
        return vertex.GetOutgoingNeighbours()
    }
    return vertex.GetNeighbours()
}

// returns the weight between the two given vertices
func (this Graph) getWeightBetween(v1, v2 graph.VertexInterface) float64 {
    for _, side := range v1.GetEdgesFast() {
        for _, edge := range side() {
            if edge.GetOtherVertex(v1) == v2 {
                return edge.GetWeight()
            }
        }
    }
    return -1
}

// a function that takes a start vertex and then traverses the graph
type TraverseFunction func(graph Graph, start graph.VertexInterface) []graph.VertexInterface

// a function that takes a start vertex and then returns the minimal spanning tree
type MinimalSpanningTreeFunction func(graph Graph, start graph.VertexInterface) (length float64, mst Graph)