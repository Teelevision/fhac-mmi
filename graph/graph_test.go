package graph

import (
    "testing"
)

// tests the graph methods
func TestGraph(t *testing.T) {

    g := UndirectedGraph()

    // test NewVertex()
    var v [10]VertexInterface
    for i := 0; i < 10; i++ {
        v[i] = g.NewVertex()
    }

    // test NewEdge()
    var e [20]EdgeInterface
    for i := 0; i < 10; i++ {
        e[i] = g.NewEdge(v[i], v[(i + 1) % 10])
    }

    // test NewWeightedEdge()
    for i := 0; i < 10; i++ {
        e[i + 10] = g.NewWeightedEdge(v[i], v[(i + 5) % 10], float64(i + 1))
    }

    // test GetVertices()
    if g.GetVertices().Count() != 10 {
        t.Errorf("Graph should have 10 vertices, got %d.", g.GetVertices().Count())
    }

    // test GetEdges()
    found := 0
    for _, edge := range g.GetEdges().All() {
        for _, edge2 := range e {
            if edge == edge2 {
                found++
            }
        }
    }
    if g.GetEdges().Count() != 20 || found != 20 {
        t.Errorf("Graph should have 20 edges, got %d of which %d match.", g.GetEdges().Count(), found)
    }

}
