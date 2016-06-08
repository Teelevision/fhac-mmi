package algorithm

import (
    _ "github.com/teelevision/fhac-mmi/graph"
)

// simple wrapper
func (this Graph) OptimalFlowCycleCancelling() ([]float64) {
    return OptimalFlowCycleCancelling(this)
}

// returns the length of the hamilton circle calculated by the nearest neighbour algorithm
func OptimalFlowCycleCancelling(graph Graph) ([]float64) {

    usage := make([]float64, graph.GetEdges().Count())

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