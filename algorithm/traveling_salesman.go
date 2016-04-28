package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
)

// simple wrapper
func (this Graph) TravelingSalesmanBruteForce() float64 {
    return TravelingSalesmanBruteForce(this)
}

type distanceMap map[graphLib.VertexInterface]map[graphLib.VertexInterface]float64

func (this distanceMap) dist(v1, v2 graphLib.VertexInterface) float64 {
    if v1 == nil {
        return 0
    }
    return this[v1][v2]
}

// returns the length of the shortest hamilton circle by brute force
func TravelingSalesmanBruteForce(graph Graph) float64 {

    num := graph.GetVertices().Count()
    if num == 0 {
        return 0
    }

    vertices := make([]graphLib.VertexInterface, num)
    copy(vertices, graph.GetVertices().All())

    // make distance map
    numEdges := graph.GetEdges().Count()
    dm := make(distanceMap, numEdges)
    for _, v := range vertices {
        dm[v] = make(map[graphLib.VertexInterface]float64, numEdges)
    }
    for _, e := range graph.GetEdges().All() {
        v1, v2, w := e.GetStartVertex(), e.GetEndVertex(), e.GetWeight()
        dm[v1][v2], dm[v2][v1] = w, w
    }

    _, length := foobar(dm, nil, vertices)

    return length
}

func foobar(dm distanceMap, front graphLib.VertexInterface, rest []graphLib.VertexInterface) (graphLib.VertexInterface, float64) {

    /*if len(rest) > 8 && front != nil {
        for i := 10 - len(rest); i > 0; i-- {
            fmt.Print(".")
        }
        fmt.Print(front.GetId())
        for _, r := range rest {
            fmt.Print(r.GetId())
        }
        fmt.Println()
    }*/

    // last element
    if len(rest) == 1 {
        return rest[0], dm.dist(front, rest[0])
    }

    // when not changing the order
    first := rest[0]
    last, l := foobar(dm, rest[0], rest[1:])
    l += dm.dist(front, rest[0])

    // combinations of changing the order
    for i := 1; i < len(rest); i++ {

        // change order
        rest[0], rest[i] = rest[i], rest[0]

        // recursion
        la, l2 := foobar(dm, rest[0], rest[1:])
        l2 += dm.dist(front, rest[0])
        if l2 < l {
            last, l, first = la, l2, rest[0]
        }

        // change back
        rest[0], rest[i] = rest[i], rest[0]
    }

    if front == nil {
        l += dm.dist(first, last)
    }
    return last, l
}