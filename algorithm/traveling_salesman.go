package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
    "math"
)

// simple wrapper
func (this Graph) TravelingSalesmanBruteForce() float64 {
    return TravelingSalesmanBruteForce(this)
}

// returns the length of the shortest hamilton circle by brute force
func TravelingSalesmanBruteForce(graph Graph) float64 {

    // get the number of vertices
    num := int(graph.GetVertices().Count())
    if num <= 1 || num > 15 {
        return 0
    }

    type vertex struct {
        graphLib.VertexInterface
        index     int
        distances [15]float64
    }

    // create vertices
    var vertices [15]*vertex
    for i, v := range graph.GetVertices().All() {

        vertices[i] = &vertex{
            VertexInterface: v,
            index: i,
        }

        for i2, v2 := range vertices[:i] {
            w := graph.getWeightBetween(v, v2.VertexInterface)
            vertices[i].distances[i2] = w
            v2.distances[i] = w
        }

    }

    // start
    startDist := vertices[0].distances

    // performs brute force to find the length of the shortest hamilton circle
    type helperFunc func(int, *vertex, float64)
    var helpers [15]helperFunc
    length := math.MaxFloat64

    // end helper
    helpers[num - 1] = func(n int, front *vertex, currentLength float64) {
        rest0index := vertices[n].index
        l := currentLength + front.distances[rest0index] + startDist[rest0index]
        if l < length {
            length = l
        }
    }

    // front to pre-end helpers
    for i := num - 2; i >= 0; i-- {
        helpers[i] = func(n int, front *vertex, currentLength float64) {

            n1, rest0, rest0b, frontDist := n + 1, vertices[n], (*vertex)(nil), &front.distances
            next := helpers[n1]

            // when not changing the order
            next(n1, rest0, currentLength + frontDist[rest0.index])

            // combinations of changing the order
            for i := n1; i < num; i++ {

                // change order
                rest0b, vertices[i] = vertices[i], rest0

                // recursion
                next(n1, rest0b, currentLength + frontDist[rest0b.index])

                // change back
                vertices[i] = rest0b
            }
        }

    }

    // perform brute force
    helpers[0](1, vertices[0], 0.0)

    // returns the length of the shortest hamilton circle
    return length
}
