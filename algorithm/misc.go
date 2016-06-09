package algorithm

import (
    "github.com/teelevision/fhac-mmi/graph"
)

type Graph struct {
    *graph.Graph
}

// returns all or only the outgoing neighbours of the vertex depending on the graph
func (this Graph) getNeighboursOfVertex(vertex graph.VertexInterface) graph.VerticesInterface {
    if this.IsDirected() {
        return vertex.GetOutgoingNeighbours()
    }
    return vertex.GetNeighbours()
}

// returns all or only the outgoing edges of the vertex depending on the graph
func (this Graph) getEdgesOfVertex(vertex graph.VertexInterface) graph.EdgesInterface {
    if this.IsDirected() {
        return vertex.GetOutgoingEdges()
    }
    return vertex.GetEdges()
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

// returns the edge from the one to the other of the two given vertices
func (this Graph) getEdgeFromTo(v1, v2 graph.VertexInterface) graph.EdgeInterface {
    for _, edge := range v1.GetOutgoingEdges().All() {
        if edge.GetEndVertex().GetId() == v2.GetId() {
            return edge
        }
    }
    return nil
}

// a function that takes a start vertex and then traverses the graph
type TraverseFunction func(graph Graph, start graph.VertexInterface) []graph.VertexInterface

// a function that takes a start vertex and then returns the minimal spanning tree
type MinimalSpanningTreeFunction func(graph Graph, start graph.VertexInterface) (length float64, mst Graph, mapping map[graph.VertexInterface]graph.VertexInterface)