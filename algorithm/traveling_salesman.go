package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
    "gopkg.in/cheggaaa/pb.v1"
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

    // progress bar
    depth := int(num) / 2;
    p := pb.StartNew(travelingSalesmanBruteForceHelperCalls(int(num), depth))

    // performs brute force to find the length of the shortest hamilton circle
    var helper func(int, *pb.ProgressBar, *vertex, []*vertex) (*vertex, float64)
    helper = func(depth int, p *pb.ProgressBar, front *vertex, rest []*vertex) (*vertex, float64) {

        // last element
        if len(rest) == 1 {
            return rest[0], front.distances[rest[0].index]
        }

        // when not changing the order
        lastVertex, length := helper(depth - 1, p, rest[0], rest[1:])
        if front != nil {
            length += front.distances[rest[0].index]
        } else {
            length += rest[0].distances[lastVertex.index]
        }

        // combinations of changing the order
        for i := 1; i < len(rest); i++ {

            // change order
            rest[0], rest[i] = rest[i], rest[0]

            // recursion
            lastVertexCandidat, lengthCandidat := helper(depth - 1, p, rest[0], rest[1:])
            if front != nil {
                lengthCandidat += front.distances[rest[0].index]
            } else {
                lengthCandidat += rest[0].distances[lastVertexCandidat.index]
            }
            if lengthCandidat < length {
                lastVertex, length = lastVertexCandidat, lengthCandidat
            }

            // change back
            rest[0], rest[i] = rest[i], rest[0]
        }

        // progress bar
        if depth > 0 {
            p.Increment()
        }

        return lastVertex, length
    }

    // perform brute force
    _, length := helper(depth, p, nil, vertices)

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