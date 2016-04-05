package parser

import (
    "testing"
    "reflect"
    graphLib "github.com/teelevision/fhac-mmi/graph"
)

// test parsing Graph2.txt
func TestParseEdges(t *testing.T) {

    graph, err := ParseEdgesFile("test/Graph2.txt", false)
    if err != nil {
        panic(err)
    }

    graphValidator(t, graph, true, 15, 17)

    // test each vertex
    for i, v := range graph.GetVertices().All() {
        num := []uint{3, 3, 3, 3, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0}
        if n := v.GetOutgoingNeighbours().Count(); n != num[i] {
            t.Errorf("Expected vertex #%d to have %d outgoing neighbours, got %d.", i, num[i], n)
        }
    }

    // for the first vertex: test if outgoing neighbours are right
    vertices := graph.GetVertices()
    v := vertices.Get(0)
    // neighbours are: 6, 9 and 13
    neighbours := []graphLib.VertexInterface{vertices.Get(6), vertices.Get(9), vertices.Get(13)}
    if ne := v.GetOutgoingNeighbours().All(); !reflect.DeepEqual(ne, neighbours) {
        t.Errorf("Expected %v, got $v.", neighbours, ne)
    }
}

// test parsing Graph2_weights.txt
func TestParseEdgesWithWeights(t *testing.T) {

    graph, err := ParseEdgesFile("test/Graph2_weights.txt", true)
    if err != nil {
        panic(err)
    }

    graphValidator(t, graph, true, 15, 17)

    // test weights
    for i, e := range graph.GetEdges().All() {
        j := float64(i + 1)
        expectWeight := j / 100 + j
        if w := e.GetWeight(); w != expectWeight {
            t.Errorf("Expected edge #%d to have weight %f, got %f.", i, expectWeight, w)
        }
    }
}

// test failing to parse Graph2_fail.txt
// In this graph last edge is incomplete (only start, no end).
func TestParseEdgesFail(t *testing.T) {
    expectError := "EOF"
    if _, err := ParseAdjacencyMatrixFile("test/Graph2_fail.txt"); err == nil {
        // did not fail
        t.Error("Expected error, got nil.")
    } else if msg := err.Error(); msg != expectError {
        // wrong error message
        t.Errorf("Expected error \"%s\", got \"%s\".", expectError, msg)
    }
}

// test parsing Graph3.txt
func TestParseEdgesGraph3(t *testing.T) {

    graph, err := ParseEdgesFile("test/Graph3.txt", false)
    if err != nil {
        panic(err)
    }

    graphValidator(t, graph, true, 1000, 3000)
}

// test parsing Graph4.txt
func TestParseEdgesGraph4(t *testing.T) {

    graph, err := ParseEdgesFile("test/Graph4.txt", false)
    if err != nil {
        panic(err)
    }

    graphValidator(t, graph, true, 1000, 2501)
}

// validates a graph's direction and number of vertices ans edges
func graphValidator(t *testing.T, graph graphLib.GraphInterface, directed bool, numVertices, numEdges uint) {

    // test direction
    if directed == true && graph.IsDirected() == false {
        t.Error("Graph should be directed, but it is not.")
    } else if directed == false && graph.IsDirected() == true {
        t.Error("Graph should not be directed, but it is.")
    }

    // test number of vertices
    if n := graph.GetVertices().Count(); n != numVertices {
        t.Errorf("Graph should have %d vertices, got %d.", numVertices, n)
    }

    // test number of edged
    if n := graph.GetEdges().Count(); n != numEdges {
        t.Errorf("Graph should have %d edges, got %d.", numEdges, n)
    }
}