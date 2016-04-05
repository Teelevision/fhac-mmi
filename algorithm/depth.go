package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
)

// simple wrapper
func (this Graph) DepthFirstSearch(start graphLib.VertexInterface) []graphLib.VertexInterface {
    return DepthFirstSearch(this, start)
}

// performs a depth-first search on the graph and returns a slice of vertices
func DepthFirstSearch(graph Graph, start graphLib.VertexInterface) []graphLib.VertexInterface {

    helper := depthFirstSearchHelper{
        graph: graph,
        result: []graphLib.VertexInterface{},
        discovered: map[graphLib.VertexInterface]bool{},
    }

    helper.search(start)

    return helper.result
}

// helper for the depth-first search
type depthFirstSearchHelper struct {
    graph      Graph

    // the result order
    result     []graphLib.VertexInterface

    // to keep track of discovered vertices
    discovered map[graphLib.VertexInterface]bool
}

// performs a depth-first search on the graph and returns a slice of vertices
func (this *depthFirstSearchHelper) search(vertex graphLib.VertexInterface) {

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
