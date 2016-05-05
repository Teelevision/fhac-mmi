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

    // get the number of vertices
    num := graph.GetVertices().Count()
    if num <= 1 || num > 15 {
        return 0
    }

    type vertex struct {
        graphLib.VertexInterface
        index     int
        distances [15]float64
    }

    // create vertices
    vertices := make([]*vertex, num)
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
    start := vertices[0]
    startIndex := start.index

    // performs brute force to find the length of the shortest hamilton circle
    var helper func(*vertex, []*vertex) float64
    helper = func(front *vertex, rest []*vertex) float64 {

        rest0, rest0b := rest[0], (*vertex)(nil)
        restFrom1 := rest[1:]

        // last element
        if len(rest) == 1 {
            return front.distances[rest0.index] + rest0.distances[startIndex]
        }

        // when not changing the order
        length := helper(rest0, restFrom1)
        length += front.distances[rest0.index]

        // combinations of changing the order
        for i := 1; i < len(rest); i++ {

            // change order
            rest0b, rest[i] = rest[i], rest0

            // recursion
            lengthCandidate := helper(rest0b, restFrom1)
            lengthCandidate += front.distances[rest0b.index]

            if lengthCandidate < length {
                length = lengthCandidate
            }

            // change back
            rest[i] = rest0b
        }

        return length
    }

    // perform brute force
    length := helper(start, vertices[1:])

    // returns the length of the shortest hamilton circle
    return length
}
