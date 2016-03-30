package graph

// vertex interface
// the general vertex has ingoing and outgoing edges
type VertexInterface interface {
    idInterface
    GetEdges() EdgesInterface
    GetIngoingEdges() editableEdgesInterface
    GetOutgoingEdges() editableEdgesInterface
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

// interface for a map of vertices
type VerticesInterface interface {
    Get(uint) VertexInterface
    Count() uint
    All() map[uint]VertexInterface
}

// interface for an editable map of vertices
type editableVerticesInterface interface {
    VerticesInterface
    add(VertexInterface)
}

// default map of vertices
type vertices map[uint]VertexInterface

// returns single vertex or nil if not found
func (this vertices) Get(id uint) VertexInterface {
    if vertex, ok := this[id]; ok {
        return vertex
    }
    return nil
}

// adds a vertex
func (this vertices) add(vertex VertexInterface) {
    this[vertex.GetId()] = vertex
}

// returns the count of vertices
func (this vertices) Count() uint {
    return uint(len(this))
}

// returns map of all vertices
func (this vertices) All() map[uint]VertexInterface {
    return this
}
