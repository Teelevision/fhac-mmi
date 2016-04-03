package graph

import (
    "testing"
    "reflect"
)

// tests the edge methods
func TestEdge(t *testing.T) {

    idProvider := idProvider(0)
    vStart := vertex{id: idProvider.NewId()}
    vEnd := vertex{id: idProvider.NewId()}
    weight := 1.23
    i := uint(234)
    e := edge{
        id: id(i),
        start: vStart,
        end: vEnd,
        weight: weight,
    }

    // test GetId()
    if e.GetId() != i {
        t.Errorf("Edge ID should be %d, got %d.", i, e.GetId())
    }

    // test GetStartVertex()
    if !reflect.DeepEqual(e.GetStartVertex(), vStart) {
        t.Error("Could not retrieve start vertex.")
    }

    // test GetEndVertex()
    if !reflect.DeepEqual(e.GetEndVertex(), vEnd) {
        t.Error("Could not retrieve end vertex.")
    }

    // test GetWeight()
    if e.GetWeight() != weight {
        t.Errorf("Edge weight should be %f, got %f.", weight, e.GetWeight())
    }

    // test SetWeight()
    weight2 := weight + 123
    e.SetWeight(weight2)
    if e.GetWeight() != weight2 {
        t.Errorf("Edge weight should be %f, got %f.", weight2, e.GetWeight())
    }
}

// tests the edges methods
func TestEdges(t *testing.T) {

    eIdProvider := idProvider(0)
    vIdProvider := idProvider(0)
    e := edges{}

    // test Add()
    vStart := vertex{id: vIdProvider.NewId()}
    var es [10]*edge
    var eid [10]id
    for i := 0; i < 10; i++ {
        eid[i] = eIdProvider.NewId()
        es[i] = &edge{
            id: eid[i],
            start: vStart,
            end: &vertex{id: vIdProvider.NewId()},
        }
        e.add(es[i])
    }

    // test Count()
    if e.Count() != 10 {
        t.Errorf("Edges should contain 10 elements, got %d.", e.Count())
    }

    // test Get()
    for i, id := range eid {
        if !reflect.DeepEqual(e.Get(uint(id)), es[i]) {
            t.Error("Could not recover edges.")
        }
    }

    // test All()
    n := 0;
    for _, ex := range e.All() {
        for i, id := range eid {
            if ex.GetId() == uint(id) {
                if !reflect.DeepEqual(ex, es[i]) {
                    t.Error("Could not recover edges.")
                }
                n++
            }
        }
    }
    if n != 10 {
        t.Errorf("Edges should contain 10 elements, got %d.", n)
    }
}

// tests the merged edges methods
func TestMergedEdges(t *testing.T) {

    eIdProvider := idProvider(0)
    vIdProvider := idProvider(0)

    var el [3]*edges
    for i := 0; i < 3; i++ {
        el[i] = &edges{}
    }

    vStart := vertex{id: vIdProvider.NewId()}
    var es [10]*edge
    var eid [10]id
    for i := 0; i < 10; i++ {
        eid[i] = eIdProvider.NewId()
        es[i] = &edge{
            id: eid[i],
            start: vStart,
            end: &vertex{id: vIdProvider.NewId()},
        }
        el[i % 3].add(es[i])
    }

    // test merge()
    e := mergedEdges{}
    for _, ed := range el {
        e.merge(ed)
    }

    // test Count()
    if e.Count() != 10 {
        t.Errorf("Edges should contain 10 elements, got %d.", e.Count())
    }

    // test Get()
    for i, id := range eid {
        if !reflect.DeepEqual(e.Get(uint(id)), es[i]) {
            t.Error("Could not recover edges.")
        }
    }

    // test All()
    n := 0;
    for _, ex := range e.All() {
        for i, id := range eid {
            if ex.GetId() == uint(id) {
                if !reflect.DeepEqual(ex, es[i]) {
                    t.Error("Could not recover edges.")
                }
                n++
            }
        }
    }
    if n != 10 {
        t.Errorf("Edges should contain 10 elements, got %d.", n)
    }
}
