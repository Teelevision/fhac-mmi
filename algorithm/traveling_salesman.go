package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
    "gopkg.in/cheggaaa/pb.v1"
    "math"
)

// simple wrapper
func (this Graph) TravelingSalesmanBruteForce() float64 {
    return TravelingSalesmanBruteForce(this)
}

// returns the length of the shortest hamilton circle by brute force
func TravelingSalesmanBruteForce(graph Graph) float64 {

    // wrapper type for vertices
    type vertex struct {
        graphLib.VertexInterface
        index     int
        distances [15]float64 // distances to other vertices, by index
    }

    // get the number of vertices
    num := int(graph.GetVertices().Count())
    if num == 0 {
        return 0
    }

    // create vertices
    vertices := make([]*vertex, num)
    for i, v := range graph.GetVertices().All() {

        vertices[i] = &vertex{
            VertexInterface: v,
            index: i,
            //distances: make([]float64, num),
        }

        // for all previous vertices, add distance
        for i2, v2 := range vertices[:i] {
            w := graph.getWeightBetween(v, v2.VertexInterface)
            vertices[i].distances[i2] = w
            v2.distances[i] = w
        }

    }

    factorial := func(n int) (r int) {
        for r = 1; n > 1; n-- {
            r *= n
        }
        return r
    }

    // progress bar
    //depth := int(num) / 2;
    p := pb.StartNew(factorial(len(vertices)))

    end := num - 1
    swap := make([]int, num - 1)
    l := num - 2
    length := math.MaxFloat64
    for a := factorial(len(vertices)); a > 0; a-- {

        // swap
        for i, c := range swap {
            if c > 0 {
                j := i + c
                vertices[i], vertices[j] = vertices[j], vertices[i]
            }
        }

        baseLength := 0.0
        for i, v := range vertices {
            if i == end {
                baseLength += v.distances[vertices[0].index]
            } else {
                baseLength += v.distances[vertices[i + 1].index]
            }
        }
        if baseLength < length {
            length = baseLength
        }

        // swap back
        for i := len(swap) - 1; i >= 0; i-- {
            c := swap[i]
            if c > 0 {
                vertices[i], vertices[i + c] = vertices[i + c], vertices[i]
            }
        }

        // increment
        for i := len(swap) - 1; i >= 0; i-- {
            swap[i]++
            if swap[i] >= (l - i + 2) {
                swap[i] = 0
            } else {
                break
            }
        }

        p.Increment()
    }

    // performs brute force to find the length of the shortest hamilton circle
    var helper func(int, *pb.ProgressBar, *vertex, []*vertex) (*vertex, float64)
    helper = func(depth int, p *pb.ProgressBar, front *vertex, rest []*vertex) (*vertex, float64) {
        rest0 := rest[0]
        rest1 := rest[1:]

        // last element
        if len(rest) == 1 {
            return rest0, front.distances[rest0.index]
        }

        // when not changing the order
        lastVertex, length := helper(depth - 1, p, rest0, rest1)
        if front != nil {
            length += front.distances[rest0.index]
        } else {
            length += rest0.distances[lastVertex.index]
        }

        // combinations of changing the order
        for i, resti := range rest1 {

            // change order
            rest0, rest[i] = resti, rest0

            // recursion
            lastVertexCandidat, lengthCandidat := helper(depth - 1, p, rest0, rest1)
            if front != nil {
                lengthCandidat += front.distances[rest0.index]
            } else {
                lengthCandidat += rest0.distances[lastVertexCandidat.index]
            }
            if lengthCandidat < length {
                lastVertex, length = lastVertexCandidat, lengthCandidat
            }

            // change back
            rest0, rest[i] = rest[i], rest0
        }

        // progress bar
        if depth > 0 {
            p.Increment()
        }

        return lastVertex, length
    }

    // perform brute force
    //_, length := helper(depth, p, nil, vertices)

    // print that we are done
    p.FinishPrint("Done.")

    // returns the length of the shortest hamilton circle
    return length
}

// calculates the number of helper function calls until the given depths
func travelingSalesmanBruteForceHelperCalls(n, depth int) int {
    if n == depth {
        return 0
    }
    return 1 + n * travelingSalesmanBruteForceHelperCalls(n - 1, depth)
}