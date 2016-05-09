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
    // this algorithm only works with 5 to 15 vertices
    if num <= 4 || num > 15 {
        return 0
    }
    numMinus1 := num - 1

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
    type helperFunc func(int, *vertex, *vertex, float64)
    var helpers [15]helperFunc
    length := math.MaxFloat64

    // end helper
    helpers[num - 3] = func(n int, front, end *vertex, currentLength float64) {
        rest0index := vertices[n].index
        l := currentLength + front.distances[rest0index] + end.distances[rest0index]
        if l < length {
            length = l
        }
    }

    // front to pre-end helpers
    for i := num - 4; i >= 1; i-- {
        helpers[i] = func(n int, front, end *vertex, currentLength float64) {

            n1, rest0, rest0Temp, frontDist := n + 1, vertices[n], (*vertex)(nil), &front.distances
            next := helpers[n]

            // when not changing the order
            next(n1, rest0, end, currentLength + frontDist[rest0.index])

            // combinations of changing the order
            for i := n1; i < numMinus1; i++ {

                // change order
                rest0Temp, vertices[i] = vertices[i], rest0

                // recursion
                next(n1, rest0Temp, end, currentLength + frontDist[rest0Temp.index])

                // change back
                vertices[i] = rest0Temp
            }
        }
    }

    currentLength, rest0, rest0Temp, end, endTemp, frontDist := 0.0, vertices[1], (*vertex)(nil), vertices[numMinus1], (*vertex)(nil), &vertices[0].distances
    next := helpers[1]

    // combinations of changing the order
    for i := 1; i < numMinus1; i++ {

        // change order
        rest0Temp, vertices[i] = vertices[i], rest0

        for j := i + 1; j < num; j++ {

            // change order
            endTemp, vertices[j] = vertices[j], end

            // recursion
            next(2, rest0Temp, endTemp, currentLength + frontDist[rest0Temp.index] + startDist[endTemp.index])

            // change back
            vertices[j] = endTemp
        }

        // change back
        vertices[i] = rest0Temp
    }

    // returns the length of the shortest hamilton circle
    return length
}
