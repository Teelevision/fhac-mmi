package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
    "gopkg.in/cheggaaa/pb.v1"
    "time"
)

// simple wrapper
func (this Graph) TravelingSalesmanBruteForce() float64 {
    return TravelingSalesmanBruteForce(this)
}

// returns the length of the shortest hamilton circle by brute force
func TravelingSalesmanBruteForce(graph Graph) float64 {

    startTime := time.Now()

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

    var fakDepth func(int, int) int
    fakDepth = func(n, depth int) int {
        if n == depth {
            return 1
        }
        return n * fakDepth(n - 1, depth)
    }

    depth := int(num) / 2;
    num1 := int(num) - 1

    calls := 0
    for d := num1; d >= depth; d-- {
        calls += fakDepth(num1, d)
    }

    // progress bar
    var p *pb.ProgressBar
    pf := 0

    // performs brute force to find the length of the shortest hamilton circle
    var helper func(int, int, *vertex, []*vertex) (float64)
    helper = func(depth int, veryFrontIndex int, front *vertex, rest []*vertex) (float64) {

        rest0 := rest[0]

        // last element
        if len(rest) == 1 {
            return front.distances[rest0.index] + rest0.distances[veryFrontIndex]
        }

        // when not changing the order
        length := helper(depth - 1, veryFrontIndex, rest0, rest[1:])
        length += front.distances[rest0.index]


        // combinations of changing the order
        for i := 1; i < len(rest); i++ {

            // change order
            rest0, rest[i] = rest[i], rest0

            // recursion
            lengthCandidat := helper(depth - 1, veryFrontIndex, rest0, rest[1:])
            lengthCandidat += front.distances[rest0.index]

            if lengthCandidat < length {
                length = lengthCandidat
            }

            // change back
            rest0, rest[i] = rest[i], rest0
        }

        // progress bar
        if depth > 0 {
            t := time.Now().Sub(startTime).Seconds()
            if t < 5.0 {
                pf++
            } else if p == nil {
                p = pb.StartNew(calls)
                p.Add(pf)
            } else {
                p.Increment()
            }
        }

        return length
    }

    // perform brute force
    length := helper(depth, vertices[0].index, vertices[0], vertices[1:])

    // print that we are done
    if p != nil {
        p.FinishPrint("Done.")
    }

    // returns the length of the shortest hamilton circle
    return length
}
