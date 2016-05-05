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
    startIndex := 0

    // performs brute force to find the length of the shortest hamilton circle
    var helper func(*vertex, int, float64)
    length := math.MaxFloat64
    helper = func(front *vertex, n int, currentLength float64) {

        n1 := n+1
        rest0, rest0b := vertices[n], (*vertex)(nil)

        // last element
        if num == n1 {
            l := currentLength + front.distances[rest0.index] + rest0.distances[startIndex]
            if l < length {
                length = l
            }
            return
        }

        // when not changing the order
        helper(rest0, n1, currentLength + front.distances[rest0.index])

        // combinations of changing the order
        for i := n1; i < num; i++ {

            // change order
            rest0b, vertices[i] = vertices[i], rest0

            // recursion
            helper(rest0b, n1, currentLength + front.distances[rest0b.index])

            // change back
            vertices[i] = rest0b
        }
    }

    // perform brute force
    helper(vertices[0], 1, 0.0)

    // returns the length of the shortest hamilton circle
    return length
}
