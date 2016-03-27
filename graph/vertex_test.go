package graph

import (
    "testing"
    "reflect"
)

// tests the vertex methods
func TestVertex(t *testing.T) {

    idProvider := idProvider(0)
    i := uint(123)
    ingoingEdges := edges{}
    ingoingEdges.add(&edge{id: idProvider.NewId()})
    outgoingEdges := edges{}
    outgoingEdges.add(&edge{id: idProvider.NewId()})

    v := vertex{
        id: id(i),
        ingoingEdges: ingoingEdges,
        outgoingEdges: outgoingEdges,
    }

    if v.GetId() != i {
        t.Errorf("Vertex ID should be %d, got %d.", i, v.GetId())
    }

    if !reflect.DeepEqual(v.GetIngoingEdges(), ingoingEdges) {
        t.Error("Could not retrieve ingoing edges.")
    }

    if !reflect.DeepEqual(v.GetOutgoingEdges(), outgoingEdges) {
        t.Error("Could not retrieve outgoing edges.")
    }
}

// tests the vertices methods
func TestVertices(t *testing.T) {

    idProvider := idProvider(0)
    v := vertices{}

    var ids [10]id
    var vs [10]vertex
    for i := 0; i < 10; i++ {
        ids[i] = idProvider.NewId()
        vs[i] = vertex{id: ids[i]}
        v.Add(vs[i])
    }

    if v.Count() != 10 {
        t.Errorf("Vertices should contain 10 elements, got %d.", v.Count())
    }

    for i, id := range ids {
        if !reflect.DeepEqual(v.Get(uint(id)), vs[i]) {
            t.Error("Could not recover vertices.")
        }
    }
}