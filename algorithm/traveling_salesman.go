package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
    "math"
)

// the maximum supported number of vertices in the traveling salesman brute force algorithm
const tsbfMax = 15

// simple wrapper
func (this Graph) TravelingSalesmanBruteForce(branchAndBound bool) float64 {
    return TravelingSalesmanBruteForce(this, branchAndBound)
}

// returns the length of the shortest hamilton circle by brute force
func TravelingSalesmanBruteForce(graph Graph, branchAndBound bool) float64 {

    // ----------------------------------------------
    // check pre-requirements
    // ----------------------------------------------

    // get the number of vertices
    num := int(graph.GetVertices().Count())

    // this algorithm only works with 5 to 15 vertices
    if num <= 4 || num > tsbfMax {
        return 0
    }

    // ----------------------------------------------
    // preparations
    // ----------------------------------------------

    // we access this one often, so better calculate it only once
    lastVertex := num - 1

    // use own vertex structure that has an index which refers to its position in the vertices array
    // it also contains an array to each other position containing the distance between the vertices
    type vertex struct {
        graphLib.VertexInterface
        index     int
        distances [tsbfMax]float64
    }

    // create vertices and calculate distance once here instead of doing it multiple times later
    var vertices [tsbfMax]*vertex
    for i, v := range graph.GetVertices().All() {

        // create vertex
        vertices[i] = &vertex{
            VertexInterface: v,
            index: i,
        }

        // add distances for all previously created vertices
        for j, v2 := range vertices[:i] {
            w := graph.getWeightBetween(v, v2.VertexInterface)
            // make the distances two-way
            vertices[i].distances[j] = w
            v2.distances[i] = w
        }

    }

    // the shortest circle's length will be stored in here
    length := math.MaxFloat64

    // ----------------------------------------------
    // build helpers
    // ----------------------------------------------
    // There will be 2 different helpers:
    // * The intermediate helper will be called recursively and ...
    //   ... creates all possible permutations of the remaining vertices.
    // * The end helper will be called at the end of the recursion and ...
    //   ... updates the length variable if a new shortest circle was found.
    // ----------------------------------------------

    // there will be 2 different helper functions. both will be stored in this array
    var helpers [tsbfMax - 2]func(int, *vertex, *vertex, float64)

    // end helper
    helpers[num - 3] = func(n int, front, end *vertex, currentLength float64) {
        rest0index := vertices[n].index
        l := currentLength + front.distances[rest0index] + end.distances[rest0index]
        if l < length {
            length = l
        }
    }

    // intermediate helper
    for i := num - 4; i >= 1; i-- {
        if branchAndBound {
            // when using branch and bound, check the length on each recursion level
            // abort if already longer than the shortest known circle
            helpers[i] = func(n int, front, end *vertex, currentLength float64) {

                n1, rest0, rest0Temp, frontDist, weightTmp := n + 1, vertices[n], (*vertex)(nil), &front.distances, 0.0
                next := helpers[n]

                // when not changing the order
                weightTmp = currentLength + frontDist[rest0.index]
                if weightTmp < length {
                    next(n1, rest0, end, weightTmp)
                }

                // combinations of changing the order
                for i := n1; i < lastVertex; i++ {

                    // change order
                    rest0Temp, vertices[i] = vertices[i], rest0

                    // recursion
                    weightTmp = currentLength + frontDist[rest0Temp.index]
                    if weightTmp < length {
                        next(n1, rest0Temp, end, weightTmp)
                    }

                    // change back
                    vertices[i] = rest0Temp
                }
            }
        } else {
            helpers[i] = func(n int, front, end *vertex, currentLength float64) {

                n1, rest0, rest0Temp, frontDist := n + 1, vertices[n], (*vertex)(nil), &front.distances
                next := helpers[n]

                // when not changing the order
                next(n1, rest0, end, currentLength + frontDist[rest0.index])

                // combinations of changing the order
                for i := n1; i < lastVertex; i++ {

                    // change order
                    rest0Temp, vertices[i] = vertices[i], rest0

                    // recursion
                    next(n1, rest0Temp, end, currentLength + frontDist[rest0Temp.index])

                    // change back
                    vertices[i] = rest0Temp
                }
            }
        }
    }

    // ----------------------------------------------
    // Strategy
    // ----------------------------------------------
    // 1. Pick 0 as start vertex.
    // 2. Pick every _set_ of two vertices and put one left and one right of the start vertex
    // 3. Use every permutation of the remaining vertices to close the circle and calculate the length.
    //    Use recursion for the permutation.
    // ----------------------------------------------

    // ----------------------------------------------
    // 1. Pick 0 as start vertex.
    // ----------------------------------------------
    startDist := vertices[0].distances

    // ----------------------------------------------
    // 2. Pick every _set_ of two vertices and put one left and one right of the start vertex
    // ----------------------------------------------

    // when swapping vertices, use this variables for the front and end to read from (faster than array access)
    rest0, end := vertices[1], vertices[lastVertex]

    // when swapping vertices, use this variables for the front and end to write to (faster than array access)
    var rest0Temp, endTemp *vertex

    // select first element of the set
    for i := 1; i < lastVertex; i++ {

        // swap. we do not actually need to swap, ...
        // ... because on the next recursion levels only the indexes 2 upwards are accessed
        rest0Temp, vertices[i] = vertices[i], rest0

        // select second element of the set that is greater than the first one
        for j := i + 1; j <= lastVertex; j++ {

            // swap. we do not actually need to swap, ...
            // ... because on the next recursion levels the last vertex is never accessed
            endTemp, vertices[j] = vertices[j], end

            // ----------------------------------------------
            // 3. Use every permutation of the remaining vertices to close the circle and calculate the length.
            // ----------------------------------------------
            helpers[1](2, rest0Temp, endTemp, startDist[i] + startDist[j])

            // swap back. since we did not really swap above, this is as simple as an assignment
            vertices[j] = endTemp
        }

        // swap back. since we did not really swap above, this is as simple as an assignment
        vertices[i] = rest0Temp
    }

    // returns the length of the shortest hamilton circle
    return length
}
