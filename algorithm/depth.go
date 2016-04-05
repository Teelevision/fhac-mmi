package algorithm

import (
    "github.com/teelevision/fhac-mmi/graph"
)

// performs a depth-first search on the graph and returns a slice of vertices
func (this Graph) DepthFirstSearch(start graph.VertexInterface) []graph.VertexInterface {

    helper := depthFirstSearchHelper{
        graph: this,
        result: []graph.VertexInterface{},
        discovered: map[graph.VertexInterface]bool{},
    }

    helper.search(start)

    return helper.result
}

// helper for the depth-first search
type depthFirstSearchHelper struct {
    graph      Graph

    // the result order
    result     []graph.VertexInterface

    // to keep track of discovered vertices
    discovered map[graph.VertexInterface]bool
}

// performs a depth-first search on the graph and returns a slice of vertices
func (this *depthFirstSearchHelper) search(vertex graph.VertexInterface) {

    // discover vertex
    this.result = append(this.result, vertex)
    this.discovered[vertex] = true

    // visit all neighbours that were not discovered yet
    for _, v := range this.graph.getNeighboursOfVertex(vertex).All() {
        if !this.discovered[v] {
            this.search(v)
        }
    }
}
