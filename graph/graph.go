package graph

// graph interface
// every graph has vertices and edges
type GraphInterface interface {
    GetVertices() VerticesInterface
    GetEdges() EdgesInterface
    IsDirected() bool
}

// default graph
type Graph struct {
    vertices           vertices
    edges              edges
    verticesIdProvider idProvider
    edgesIdProvider    idProvider
    directed           bool
}

// init a new graph
func CreateNewGraph(directed bool) *Graph {
    // contains empty maps of vertices and edges
    return &Graph{
        vertices: vertices{},
        edges: edges{},
        verticesIdProvider: idProvider(0),
        edgesIdProvider: idProvider(0),
        directed: directed,
    }
}

// init a new directed graph
func DirectedGraph() *Graph {
    return CreateNewGraph(true)
}

// init a new undirected graph
func UndirectedGraph() *Graph {
    return CreateNewGraph(false)
}

// returns whether the graph is directed
func (this Graph) IsDirected() bool {
    return this.directed
}

// returns whether the graph is directed
func (this *Graph) SetDirected(directed bool) {
    this.directed = directed
}

// returns the vertices map
func (this Graph) GetVertices() VerticesInterface {
    return this.vertices
}

// returns the edges map
func (this Graph) GetEdges() EdgesInterface {
    return this.edges
}

// creates, adds and returns a new vertex
func (this *Graph) NewVertex() *vertex {

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
func (this *Graph) NewEdge(start, end VertexInterface) EdgeInterface {
    return this.NewWeightedEdge(start, end, 1.0)
}


// creates, adds and returns a new edge
func (this *Graph) NewWeightedEdge(start, end VertexInterface, weight float64) EdgeInterface {

    // create edge with the given source and target and a default weight
    edge := &edge{
        id: this.verticesIdProvider.NewId(),
        start: start,
        end: end,
        weight: weight,
    }

    // add edge as outgoing/ingoing to the source/target
    start.GetOutgoingEdges().add(edge)
    end.GetIngoingEdges().add(edge)

    // add to this graph
    this.edges.add(edge)

    return edge
}
