package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
    "github.com/teelevision/fhac-mmi/parser"
    "fmt"
    "math"
    "errors"
)

// simple wrapper
func (this Graph) OptimalFlowCycleCancelling() (float64, []float64, error) {
    return OptimalFlowCycleCancelling(this)
}

// returns the length of the hamilton circle calculated by the nearest neighbour algorithm
func OptimalFlowCycleCancelling(graph Graph) (float64, []float64, error) {

    /**
     * 1. Calculate the maximum flow through the graph.
     */

    // add super source and super destination
    superGraph, superSource, superDestination, sumSource, sumDestination := graph.createSuperSourceAndDestinationGraph()
    if sumSource != sumDestination {
        return 0.0, nil, errors.New(fmt.Sprintf("Source and destination sizes do not match (%f vs %f).", sumSource, sumDestination))
    }

    // get the maximum flow through the graph
    maxFlow, maxFlowEdges := superGraph.MaxFlowEdmondsKarp(superSource, superDestination)
    if maxFlow != sumSource {
        return 0.0, nil, errors.New(fmt.Sprintf("No flow was found (needed %f, got %f).", sumSource, maxFlow))
    }

    /**
     * 2. Prepare
     */

    // create graph we can work on
    G := &FlowGraph{Graph{graph.Transform(nil, func(e graphLib.EdgeInterface) graphLib.EdgeInterface {
        edge := e.Clone().(FlowEdge)
        edge.SetFlow(maxFlowEdges[e.GetPos()].GetFlow())
        return edge
    })}}
    vertices := G.GetVertices()
    edges := G.GetEdges()

    /**
     * 3. Finding negative cycles and improve flow through graph.
     */

    // find negative cycles
    for go_on := true; go_on; {
        go_on = false

        // try every vertices
        for _, someVertex := range vertices.All() {

            // build residual graph
            resiG := Graph{G.getResidualGraph().Graph}

            // find cycle
            someVertex = resiG.GetVertices().Get(someVertex.GetId())
            _, _, cycle := resiG.ShortestPathsMBF(someVertex, someVertex)

            // if cycle was found
            if cycle != nil {
                l := len(cycle)

                // keep track of the edges that we need to update
                edgesToUpdate := make([]struct {
                    FlowEdge
                    factor float64
                }, l)

                // find maximum
                maxFlow := math.MaxFloat64
                for u, v := -1, 0; v < l; u, v = v, v + 1 {

                    // get edge
                    f, t := vertices.Get(cycle[(u + l) % l].GetId()), vertices.Get(cycle[v].GetId())
                    e, revert, factor := G.getEdgeFromTo(f, t), false, 1.0
                    if e == nil {
                        e, revert, factor = G.getEdgeFromTo(t, f), true, -1.0
                    }
                    edge := edges.GetPos(e.GetPos()).(FlowEdge)

                    // get the max flow over this edge
                    w := edge.GetFlow()
                    if !revert {
                        w = edge.GetCapacity() - edge.GetFlow();
                    }

                    // update the cycle's max flow if lower
                    if w < maxFlow {
                        maxFlow = w
                    }

                    // update this edge after the cycle's max flow is found
                    edgesToUpdate[v].FlowEdge = edge
                    edgesToUpdate[v].factor = factor
                }

                // apply flow to edges
                for _, e := range edgesToUpdate {
                    e.SetFlow(e.GetFlow() + e.factor * maxFlow)
                }

                // we probably aren't finished yet
                go_on = true
            }
        }
    }

    /**
     * 4. Build result
     */
    usage, cost := make([]float64, graph.GetEdges().Count()), 0.0
    for i, ee := range G.GetEdges().All() {
        e := ee.(FlowEdge)
        usage[i] = e.GetFlow()
        cost += e.GetCost() * e.GetFlow()
    }
    return cost, usage, nil
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
func (this Graph) createSuperSourceAndDestinationGraph() (Graph, graphLib.VertexInterface, graphLib.VertexInterface, float64, float64) {

    // create a new graph
    graph := this.Clone()

    // add super source and destination vertices
    superSource, superDestionation := graph.NewVertex(), graph.NewVertex()

    // sum up source and destination sizes
    sumSource, sumDestination := 0.0, 0.0

    // connect each vertex with a balance != 0 with either source or destination
    for _, v := range graph.GetVertices().All() {
        if v != superSource && v != superDestionation {
            if b := v.(*parser.FlowVertex).Balance; b > 0 {
                graph.NewWeightedEdge(superSource, v, b)
                sumSource += b
            } else if b < 0 {
                graph.NewWeightedEdge(v, superDestionation, -1 * b)
                sumDestination -= b
            }
        }
    }

    return Graph{graph}, superSource, superDestionation, sumSource, sumDestination
}