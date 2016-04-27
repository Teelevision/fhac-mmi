package algorithm

import (
    "testing"
    "github.com/teelevision/fhac-mmi/parser"
)

func BenchmarkNearestNeighbourHamiltonCircleLength(b *testing.B) {

    g, err := parser.ParseEdgesFile("test/K_100.txt", true)
    if err != nil {
        panic(err)
    }
    graph := Graph{g}
    start := graph.GetVertices().Get(0)

    for n := 0; n < b.N; n++ {
        length := graph.NearestNeighbourHamiltonCircleLength(start)
        if length != 323.93 {
            panic("NearestNeighbourHamiltonCircleLength() result is wrong")
        }
    }

}
