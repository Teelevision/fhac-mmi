package parser

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
    "io"
    "os"
)

// parses an file containing a flow graph
func ParseFlowFile(file string) (*graphLib.Graph, error) {
    f, _ := os.Open(file)
    graph, err := ParseFlow(f)
    return graph, err
}

type FlowVertex struct {
    graphLib.VertexInterface
    Balance float64
}

type FlowEdge struct {
    graphLib.EdgeInterface
    Capacity float64
}

// parses an file containing a flow graph
func ParseFlow(reader io.Reader) (*graphLib.Graph, error) {

    // parse vertices
    graph, vertices, scanner, err := parseHeader(reader)
    if err != nil {
        return graph, err
    }

    // add balance to vertices
    graph = graph.Transform(func(vertex graphLib.VertexInterface) graphLib.VertexInterface {
        balance, err := parseFloat(scanner)
        if err != nil {
            panic(err)
        }
        return &FlowVertex{
            VertexInterface: vertex,
            Balance: balance,
        }
    }, nil)


    // create edges
    capacities := make([]float64, 0)
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
        weight, err := parseFloat(scanner)
        if err != nil {
            return graph, err
        }

        // parse capacity
        capacity, err := parseFloat(scanner)
        if err != nil {
            return graph, err
        }
        capacities = append(capacities, capacity)

        // create edge
        graph.NewWeightedEdge(vertices[start], vertices[end], weight)
    }

    // add capacity to edges
    graph = graph.Transform(nil, func(edge graphLib.EdgeInterface) graphLib.EdgeInterface {
        return &FlowEdge{
            EdgeInterface: edge,
            Capacity: capacities[edge.GetPos()],
        }
    })

    return graph, nil
}