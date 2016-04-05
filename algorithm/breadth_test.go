package algorithm

import (
    "testing"
    "github.com/teelevision/fhac-mmi/graph"
    "reflect"
)

// test the Breitensuche
func TestBreitensuche(t *testing.T) {

    g := graph.UndirectedGraph()
    a := Graph{g}

    // add 10 vertices
    var v [10]graph.VertexInterface
    for i := 0; i < 10; i++ {
        v[i] = g.NewVertex()
    }

    test := func(start graph.VertexInterface, orderDirected []graph.VertexInterface, orderUndirected []graph.VertexInterface) {
        g.SetDirected(true)
        vertices := a.BreadthFirstSearch(start)
        if !reflect.DeepEqual(vertices, orderDirected) {
            t.Errorf("Expected\n    %v,\ngot %v.", orderDirected, vertices)
        }
        g.SetDirected(false)
        vertices = a.BreadthFirstSearch(start)
        if !reflect.DeepEqual(vertices, orderUndirected) {
            t.Errorf("Expected\n    %v,\ngot %v.", orderUndirected, vertices)
        }
    }

    // test (direction from lower to higher):
    // [0]---(1)    (6)--(7)  (9)
    //  | \    \     |
    // (2)(3)--(4)  (5)--(8)
    g.NewEdge(v[0], v[1])
    g.NewEdge(v[0], v[2])
    g.NewEdge(v[0], v[3])
    g.NewEdge(v[1], v[4])
    g.NewEdge(v[3], v[4])
    g.NewEdge(v[5], v[6])
    g.NewEdge(v[5], v[8])
    g.NewEdge(v[6], v[7])
    test(v[0],
        []graph.VertexInterface{v[0], v[1], v[2], v[3], v[4]},
        []graph.VertexInterface{v[0], v[1], v[2], v[3], v[4]})

    // test (direction from lower to higher):
    // [0]---(1)    (6)--(7)  (9)
    //  | \    \     |
    // (2)(3)--(4)--(5)--(8)
    g.NewEdge(v[4], v[5])
    test(v[0],
        []graph.VertexInterface{v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[8], v[7]},
        []graph.VertexInterface{v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[8], v[7]})

    // test (direction from lower to higher):
    // [0]---(1)<---(6)--(7)  (9)
    //  | \    \     |
    // (2)(3)--(4)--(5)--(8)
    g.NewEdge(v[6], v[1])
    test(v[0],
        []graph.VertexInterface{v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[8], v[7]},
        []graph.VertexInterface{v[0], v[1], v[2], v[3], v[4], v[6], v[5], v[7], v[8]})

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
        []graph.VertexInterface{v[6], v[7], v[1], v[5], v[4], v[0], v[8], v[3], v[2]})
}
