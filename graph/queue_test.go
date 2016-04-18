package graph

import (
    "testing"
)

// tests the vertex methods
func TestNearestVertexQueue(t *testing.T) {

    // some vertices, the id will be used as distance
    vertices := []vertex{
        vertex{id: id(10)},
        vertex{id: id(1)},
        vertex{id: id(4)},
        vertex{id: id(3)},
        vertex{id: id(3)},
        vertex{id: id(7)},
        vertex{id: id(77)},
        vertex{id: id(432)},
        vertex{id: id(0)},
    }
    l := len(vertices)

    // new queue that holds the nearest vertices
    q := NewNearestVertexQueue(uint(l))
    for _, vertex := range vertices {
        q.PushVertex(vertex, float64(vertex.id), nil)
    }

    // recover vertices in the right order
    for _, d := range []float64{0, 1, 3, 3, 4, 7, 10, 77, 432} {
        if q.IsEmpty() {
            t.Error("Queue is empty, but should not.")
        }

        _, distance, _ := q.PopNearestVertex()
        if distance != d {
            t.Errorf("Distance should be %.f, but is %.f.", d, distance)
        }

        l--
    }

    // check if all were recovered
    if l > 0 {
        t.Error("Queue did not return every vertex.")
    }

}
