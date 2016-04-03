package graph

// edge interface
// an edge has both start and end vertex and a weight
type EdgeInterface interface {
    idInterface
    GetStartVertex() VertexInterface
    GetEndVertex() VertexInterface
    GetWeight() float64
    SetWeight(weight float64)
}

// a basic edge
type edge struct {
    id
    start  VertexInterface
    end    VertexInterface
    weight float64
}

// returns the start vertex
func (this edge) GetStartVertex() VertexInterface {
    return this.start;
}

// returns the end vertex
func (this edge) GetEndVertex() VertexInterface {
    return this.end;
}

// returns the weight
func (this edge) GetWeight() float64 {
    return this.weight;
}

// sets the weight
func (this *edge) SetWeight(weight float64) {
    this.weight = weight;
}

// interface for a map of edges
type EdgesInterface interface {
    Get(uint) EdgeInterface
    Count() uint
    All() []EdgeInterface
}

// interface for an editable map of edges
type editableEdgesInterface interface {
    EdgesInterface
    add(EdgeInterface)
}

// a simple map of edges
type edges []EdgeInterface

// returns the edge with the given id or nil if not found
func (this edges) Get(id uint) EdgeInterface {
    for _, e := range this {
        if e.GetId() == id {
            return e
        }
    }
    return nil
}

// adds an edge
func (this *edges) add(edge EdgeInterface) {
    *this = append(*this, edge)
}

// returns the number of edges
func (this edges) Count() uint {
    return uint(len(this))
}

// returns slice of all edges
func (this edges) All() []EdgeInterface {
    return this
}


// a slice of maps of edges
// can be used to combine edge maps for reading
type mergedEdges []EdgesInterface

// returns the edge for the id or nil if not found
func (this *mergedEdges) merge(edges EdgesInterface) {
    *this = append(*this, edges)
}

// returns the edge for the id or nil if not found
func (this mergedEdges) Get(id uint) EdgeInterface {
    for _, edges := range this {
        if edge := edges.Get(id); edge != nil {
            return edge
        }
    }
    return nil
}

// returns the number of edges
func (this mergedEdges) Count() (num uint) {
    for _, edges := range this {
        num += edges.Count()
    }
    return
}

// returns a map of all edges
func (this mergedEdges) All() []EdgeInterface {
    allEdges := make([]EdgeInterface, 0, this.Count())
    for _, edges := range this {
        for _, edge := range edges.All() {
            allEdges = append(allEdges, edge)
        }
    }
    return allEdges
}
