package graph

// vertex interface
// the general vertex has ingoing and outgoing edges
type VertexInterface interface {
    idInterface
    GetEdges() EdgesInterface
    GetIngoingEdges() editableEdgesInterface
    GetOutgoingEdges() editableEdgesInterface
    GetNeighbours() VerticesInterface
    GetIngoingNeighbours() VerticesInterface
    GetOutgoingNeighbours() VerticesInterface
}

// a basic vertex
type vertex struct {
    id
    ingoingEdges  editableEdgesInterface
    outgoingEdges editableEdgesInterface
}

// returns a combination of ingoing an outgoing edges
func (this vertex) GetEdges() EdgesInterface {
    return &mergedEdges{this.ingoingEdges, this.outgoingEdges}
}

// returns the ingoing edges
func (this vertex) GetIngoingEdges() editableEdgesInterface {
    return this.ingoingEdges
}

// returns the outgoing edges
func (this vertex) GetOutgoingEdges() editableEdgesInterface {
    return this.outgoingEdges
}

// returns a combination of ingoing an outgoing edges
func (this vertex) GetNeighbours() VerticesInterface {
    v := this.GetOutgoingNeighbours().(vertices)
    for _, n := range this.GetIngoingNeighbours().All() {
        v.add(n)
    }
    return v
}

// returns the ingoing edges
func (this vertex) GetIngoingNeighbours() VerticesInterface {
    v := vertices{}
    for _, edge := range this.ingoingEdges.All() {
        v.add(edge.GetStartVertex())
    }
    return v
}

// returns the outgoing edges
func (this vertex) GetOutgoingNeighbours() VerticesInterface {
    v := vertices{}
    for _, edge := range this.outgoingEdges.All() {
        v.add(edge.GetEndVertex())
    }
    return v
}

// interface for a map of vertices
type VerticesInterface interface {
    Get(uint) VertexInterface
    Count() uint
    All() []VertexInterface
}

// interface for an editable map of vertices
type editableVerticesInterface interface {
    VerticesInterface
    add(VertexInterface)
}

// default map of vertices
type vertices []VertexInterface

// returns single vertex or nil if not found
func (this vertices) Get(id uint) VertexInterface {
    for _, v := range this {
        if v.GetId() == id {
            return v
        }
    }
    return nil
}

// adds a vertex
func (this *vertices) add(vertex VertexInterface) {
    *this = append(*this, vertex)
}

// returns the count of vertices
func (this vertices) Count() uint {
    return uint(len(this))
}

// returns slice of all vertices
func (this vertices) All() []VertexInterface {
    return this
}
