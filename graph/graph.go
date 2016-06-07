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
    return CreateNewGraphWithNumVerticesAndNumEdges(directed, 10, 10)
}

// init a new graph with given sizes
func CreateNewGraphWithNumVerticesAndNumEdges(directed bool, numVertices, numEdges uint) *Graph {
    return &Graph{
        vertices: make(vertices, 0, numVertices),
        edges: make(edges, 0, numEdges),
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
    return this.NewVertexWithId(uint(this.edgesIdProvider.NewId()))
}

// creates, adds and returns a new vertex with given id
func (this *Graph) NewVertexWithId(i uint) *vertex {

    // create with empty map of ingoing and outgoing edges
    vertex := &vertex{
        id: id(i),
        ingoingEdges: newEdges(10),
        outgoingEdges: newEdges(10),
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

func (this Graph) Transform(tV func(VertexInterface) VertexInterface, tE func(EdgeInterface) EdgeInterface) *Graph {
    g := &Graph{
        vertices: make(vertices, 0, this.GetVertices().Count()),
        edges: make(edges, 0, this.GetEdges().Count()),
        verticesIdProvider: idProvider(0),
        edgesIdProvider: idProvider(0),
        directed: this.IsDirected(),
    }

    // transform vertices
    if tV == nil {
        // default: clone
        tV = func(v VertexInterface) VertexInterface {
            return &vertex{
                id: id(v.GetId()),
                ingoingEdges: v.GetIngoingEdges(),
                outgoingEdges: v.GetOutgoingEdges(),
            }
        }
    }
    for _, v := range this.vertices.All() {
        g.vertices.add(tV(v))
    }

    // transform edges
    if tE == nil {
        // default: clone
        tE = func(e EdgeInterface) EdgeInterface {
            return &edge{
                id: id(e.GetId()),
                start: e.GetStartVertex(),
                end: e.GetEndVertex(),
                weight: e.GetWeight(),
            }
        }
    }
    for _, v := range this.edges.All() {
        g.edges.add(tE(v))
    }

    return g
}

// clone the graph without edges
func CloneGraphWithoutEdges(graph GraphInterface, numEdges uint) *Graph {
    return &Graph{
        vertices: vertices(graph.GetVertices().All()),
        edges: make(edges, 0, numEdges),
        verticesIdProvider: idProvider(graph.GetVertices().Count()),
        edgesIdProvider: idProvider(0),
        directed: graph.IsDirected(),
    }
}