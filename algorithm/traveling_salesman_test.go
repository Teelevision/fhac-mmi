package algorithm

import (
    "testing"
    "github.com/teelevision/fhac-mmi/parser"
)

func BenchmarkTravelingSalesmanBruteForce(b *testing.B) {

    g, err := parser.ParseEdgesFile("test/K_10e.txt", true)
    if err != nil {
        panic(err)
    }
    graph := Graph{g}

    for n := 0; n < b.N; n++ {
        length := graph.TravelingSalesmanBruteForce()
        if length != 27.259999999999994 {
            panic("TravelingSalesmanBruteForce() result is wrong")
        }
    }

}
