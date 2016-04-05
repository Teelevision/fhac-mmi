package algorithm

import (
    "testing"
    "github.com/teelevision/fhac-mmi/graph"
)

// test the calculation of connected components
func TestConnectedComponents(t *testing.T) {

    g := graph.UndirectedGraph()
    g.SetDirected(true)
    a := Graph{g}

    // test function
    test := func(num uint) {
        if n := a.GetNumConnectedComponents(); n != num {
            t.Errorf("Expected %d connected components, got %d.", num, n)
        }
    }

    // test: empty
    test(0)

    // add 10 vertices
    var v [10]graph.VertexInterface
    for i := 0; i < 10; i++ {
        v[i] = g.NewVertex()
    }

    // add edge function
    edge := func(i, j uint) {
        g.NewEdge(v[i], v[j])
    }

    // test: no edges
    test(10)

    // test:
    // (0)---(1)     (6)--(7)  (9)
    //  |             |
    // (2) (3)--(4)  (5)--(8)
    edge(0, 1)
    edge(0, 2)
    edge(3, 4)
    edge(5, 6)
    edge(5, 8)
    edge(6, 7)
    test(4)

    // test:
    // (0)---(1)    (6)--(7)  (9)
    //  |    / \     |    |
    // (2) (3)-(4)  (5)--(8)
    edge(1, 3)
    edge(1, 4)
    edge(7, 8)
    test(3)

    // test:
    // (0)---(1)    (6)--(7)  (9)
    //  |    / \     |    |
    // (2) (3)-(4)--(5)--(8)
    edge(4, 5)
    test(2)

    // test:
    // (0)---(1)    (6)--(7)--(9)
    //  |    / \     |    |
    // (2) (3)-(4)--(5)--(8)
    edge(7, 9)
    test(1)

}
