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

// adds a new vertex
func (this *Graph) AddVertex(vertex VertexInterface) VertexInterface {
    vertex.setPos(len(this.vertices))
    this.vertices.add(vertex)
    return vertex
}

// creates, adds and returns a new vertex
func (this *Graph) NewVertex() VertexInterface {
    return this.NewVertexWithId(uint(this.edgesIdProvider.NewId()))
}

// creates, adds and returns a new vertex with given id
func (this *Graph) NewVertexWithId(i uint) VertexInterface {

    // create with empty map of ingoing and outgoing edges
    vertex := &vertex{
        id: id(i),
        ingoingEdges: newEdges(10),
        outgoingEdges: newEdges(10),
    }

    return this.AddVertex(vertex)
}

// adds a new custom vertex with given id
func (this *Graph) NewCustomVertex(transform func(VertexInterface) VertexInterface) VertexInterface {

    // create with empty map of ingoing and outgoing edges
    vertex := &vertex{
        id: this.edgesIdProvider.NewId(),
        ingoingEdges: newEdges(10),
        outgoingEdges: newEdges(10),
    }

    return this.AddVertex(transform(vertex))
}

// creates, adds and returns a new edge with a weight of 1
func (this *Graph) AddEdge(edge EdgeInterface) EdgeInterface {
    // add edge as outgoing/ingoing to the source/target
    edge.GetStartVertex().GetOutgoingEdges().add(edge)
    edge.GetEndVertex().GetIngoingEdges().add(edge)

    // add to this graph
    edge.setPos(len(this.edges))
    this.edges.add(edge)
    return edge
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

    return this.AddEdge(edge)
}

// creates, adds and returns a new custom edge
func (this *Graph) NewCustomEdge(start, end VertexInterface, transform func(EdgeInterface) EdgeInterface) EdgeInterface {

    // create edge with the given source and target
    edge := &edge{
        id: this.verticesIdProvider.NewId(),
        start: start,
        end: end,
    }

    return this.AddEdge(transform(edge))
}

func (this Graph) Transform(tV func(VertexInterface) VertexInterface, tE func(EdgeInterface) EdgeInterface) *Graph {
    g := &Graph{
        vertices: make(vertices, 0, this.GetVertices().Count()),
        edges: make(edges, 0, this.GetEdges().Count()),
        verticesIdProvider: this.verticesIdProvider,
        edgesIdProvider: this.edgesIdProvider,
        directed: this.IsDirected(),
    }

    // transform vertices
    if tV == nil {
        // default: clone
        tV = func(v VertexInterface) VertexInterface {
            return v.Clone()
        }
    }
    for _, v := range this.vertices.All() {
        g.AddVertex(tV(v))
    }

    // transform edges
    if tE == nil {
        // default: clone
        tE = func(e EdgeInterface) EdgeInterface {
            return e.Clone()
        }
    }
    for _, v := range this.edges.All() {
        g.AddEdge(tE(v))
    }

    return g
}

// clones the graph
func (this Graph) Clone() *Graph {
    return this.Transform(nil, nil)
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