package parser

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
    "io"
    "os"
)

// parses an file containing a bipartite graph
func ParseBipartiteFile(file string) (*graphLib.Graph, error) {
    f, _ := os.Open(file)
    graph, err := ParseBipartite(f)
    return graph, err
}

type GroupVertex struct {
    graphLib.VertexInterface
    Group int
}

func (this GroupVertex) Clone() graphLib.VertexInterface {
    return &GroupVertex{
        VertexInterface: this.VertexInterface.Clone(),
        Group: this.Group,
    }
}

func (this GroupVertex) GetGroup() int {
    return this.Group
}

// parses an file containing a bipartite graph
func ParseBipartite(reader io.Reader) (*graphLib.Graph, error) {

    // parse vertices
    graph, vertices, scanner, err := parseHeader(reader)
    if err != nil {
        return graph, err
    }

    // parse num of vertices in the first group
    numFirstGroup, err := parseInt(scanner)
    if err != nil {
        return graph, err
    }

    // use GroupVertex object which will contain the number of the group
    graph = graph.Transform(func(v graphLib.VertexInterface) graphLib.VertexInterface {
        group := 0
        if v.GetPos() >= numFirstGroup {
            group = 1
        }
        return &GroupVertex{
            VertexInterface: v,
            Group: group,
        }
    }, nil)

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

        // create edge
        graph.NewWeightedEdge(vertices[start], vertices[end], 1.0)
    }

    return graph, nil
}