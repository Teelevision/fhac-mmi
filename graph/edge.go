package graph

// edge interface
// an edge has both start and end vertex and a weight
type EdgeInterface interface {
    idInterface
    GetPos() int
    setPos(int)
    GetStartVertex() VertexInterface
    GetEndVertex() VertexInterface
    SwapStartAndEnd()
    GetOtherVertex(VertexInterface) VertexInterface
    GetWeight() float64
    SetWeight(weight float64)
    Clone() EdgeInterface
}

// a basic edge
type edge struct {
    id
    pos    int
    start  VertexInterface
    end    VertexInterface
    weight float64
}

// returns the position
func (this edge) GetPos() int {
    return this.pos
}

// sets the position
func (this *edge) setPos(pos int) {
    this.pos = pos
}

// returns the start vertex
func (this edge) GetStartVertex() VertexInterface {
    return this.start
}

// returns the end vertex
func (this edge) GetEndVertex() VertexInterface {
    return this.end
}

// swaps start and end
func (this *edge) SwapStartAndEnd() {
    this.start, this.end = this.end, this.start
}

// returns the other vertex that is not the given one
func (this edge) GetOtherVertex(vertex VertexInterface) VertexInterface {
    if this.start == vertex {
        return this.end
    }
    return this.start
}

// returns the weight
func (this edge) GetWeight() float64 {
    return this.weight
}

// sets the weight
func (this *edge) SetWeight(weight float64) {
    this.weight = weight
}

// clones the edge
func (this edge) Clone() EdgeInterface {
    return &edge{
        id: id(this.GetId()),
        start: this.GetStartVertex(),
        end: this.GetEndVertex(),
        weight: this.GetWeight(),
    }
}

// interface for a map of edges
type EdgesInterface interface {
    Get(uint) EdgeInterface
    GetPos(int) EdgeInterface
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

// creates new instance of edges with the given capacity
func newEdges(capacity int) *edges {
    e := edges(make([]EdgeInterface, 0, capacity))
    return &e
}

// returns the edge with the given id or nil if not found
func (this edges) Get(id uint) EdgeInterface {
    for _, e := range this {
        if e.GetId() == id {
            return e
        }
    }
    return nil
}

// returns the edge at the given position
func (this edges) GetPos(pos int) EdgeInterface {
    return this[pos]
}

// adds an edge
func (this *edges) add(edge EdgeInterface) {
    // if slice is full, double its size
    if len(*this) == cap(*this) {
        newSlice := make([]EdgeInterface, len(*this), 2 * cap(*this) + 1)
        copy(newSlice, *this)
        *this = newSlice
    }
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
        allEdges = append(allEdges, edges.All()...)
    }
    return allEdges
}

// returns the edge at the given position
func (this mergedEdges) GetPos(pos int) EdgeInterface {
    for _, edges := range this {
        if c := int(edges.Count()); pos < c {
            return edges.GetPos(pos)
        } else {
            pos -= c
        }
    }
    return nil
}
