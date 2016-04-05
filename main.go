package main

import (
    "github.com/teelevision/fhac-mmi/parser"
    "github.com/teelevision/fhac-mmi/algorithm"
    "fmt"
    "time"
)

func main() {

    startTime := time.Now()

    g, err := parser.ParseEdgesFile("prakt/Graph4.txt", false)
    if err != nil {
        panic(err)
    }
    g.SetDirected(false)
    graph := algorithm.Graph{g}

    numComponents := graph.GetNumConnectedComponents()

    endTime := time.Now()

    fmt.Println("Zusammenhangskomponentent: ", numComponents)
    fmt.Printf("Duration: %v\n", endTime.Sub(startTime))
}
