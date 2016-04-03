package graph

// graph interface
// every graph has vertices and edges
type GraphInterface interface {
    GetVertices() VerticesInterface
    GetEdges() EdgesInterface
    IsDirected() bool
}

// default graph
type graph struct {
    vertices           vertices
    edges              edges
    verticesIdProvider idProvider
    edgesIdProvider    idProvider
    directed           bool
}

// init a new graph
func CreateNewGraph(directed bool) *graph {
    // contains empty maps of vertices and edges
    return &graph{
        vertices: vertices{},
        edges: edges{},
        verticesIdProvider: idProvider(0),
        edgesIdProvider: idProvider(0),
        directed: directed,
    }
}

// init a new directed graph
func DirectedGraph() *graph {
    return CreateNewGraph(true)
}

// init a new undirected graph
func UndirectedGraph() *graph {
    return CreateNewGraph(false)
}

// returns whether the graph is directed
func (this graph) IsDirected() bool {
    return this.directed
}

// returns whether the graph is directed
func (this *graph) SetDirected(directed bool) {
    this.directed = directed
}

// returns the vertices map
func (this graph) GetVertices() VerticesInterface {
    return this.vertices
}

// returns the edges map
func (this graph) GetEdges() EdgesInterface {
    return this.edges
}

// creates, adds and returns a new vertex
func (this *graph) NewVertex() *vertex {

    // create with empty map of ingoing and outgoing edges
    vertex := &vertex{
        id: this.edgesIdProvider.NewId(),
        ingoingEdges: &edges{},
        outgoingEdges: &edges{},
    }

    // add to this graph
    this.vertices.add(vertex)

    return vertex
}

// creates, adds and returns a new edge with a weight of 1
func (this *graph) NewEdge(source, target VertexInterface) EdgeInterface {
    return this.NewWeightedEdge(source, target, 1.0)
}


// creates, adds and returns a new edge
func (this *graph) NewWeightedEdge(source, target VertexInterface, weight float64) EdgeInterface {

    // create edge with the given source and target and a default weight
    edge := &edge{
        id: this.verticesIdProvider.NewId(),
        start: source,
        end: target,
        weight: weight,
    }

    // add edge as outgoing/ingoing to the source/target
    source.GetOutgoingEdges().add(edge)
    target.GetIngoingEdges().add(edge)

    // add to this graph
    this.edges.add(edge)

    return edge
}
