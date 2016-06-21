package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
    "github.com/teelevision/fhac-mmi/parser"
)

// simple wrapper
func (this Graph) MaxMatching() ([]graphLib.EdgeInterface) {
    return MaxMatching(this)
}

// returns the maximum flow using the Edmonds-Karp algorithm
func MaxMatching(graph Graph) ([]graphLib.EdgeInterface) {

    /*
     * 1. Create graph where all edges go from group 0 to 1.
     */
    G := Graph{graph.Transform(nil, func(e graphLib.EdgeInterface) graphLib.EdgeInterface {
        edge := e.Clone()

        // set weight to 1
        edge.SetWeight(1.0)

        // swap direction if edge starts in group 1
        if graph.GetVertices().GetPos(e.GetStartVertex().GetPos()).(*parser.GroupVertex).GetGroup() == 1 {
            edge.SwapStartAndEnd()
        }

        return edge
    })};
    G.SetDirected(true)

    /*
     * 2. Add super source and super target.
     */
    superSource, superTarget := G.NewVertex(), G.NewVertex()

    /*
     * 3. Connect super source to all vertices in group 0 and all vertices in group 1 to super target.
     */
    for _, v := range G.GetVertices().All() {
        if v != superSource && v != superTarget {
            if v.(*parser.GroupVertex).GetGroup() == 0 {
                G.NewWeightedEdge(superSource, v, 1.0)
            } else {
                G.NewWeightedEdge(v, superTarget, 1.0)
            }
        }
    }

    /*
     * 4. Set capacity of all edges to 1.
     */
    // already done in 1. and 3. in form of weights

    /*
     * 5. Find maximum flow between the super source and target.
     */
    maxFlow, flowEdges := G.MaxFlowEdmondsKarp(superSource, superTarget)

    /*
     * 6. Build response.
     */
    matchedEdges := make([]graphLib.EdgeInterface, 0, int(maxFlow))
    for _, e := range flowEdges {
        if s, t := e.GetStartVertex().GetPos(), e.GetEndVertex().GetPos(); s != superSource.GetPos() && t != superTarget.GetPos() {
            if e.GetFlow() >= 1.0 {
                matchedEdges = append(matchedEdges, e)
            }
        }
    }

    // assert
    if int(maxFlow) != len(matchedEdges) {
        panic("The maximum flow should match the number of edges with a flow.")
    }

    return matchedEdges
}
