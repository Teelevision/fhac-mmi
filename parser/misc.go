package parser

import (
    "bufio"
    "strconv"
    "errors"
    "io"
    graphLib "github.com/teelevision/fhac-mmi/graph"
)

// parses an integer
func parseInt(s *bufio.Scanner) (int, error) {
    if !s.Scan() {
        return 0, errors.New("EOF")
    }
    if num, err := strconv.Atoi(s.Text()); err != nil {
        return 0, err
    } else {
        return num, nil
    }
}

// parses a float
func parseFloat(s *bufio.Scanner) (float64, error) {
    if !s.Scan() {
        return 0, errors.New("EOF")
    }
    if num, err := strconv.ParseFloat(s.Text(), 64); err != nil {
        return 0, err
    } else {
        return num, nil
    }
}


// parses the header of a adjacency matrix or edge list
// The header contains the number of vertices.
func parseHeader(reader io.Reader) (*graphLib.Graph, []graphLib.VertexInterface, *bufio.Scanner, error) {

    graph := graphLib.DirectedGraph()

    scanner := bufio.NewScanner(reader)
    scanner.Split(bufio.ScanWords)

    // get number of vertices
    numVertices, err := parseInt(scanner)
    if err != nil {
        return graph, nil, scanner, err
    }

    // create vertices
    vertices := make([]graphLib.VertexInterface, numVertices)
    for v := 0; v < numVertices; v++ {
        vertices[v] = graph.NewVertex()
    }

    return graph, vertices, scanner, nil
}