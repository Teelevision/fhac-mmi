package algorithm

import (
    "testing"
    "github.com/teelevision/fhac-mmi/parser"
)

func BenchmarkKruskalLength(b *testing.B) {

    g, err := parser.ParseEdgesFile("test/G_100_200.txt", true)
    if err != nil {
        panic(err)
    }
    graph := Graph{g}

    for n := 0; n < b.N; n++ {
        length := graph.KruskalLength()
        if length != 27450.617104929264 {
            panic("KruskalLength() result is wrong")
        }
    }

}
