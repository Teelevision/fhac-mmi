package algorithm

import (
    "testing"
    "github.com/teelevision/fhac-mmi/parser"
)

func BenchmarkDoubleTreeHamiltonCircleLength(b *testing.B) {

    g, err := parser.ParseEdgesFile("test/K_100.txt", true)
    if err != nil {
        panic(err)
    }
    graph := Graph{g}
    start := graph.GetVertices().Get(0)

    for n := 0; n < b.N; n++ {
        _, length := graph.DoubleTreeHamiltonCircle(Prim, start)
        if length != 385.44999999999993 {
            panic("DoubleTreeHamiltonCircleLength() result is wrong")
        }
    }

}
