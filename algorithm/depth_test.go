package algorithm

import (
    "testing"
    "github.com/teelevision/fhac-mmi/graph"
    "reflect"
)

// test the depth-first search
func TestDepthFirstSearch(t *testing.T) {

    g := graph.UndirectedGraph()
    a := Graph{g}

    // add 10 vertices
    var v [10]graph.VertexInterface
    for i := 0; i < 10; i++ {
        v[i] = g.NewVertex()
    }

    // test function
    test := func(start graph.VertexInterface, orderDirected []graph.VertexInterface, orderUndirected []graph.VertexInterface) {
        g.SetDirected(true)
        vertices := a.DepthFirstSearch(start)
        if !reflect.DeepEqual(vertices, orderDirected) {
            t.Errorf("Expected\n    %v,\ngot %v.", orderDirected, vertices)
        }
        g.SetDirected(false)
        vertices = a.DepthFirstSearch(start)
        if !reflect.DeepEqual(vertices, orderUndirected) {
            t.Errorf("Expected\n    %v,\ngot %v.", orderUndirected, vertices)
        }
    }

    // add edge function
    edge := func(i, j uint) {
        g.NewEdge(v[i], v[j])
    }

    // test (direction from lower to higher):
    // [0]---(1)    (6)--(7)  (9)
    //  | \    \     |
    // (2)(3)--(4)  (5)--(8)
    edge(0, 1)
    edge(0, 2)
    edge(0, 3)
    edge(1, 4)
    edge(3, 4)
    edge(5, 6)
    edge(5, 8)
    edge(6, 7)
    test(v[0],
        []graph.VertexInterface{v[0], v[1], v[4], v[2], v[3]},
        []graph.VertexInterface{v[0], v[1], v[4], v[3], v[2]})

    // test (direction from lower to higher):
    // [0]---(1)    (6)--(7)  (9)
    //  | \    \     |
    // (2)(3)--(4)--(5)--(8)
    edge(4, 5)
    test(v[0],
        []graph.VertexInterface{v[0], v[1], v[4], v[5], v[6], v[7], v[8], v[2], v[3]},
        []graph.VertexInterface{v[0], v[1], v[4], v[5], v[6], v[7], v[8], v[3], v[2]})

    // test (direction from lower to higher):
    // [0]---(1)<---(6)--(7)  (9)
    //  | \    \     |
    // (2)(3)--(4)--(5)--(8)
    edge(6, 1)
    test(v[0],
        []graph.VertexInterface{v[0], v[1], v[4], v[5], v[6], v[7], v[8], v[2], v[3]},
        []graph.VertexInterface{v[0], v[1], v[4], v[5], v[6], v[7], v[8], v[3], v[2]})

    // test (direction from lower to higher):
    // (0)---(1)<---(6)--(7)  [9]
    //  | \    \     |
    // (2)(3)--(4)--(5)--(8)
    test(v[9],
        []graph.VertexInterface{v[9]},
        []graph.VertexInterface{v[9]})

    // test (direction from lower to higher):
    // (0)---(1)<---[6]--(7)  (9)
    //  | \    \     |
    // (2)(3)--(4)--(5)--(8)
    test(v[6],
        []graph.VertexInterface{v[6], v[7], v[1], v[4], v[5], v[8]},
        []graph.VertexInterface{v[6], v[7], v[1], v[4], v[5], v[8], v[3], v[0], v[2]})
}
