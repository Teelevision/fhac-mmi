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
