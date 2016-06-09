package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
    "math"
)

// simple wrapper
func (this Graph) MaxFlowEdmondsKarp(start, end graphLib.VertexInterface) (float64, []FlowEdge) {
    return MaxFlowEdmondsKarp(this, start, end)
}

// returns the maximum flow using the Edmonds-Karp algorithm
func MaxFlowEdmondsKarp(graph Graph, start, end graphLib.VertexInterface) (float64, []FlowEdge) {

    // create flow graph
    G := &FlowGraph{Graph{graph.Transform(nil, func(e graphLib.EdgeInterface) graphLib.EdgeInterface {
        return &ekEdge{
            EdgeInterface: e,
            flow: 0.0,
        }
    })}}

    // get paths from start to end until none can be found anymore
    s, e, maxFlow := start.GetId(), end.GetId(), 0.0
    for path, flow := G.nextPath(s, e); path != nil; path, flow = G.nextPath(s, e) {

        // add flow to all normal edges and subtract it from all reverted edges of the path
        for _, edge := range path {
            if edge.revert {
                edge.SetFlow(edge.GetFlow() - flow)
            } else {
                edge.SetFlow(edge.GetFlow() + flow)
            }
        }

        // sum up flows
        maxFlow += flow

    }

    return maxFlow, G.GetAllFlowEdges()
}

type FlowEdge interface {
    graphLib.EdgeInterface
    GetFlow() float64
    SetFlow(float64)
    GetCapacity() float64
    GetCost() float64
}

type FlowGraph struct {
    Graph
}

func (this FlowGraph) GetAllFlowEdges() []FlowEdge {
    edges := make([]FlowEdge, this.GetEdges().Count())
    for i, e := range this.GetEdges().All() {
        edges[i] = e.(FlowEdge)
    }
    return edges
}

type ekEdge struct {
    graphLib.EdgeInterface
    flow float64
}

func (this ekEdge) GetFlow() float64 {
    return this.flow
}

func (this *ekEdge) SetFlow(flow float64) {
    this.flow = flow
}

func (this ekEdge) GetCost() float64 {
    return 0
}

func (this ekEdge) GetCapacity() float64 {
    return this.GetWeight();
}

type ekEdgeFlow struct {
    FlowEdge
    revert bool
}

type residualGraph struct {
    *graphLib.Graph
}

// creates a residual graph
func (this FlowGraph) getResidualGraph() (residualGraph) {
    resiG := graphLib.DirectedGraph()

    // copy vertices
    vertices := this.GetVertices()
    for _, v := range vertices.All() {
        resiG.NewVertexWithId(v.GetId())
    }

    // add edges
    rVertices := resiG.GetVertices()
    for _, e := range this.GetAllFlowEdges() {
        a, b := rVertices.Get(e.GetStartVertex().GetId()), rVertices.Get(e.GetEndVertex().GetId())

        // normal direction
        if restCapacity := e.GetCapacity() - e.GetFlow(); restCapacity > 0 {
            resiG.NewWeightedEdge(a, b, e.GetCost())
        }

        // opposite direction
        if e.GetFlow() > 0 {
            resiG.NewWeightedEdge(b, a, -1 * e.GetCost())
        }

    }
    return residualGraph{resiG}
}

// returns the next path
func (this *FlowGraph) nextPath(start, end uint) ([]*ekEdgeFlow, float64) {
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
    edges := this.GetAllFlowEdges()
    for i := len(path) - 1; i > 0; i-- {

        // path is reverted
        from, to := vertices.Get(path[i]), vertices.Get(path[i - 1])

        // get edge
        edge, revert := this.getEdgeFromTo(from, to), false
        if edge == nil {
            edge, revert = this.getEdgeFromTo(to, from), true
        }
        edgeFlow := &ekEdgeFlow{
            FlowEdge: edges[edge.GetPos()],
            revert: revert,
        }

        // look for the bottleneck
        if restCapacity := edgeFlow.GetCapacity() - edgeFlow.GetFlow(); restCapacity < maxFlow {
            maxFlow = restCapacity
        }

        // add vertex to reverse path
        orderedEdges = append(orderedEdges, edgeFlow)
    }

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