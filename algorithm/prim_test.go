package algorithm

import (
    "testing"
    "github.com/teelevision/fhac-mmi/parser"
)

func BenchmarkPrimLength(b *testing.B) {

    g, err := parser.ParseEdgesFile("test/G_100_200.txt", true)
    if err != nil {
        panic(err)
    }
    start := g.GetVertices().Get(0)
    graph := Graph{g}

    for n := 0; n < b.N; n++ {
        length, _ := graph.PrimLength(start)
        if length != 27450.617104929115 {
            panic("PrimLength() result is wrong")
        }
    }

}
