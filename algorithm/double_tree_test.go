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
        length := graph.DoubleTreeHamiltonCircleLength(start)
        if length != 527.1999999999998 {
            panic("DoubleTreeHamiltonCircleLength() result is wrong")
        }
    }

}
