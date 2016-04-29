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

    // get the number of vertices
    num := graph.GetVertices().Count()
    if num == 0 {
        return 0
    }

    // copy slice of vertices, because we are going to write to it
    vertices := make([]graphLib.VertexInterface, num)
    copy(vertices, graph.GetVertices().All())

    // make distance map that contains every distance between every two vertices
    numEdges := graph.GetEdges().Count()
    dm := make(travelingSalesmanBruteForceDistanceMap, numEdges)
    for _, v := range vertices {
        dm[v] = make(map[graphLib.VertexInterface]float64, numEdges)
    }

    // fill distance map
    for _, e := range graph.GetEdges().All() {
        v1, v2, w := e.GetStartVertex(), e.GetEndVertex(), e.GetWeight()
        dm[v1][v2], dm[v2][v1] = w, w
    }

    // progress bar
    depth := int(num) / 2;
    p := pb.StartNew(travelingSalesmanBruteForceHelperCalls(int(num), depth))

    // perform brute force
    _, length := travelingSalesmanBruteForceHelper(depth, p, dm, nil, vertices)

    // print that we are done
    p.FinishPrint("Done.")

    // returns the length of the shortest hamilton circle
    return length
}

// distance map used by the traveling salesman brute force algorithm
type travelingSalesmanBruteForceDistanceMap map[graphLib.VertexInterface]map[graphLib.VertexInterface]float64

// distance between two vertices or 0 if the first is nil
func (this travelingSalesmanBruteForceDistanceMap) dist(v1, v2 graphLib.VertexInterface) float64 {
    if v1 == nil {
        return 0
    }
    return this[v1][v2]
}

// calculates the number of helper function calls until the given depths
func travelingSalesmanBruteForceHelperCalls(n, depth int) int {
    if n == depth {
        return 0
    }
    return 1 + n * travelingSalesmanBruteForceHelperCalls(n - 1, depth)
}

// performs brute force to find the length of the shortest hamilton circle
func travelingSalesmanBruteForceHelper(depth int, p *pb.ProgressBar, dm travelingSalesmanBruteForceDistanceMap, front graphLib.VertexInterface, rest []graphLib.VertexInterface) (graphLib.VertexInterface, float64) {

    // last element
    if len(rest) == 1 {
        return rest[0], dm.dist(front, rest[0])
    }

    // when not changing the order
    lastVertex, length := travelingSalesmanBruteForceHelper(depth - 1, p, dm, rest[0], rest[1:])
    length += dm.dist(front, rest[0])
    if front == nil {
        length += dm.dist(rest[0], lastVertex)
    }

    // combinations of changing the order
    for i := 1; i < len(rest); i++ {

        // change order
        rest[0], rest[i] = rest[i], rest[0]

        // recursion
        lastVertexCandidat, lengthCandidat := travelingSalesmanBruteForceHelper(depth - 1, p, dm, rest[0], rest[1:])
        lengthCandidat += dm.dist(front, rest[0])
        if front == nil {
            lengthCandidat += dm.dist(rest[0], lastVertexCandidat)
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