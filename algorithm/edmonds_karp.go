package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
    "math"
    "fmt"
)

// simple wrapper
func (this Graph) MaxFlowEdmondsKarp(start, end graphLib.VertexInterface) (float64) {
    return MaxFlowEdmondsKarp(this, start, end)
}

// returns the maximum flow using the Edmonds-Karp algorithm
func MaxFlowEdmondsKarp(graph Graph, start, end graphLib.VertexInterface) (float64) {

    numEdges := graph.GetEdges().Count()

    // wrap graph so that we can use our own edges
    G := &ekGraph{
        Graph: graph,
        edges: make([]*ekEdge, numEdges),
        mapEdges: make(map[graphLib.EdgeInterface]*ekEdge, numEdges),
    }

    // create vertices
    for i, e := range G.GetEdges().All() {
        edge := &ekEdge{
            EdgeInterface: e,
            flow: 0.0,
        }
        G.edges[i] = edge
        G.mapEdges[e] = edge
    }

    // get paths from start to end until none can be found anymore
    s, e, maxFlow := start.GetId(), end.GetId(), 0.0
    for path, flow := G.nextPath(s, e); path != nil; path, flow = G.nextPath(s, e) {

        // add flow to all normal edges and subtract it from all reverted edges of the path
        for _, edge := range path {
            if edge.revert {
                edge.flow -= flow
            } else {
                edge.flow += flow
            }
        }

        // sum up flows
        maxFlow += flow

    }

    return maxFlow
}

type ekGraph struct {
    Graph
    edges    []*ekEdge
    mapEdges map[graphLib.EdgeInterface]*ekEdge
}

type ekEdge struct {
    graphLib.EdgeInterface
    flow float64
}

func (this ekEdge) getCapacity() float64 {
    return this.GetWeight() - this.flow;
}

type ekEdgeFlow struct {
    *ekEdge
    revert bool
}

type residualGraph struct {
    *graphLib.Graph
}

// creates a residual graph
func (this ekGraph) getResidualGraph() (residualGraph) {
    resiG := graphLib.DirectedGraph()

    // copy vertices
    vertices := this.GetVertices()
    for _, v := range vertices.All() {
        resiG.NewVertexWithId(v.GetId())
    }

    // add edges
    rVertices := resiG.GetVertices()
    for _, e := range this.edges {
        a, b := rVertices.Get(e.GetStartVertex().GetId()), rVertices.Get(e.GetEndVertex().GetId())

        // normal direction
        if capacity := e.getCapacity(); capacity > 0 {
            resiG.NewWeightedEdge(a, b, capacity)
        }

        // opposite direction
        if e.flow > 0 {
            resiG.NewWeightedEdge(b, a, e.flow)
        }

    }
    return residualGraph{resiG}
}

// returns the next path
func (this *ekGraph) nextPath(start, end uint) ([]*ekEdgeFlow, float64) {
    vertices := this.GetVertices()

    // create residual graph
    resiG := this.getResidualGraph()

    // get path
    path := resiG.pathWithTheLeastHops(start, end)
    if path == nil {
        return nil, 0.0
    }

    // reverse path
    orderedEdges := make([]*ekEdgeFlow, 0, len(path))

    // get possible flow
    maxFlow := math.MaxFloat64
    for i := len(path) - 1; i > 0; i-- {

        // path is reverted
        from, to := vertices.Get(path[i]), vertices.Get(path[i - 1])

        // get edge
        edgeFlow := &ekEdgeFlow{
            ekEdge: this.mapEdges[this.getEdgeFromTo(from, to)],
            revert: false,
        }
        if edgeFlow.ekEdge == nil {
            // get other edge
            edgeFlow.ekEdge = this.mapEdges[this.getEdgeFromTo(to, from)]
            edgeFlow.revert = true
        }

        // look for the bottleneck
        if edgeFlow.getCapacity() < maxFlow {
            maxFlow = edgeFlow.getCapacity()
        }

        // add vertex to reverse path
        orderedEdges = append(orderedEdges, edgeFlow)
    }

    fmt.Println(path, maxFlow)
    return orderedEdges, maxFlow
};

// returns a path with the least hops between start and end
func (this residualGraph) pathWithTheLeastHops(start, end uint) ([]uint) {
    vertices := this.GetVertices()
    num := vertices.Count()

    // own vertex type
    type rVertex struct {
        graphLib.VertexInterface
        prev   *rVertex
        queued bool
    }

    // map vertices to own type
    vs := make(map[graphLib.VertexInterface]*rVertex, num)
    for _, v := range vertices.All() {
        vs[v] = &rVertex{
            VertexInterface: v,
            prev: nil,
            queued: false,
        }
    }

    // queue
    q := make([]*rVertex, 0, num)
    queue := func(v *rVertex) {
        q = append(q, v)
        v.queued = true
    }

    // add start to queue
    queue(vs[vertices.Get(start)])

    // go through queue
    for i := 0; i < len(q); i++ {
        vertex := q[i]

        // go through neighbours
        for _, vv := range vertex.GetOutgoingNeighbours().All() {
            v := vs[vv]
            if !v.queued {
                v.prev = vertex

                // flow to end found
                if v.GetId() == end {
                    // return slice of IDs (reverse order)
                    result := make([]uint, 0)
                    for current := v; current != nil; current = current.prev {
                        result = append(result, current.GetId())
                    }
                    return result
                }

                // add to queue
                queue(v)

            }
        }
    }

    // none found
    return nil
}