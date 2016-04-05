package parser

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
    "io"
    "os"
)

// parses a file that contains an edges list
func ParseEdgesFile(file string, withWeights bool) (*graphLib.Graph, error) {
    f, _ := os.Open(file)
    graph, err := ParseEdges(f, withWeights)
    return graph, err
}

// parses an edges list with or without weights
func ParseEdges(reader io.Reader, withWeights bool) (*graphLib.Graph, error) {

    // parse vertices
    graph, vertices, scanner, err := parseHeader(reader)
    if err != nil {
        return graph, err
    }

    // create edges
    for {

        // parse start vertex and test if input is empty
        start, err := parseInt(scanner)
        if err != nil && err.Error() == "EOF" {
            break
        } else if err != nil {
            return graph, err
        }

        // parse end vertex
        end, err := parseInt(scanner)
        if err != nil {
            return graph, err
        }

        // parse weight
        weight := 1.0
        if withWeights {
            weight, err = parseFloat(scanner)
            if err != nil {
                return graph, err
            }
        }

        // create edge
        graph.NewWeightedEdge(vertices[start], vertices[end], weight)
    }

    return graph, nil
}