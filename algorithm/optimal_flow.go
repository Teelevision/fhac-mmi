package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
    "github.com/teelevision/fhac-mmi/parser"
    "fmt"
)

// simple wrapper
func (this Graph) OptimalFlowCycleCancelling() ([]float64) {
    return OptimalFlowCycleCancelling(this)
}

// returns the length of the hamilton circle calculated by the nearest neighbour algorithm
func OptimalFlowCycleCancelling(graph Graph) ([]float64) {

    usage := make([]float64, graph.GetEdges().Count())

    // add super source and super destination
    superGraph, source, destination := graph.createSuperSourceAndDestinationGraph()

    maxFlow, edges := superGraph.MaxFlowEdmondsKarp(source, destination)
    fmt.Println(maxFlow)
    for _, e := range edges {
        fmt.Printf("%d -> %d = %f\n", e.GetStartVertex().GetPos(), e.GetEndVertex().GetPos(), e.flow)
    }

    return usage
}

//
// -----------------------------
//


// simple wrapper
func (this Graph) OptimalFlowSuccessiveShortestPath() ([]float64) {
    return OptimalFlowSuccessiveShortestPath(this)
}

// returns the length of the hamilton circle calculated by the nearest neighbour algorithm
func OptimalFlowSuccessiveShortestPath(graph Graph) ([]float64) {

    usage := make([]float64, graph.GetEdges().Count())

    return usage
}

//
// -----------------------------
// Helpers

// returns a new graph that contains a super source and destination and is otherwise just a copy of the base graph
// vertices must be of type parser.FlowVertex
func (this Graph) createSuperSourceAndDestinationGraph() (Graph, graphLib.VertexInterface, graphLib.VertexInterface) {

    // create a new graph
    graph := this.Clone()

    // add super source and destination vertices
    source, destination := graph.NewVertex(), graph.NewVertex()

    // connect each vertex with a balance != 0 with either source or destination
    for _, v := range graph.GetVertices().All() {
        if v != source && v != destination {
            if b := v.(*parser.FlowVertex).Balance; b > 0 {
                graph.NewWeightedEdge(source, v, b)
            } else if b < 0 {
                graph.NewWeightedEdge(v, destination, -1 * b)
            }
        }
    }

    return Graph{graph}, source, destination
}