package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
)

// simple wrapper
func (this Graph) KruskalLength() float64 {
    return KruskalLength(this)
}

type kruskalHelperItem struct {
    vertices []graphLib.VertexInterface
}

// kruskal algorithm with result length
func KruskalLength(graph Graph) float64 {

    // the number of edges and vertices
    num := graph.GetVertices().Count()
    requiredNumEdges := num - 1

    components := make(map[graphLib.VertexInterface]*kruskalHelperItem, num)

    for _, vertex := range graph.GetVertices().All() {
        components[vertex] = &kruskalHelperItem{[]graphLib.VertexInterface{vertex}}
    }

    q := graphLib.NewCheapestEdgeQueue(graph.GetEdges().Count())
    for _, edge := range graph.GetEdges().All() {
        q.PushEdge(edge)
    }

    // keep track of the length of the minimal spanning tree
    length := 0.0

    for n, edge := uint(0), q.PopCheapestEdge(); edge != nil; edge = q.PopCheapestEdge() {

        cStart, cEnd := components[edge.GetStartVertex()], components[edge.GetEndVertex()]

        // compare pointers, if different, then those are different components
        if cStart != cEnd {

            // merge the smaller one into the bigger one
            if len(cStart.vertices) > len(cEnd.vertices) {
                cStart.vertices = append(cStart.vertices, cEnd.vertices...)
                for _, e := range cEnd.vertices {
                    components[e] = cStart
                }
            } else {
                cEnd.vertices = append(cEnd.vertices, cStart.vertices...)
                for _, e := range cStart.vertices {
                    components[e] = cEnd
                }
            }

            // add length
            length += edge.GetWeight()
            n++

            // if we got numVertices - 1 edges, we are finished
            if n >= requiredNumEdges {
                return length
            }
        }
    }

    // return the length/weight of the minimal spanning tree
    return length
}