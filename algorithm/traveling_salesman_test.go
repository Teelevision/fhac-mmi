package algorithm

import (
    "testing"
    "github.com/teelevision/fhac-mmi/parser"
)

func TravelingSalesmanBruteForceBenchmark(b *testing.B, file string, result float64) {

    g, err := parser.ParseEdgesFile(file, true)
    if err != nil {
        panic(err)
    }
    graph := Graph{g}

    for n := 0; n < b.N; n++ {
        length := graph.TravelingSalesmanBruteForce()
        if float64(int(length * 100 + 0.5)) / 100 != result {
            panic("TravelingSalesmanBruteForce() result is wrong")
        }
    }

}

func BenchmarkTravelingSalesmanBruteForce10(b *testing.B) {
    TravelingSalesmanBruteForceBenchmark(b, "test/K_10.txt", 38.41)
}

func BenchmarkTravelingSalesmanBruteForce10e(b *testing.B) {
    TravelingSalesmanBruteForceBenchmark(b, "test/K_10e.txt", 27.26)
}

func BenchmarkTravelingSalesmanBruteForce12(b *testing.B) {
    TravelingSalesmanBruteForceBenchmark(b, "test/K_12.txt", 45.19)
}

func BenchmarkTravelingSalesmanBruteForce12e(b *testing.B) {
    TravelingSalesmanBruteForceBenchmark(b, "test/K_12e.txt", 36.13)
}
