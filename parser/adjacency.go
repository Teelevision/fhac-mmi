package parser

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
    "io"
    "os"
)

// parses a file that contains an adjacency matrix
func ParseAdjacencyMatrixFile(file string) (graphLib.GraphInterface, error) {
    f, _ := os.Open(file)
    graph, err := ParseAdjacencyMatrix(f)
    return graph, err
}

// parses an adjacency matrix
func ParseAdjacencyMatrix(reader io.Reader) (graphLib.GraphInterface, error) {

    // parse vertices
    graph, vertices, scanner, err := parseHeader(reader)
    if err != nil {
        return graph, err
    }

    // create edges
    for row := 0; row < len(vertices); row++ {
        for col := 0; col < len(vertices); col++ {
            if weight, err := parseFloat(scanner); err != nil {
                return graph, err
            } else if weight > 0 {
                graph.NewWeightedEdge(vertices[row], vertices[col], weight)
            }
        }
    }

    return graph, nil
}