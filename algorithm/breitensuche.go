package algorithm

import (
    "github.com/teelevision/fhac-mmi/graph"
)

type Graph struct {
    graph.GraphInterface
}

func (this Graph) Breitensuche(start graph.VertexInterface) []graph.VertexInterface {

    result := []graph.VertexInterface{start}
    checked := map[graph.VertexInterface]bool{start: true}

    for i := 0; i < len(result); i++ {
        vertex := result[i]
        var neighbours graph.VerticesInterface
        if this.IsDirected() {
            neighbours = vertex.GetOutgoingNeighbours()
        } else {
            neighbours = vertex.GetNeighbours()
        }
        for _, v := range neighbours.All() {
            if !checked[v] {
                checked[v] = true
                result = append(result, v)
            }
        }
    }

    return result
}
