package algorithm

import (
    "testing"
    "github.com/teelevision/fhac-mmi/parser"
)

func TravelingSalesmanBruteForceBenchmark(b *testing.B, file string, result float64, branchAndBound bool) {

    g, err := parser.ParseEdgesFile(file, true)
    if err != nil {
        panic(err)
    }
    graph := Graph{g}

    for n := 0; n < b.N; n++ {
        length := graph.TravelingSalesmanBruteForce(branchAndBound)
        if float64(int(length * 100 + 0.5)) / 100 != result {
            panic("TravelingSalesmanBruteForce() result is wrong")
        }
    }

}

func BenchmarkTravelingSalesmanBruteForce10(b *testing.B) {
    TravelingSalesmanBruteForceBenchmark(b, "test/K_10.txt", 38.41, false)
}

func BenchmarkTravelingSalesmanBruteForce10e(b *testing.B) {
    TravelingSalesmanBruteForceBenchmark(b, "test/K_10e.txt", 27.26, false)
}

func BenchmarkTravelingSalesmanBruteForce12(b *testing.B) {
    TravelingSalesmanBruteForceBenchmark(b, "test/K_12.txt", 45.19, false)
}

func BenchmarkTravelingSalesmanBruteForce12e(b *testing.B) {
    TravelingSalesmanBruteForceBenchmark(b, "test/K_12e.txt", 36.13, false)
}

func BenchmarkTravelingSalesmanBranchAndBound10(b *testing.B) {
    TravelingSalesmanBruteForceBenchmark(b, "test/K_10.txt", 38.41, true)
}

func BenchmarkTravelingSalesmanBranchAndBound10e(b *testing.B) {
    TravelingSalesmanBruteForceBenchmark(b, "test/K_10e.txt", 27.26, true)
}

func BenchmarkTravelingSalesmanBranchAndBound12(b *testing.B) {
    TravelingSalesmanBruteForceBenchmark(b, "test/K_12.txt", 45.19, true)
}

func BenchmarkTravelingSalesmanBranchAndBound12e(b *testing.B) {
    TravelingSalesmanBruteForceBenchmark(b, "test/K_12e.txt", 36.13, true)
}
