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

    initTime := time.Now()

    numComponents1 := algorithm.GetNumConnectedComponents(graph, algorithm.BreadthFirstSearch)
    numComponents2 := algorithm.GetNumConnectedComponents(graph, algorithm.DepthFirstSearch)

    endTime := time.Now()

    fmt.Println("Zusammenhangskomponentent:", numComponents1, numComponents2)
    fmt.Printf("Duration: %v\n", endTime.Sub(startTime))
    fmt.Printf(" - Initialization: %v\n", initTime.Sub(startTime))
    fmt.Printf(" - Calulation: %v\n", endTime.Sub(initTime))
}
