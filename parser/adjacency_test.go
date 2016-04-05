package parser

import (
    "testing"
    "reflect"
    graphLib "github.com/teelevision/fhac-mmi/graph"
)

// test parsing Graph1.txt
func TestParseAdjacencyMatrix(t *testing.T) {

    graph, err := ParseAdjacencyMatrixFile("test/Graph1.txt")
    if err != nil {
        panic(err)
    }

    // test direction
    if graph.IsDirected() != true {
        t.Error("Graph should be directed, but is not.")
    }

    // test number of vertices
    if n := graph.GetVertices().Count(); n != 15 {
        t.Errorf("Graph should have 15 vertices, got %d.", n)
    }

    // test number of edged
    if n := graph.GetEdges().Count(); n != 30 {
        t.Errorf("Graph should have 30 edges, got %d.", n)
    }

    // test each vertex
    for i, v := range graph.GetVertices().All() {
        num := []uint{6, 4, 6, 4, 4, 2, 6, 6, 4, 6, 2, 2, 4, 2, 2}
        if n := v.GetNeighbours().Count(); n != num[i] {
            t.Errorf("Expected vertex #%d to have %d neighbours, got %d.", i, num[i], n)
        }
    }

    // for the first vertex: test if neighbours are right
    vertices := graph.GetVertices()
    v := vertices.Get(0)
    // neighbours are: 6, 9 and 13
    neighbours := []graphLib.VertexInterface{vertices.Get(6), vertices.Get(9), vertices.Get(13)}
    if ne := v.GetIngoingNeighbours().All(); !reflect.DeepEqual(ne, neighbours) {
        t.Errorf("Expected %v, got $v.", neighbours, ne)
    }
}

// test failing to parse Graph1_fail.txt
// In this graph the number of vertices is too high.
func TestParseAdjacencyMatrixFail(t *testing.T) {
    expectError := "EOF"
    if _, err := ParseAdjacencyMatrixFile("test/Graph1_fail.txt"); err == nil {
        // did not fail
        t.Error("Expected error, got nil.")
    } else if msg := err.Error(); msg != expectError {
        // wrong error message
        t.Errorf("Expected error \"%s\", got \"%s\".", expectError, msg)
    }
}

// test failing to parse Graph1_fail2.txt
// In this graph the number of vertices is not a number.
func TestParseAdjacencyMatrixFail2(t *testing.T) {
    expectError := "strconv.ParseInt: parsing \"notanumber\": invalid syntax"
    if _, err := ParseAdjacencyMatrixFile("test/Graph1_fail2.txt"); err == nil {
        // did not fail
        t.Error("Expected error, got nil.")
    } else if msg := err.Error(); msg != expectError {
        // wrong error message
        t.Errorf("Expected error \"%s\", got \"%s\".", expectError, msg)
    }
}