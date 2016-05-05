package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
)

// simple wrapper
func (this Graph) TravelingSalesmanBruteForce() float64 {
    return TravelingSalesmanBruteForce(this)
}

// returns the length of the shortest hamilton circle by brute force
func TravelingSalesmanBruteForce(graph Graph) float64 {

    type vertex struct {
        graphLib.VertexInterface
        index     int
        distances []float64
    }

    // get the number of vertices
    num := graph.GetVertices().Count()
    if num == 0 {
        return 0
    }

    // create vertices
    vertices := make([]*vertex, num)
    for i, v := range graph.GetVertices().All() {

        vertices[i] = &vertex{
            VertexInterface: v,
            index: i,
            distances: make([]float64, num),
        }

        for i2, v2 := range vertices[:i] {
            w := graph.getWeightBetween(v, v2.VertexInterface)
            vertices[i].distances[i2] = w
            v2.distances[i] = w
        }

    }

    // performs brute force to find the length of the shortest hamilton circle
    var helper func(int, *vertex, []*vertex) float64
    helper = func(veryFrontIndex int, front *vertex, rest []*vertex) float64 {

        rest0 := rest[0]

        // last element
        if len(rest) == 1 {
            return front.distances[rest0.index] + rest0.distances[veryFrontIndex]
        }

        // when not changing the order
        length := helper(veryFrontIndex, rest0, rest[1:])
        length += front.distances[rest0.index]

        // combinations of changing the order
        for i := 1; i < len(rest); i++ {

            // change order
            rest0, rest[i] = rest[i], rest0

            // recursion
            lengthCandidat := helper(veryFrontIndex, rest0, rest[1:])
            lengthCandidat += front.distances[rest0.index]

            if lengthCandidat < length {
                length = lengthCandidat
            }

            // change back
            rest0, rest[i] = rest[i], rest0
        }

        return length
    }

    // perform brute force
    start := vertices[0]
    length := helper(start.index, start, vertices[1:])

    // returns the length of the shortest hamilton circle
    return length
}
