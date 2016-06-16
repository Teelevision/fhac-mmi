package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
    "math"
    "container/heap"
    "fmt"
)

// simple wrapper
func (this Graph) ShortestPathsDijkstra(start, end graphLib.VertexInterface) {
    ShortestPathsDijkstra(this, start, end)
}

//
func ShortestPathsDijkstra(graph Graph, start, end graphLib.VertexInterface) {

    // number of vertices
    num := graph.GetVertices().Count()

    // queue
    q := make(shortestPathQueue, num)
    // map to map vertices to the objects that are used here
    m := make(map[graphLib.VertexInterface]*shortestPathVertex, num)
    // create objects
    for i, v := range graph.GetVertices().All() {
        vertex := &shortestPathVertex{
            VertexInterface: v,
            prev: nil,
            distance: math.MaxFloat64,
            index: i,
        }
        m[v], q[i] = vertex, vertex
    }

    // start
    m[start].prev = m[start]
    m[start].distance = 0

    // init
    q.init()

    // take each item
    for q.Len() > 0 {
        current := q.popNearest()

        // abort if not way was found to the current vertex
        if current.prev == nil {
            fmt.Printf("No way found to vertex %d. Aborting.\n", current.GetId())
            return
        }

        // go through edges
        for _, edge := range graph.getEdgesOfVertex(current.VertexInterface).All() {
            weight := edge.GetWeight()

            // abort if edge has negative weight
            if weight < 0 {
                fmt.Println("Negative edge weight found. Aborting.")
                return
            }

            neighbour := m[edge.GetOtherVertex(current.VertexInterface)]

            distance := current.distance + weight
            if distance < neighbour.distance {
                // shorter path found
                q.update(neighbour, current, distance)
            }
        }

        // check if end vertex was visited
        if current.VertexInterface == end {
            break
        }

    }

    // print path to end vertex or every path if no end is defined
    if end != nil {
        printShortestPath(m[end], start)
    } else {
        for _, v := range m {
            printShortestPath(v, start)
        }
    }

}

// prints the shortest path and its weight
func printShortestPath(v *shortestPathVertex, start graphLib.VertexInterface) {
    printShortestPathVertex(v, start)
    fmt.Println("=", v.distance)
}

func printShortestPathVertex(v *shortestPathVertex, start graphLib.VertexInterface) {
    if v.VertexInterface != start {
        printShortestPathVertex(v.prev, start)
    }
    fmt.Print(v.GetId(), " ")
}

// shortest path helper vertex
// knows about its previous vertex and its distance to the start
type shortestPathVertex struct {
    graphLib.VertexInterface
    prev     *shortestPathVertex
    distance float64
    index    int
}

// the priority queue for the shortest path
type shortestPathQueue []*shortestPathVertex

func (this shortestPathQueue) Len() int {
    return len(this)
}

func (this shortestPathQueue) Less(i, j int) bool {
    return this[i].distance < this[j].distance
}

func (this shortestPathQueue) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
    this[i].index = i
    this[j].index = j
}

func (this *shortestPathQueue) Push(x interface{}) {
    item := x.(*shortestPathVertex)
    item.index = len(*this)
    *this = append(*this, item)
}

func (this *shortestPathQueue) Pop() interface{} {
    old := *this
    n := len(old)
    item := old[n - 1]
    item.index = -1
    *this = old[0 : n - 1]
    return item
}

func (this *shortestPathQueue) init() {
    heap.Init(this)
}

// update a vertex' distance and previous vertex
func (this *shortestPathQueue) update(item *shortestPathVertex, prev *shortestPathVertex, distance float64) {
    item.prev = prev
    item.distance = distance
    heap.Fix(this, item.index)
}

// returns the nearest vertex and removes it from the queue
func (this *shortestPathQueue) popNearest() *shortestPathVertex {
    return heap.Pop(this).(*shortestPathVertex)
}



// simple wrapper
func (this Graph) ShortestPathsMBF(start, end graphLib.VertexInterface) (float64, []graphLib.VertexInterface, []graphLib.VertexInterface) {
    return ShortestPathsMBF(this, start, end)
}

//
func ShortestPathsMBF(graph Graph, start, end graphLib.VertexInterface) (float64, []graphLib.VertexInterface, []graphLib.VertexInterface) {

    // number of vertices
    num := graph.GetVertices().Count()

    // map to map vertices to the objects that are used here
    m := make([]*shortestPathVertex, num)
    // create objects
    for i, v := range graph.GetVertices().All() {
        m[v.GetPos()] = &shortestPathVertex{
            VertexInterface: v,
            prev: nil,
            distance: math.MaxFloat64,
            index: i,
        }
    }

    edges := graph.GetEdges().All()

    // start
    m[start.GetPos()].prev = m[start.GetPos()]
    m[start.GetPos()].distance = 0

    var changed *shortestPathVertex
    checkAndUpdate := func(u, v *shortestPathVertex, weight float64) {
        if d := u.distance + weight; d < v.distance {
            v.distance = d
            v.prev = u
            changed = v
        }
    }

    // main loop (num times)
    for n := int(num); n >= 0; n-- {
        changed = nil

        // go through each edge
        for _, e := range edges {

            // update v if distance over u is shorter
            u, v := m[e.GetStartVertex().GetPos()], m[e.GetEndVertex().GetPos()]
            if u.prev != nil {
                checkAndUpdate(u, v, e.GetWeight())
            }
            if !graph.IsDirected() && v.prev != nil {
                checkAndUpdate(v, u, e.GetWeight())
            }

        }

        // abort if no more changes were made
        if changed == nil {
            break;
        }

    }

    // check for loop
    if changed != nil {
        // go num times back
        for n := int(num); n >= 0; n-- {
            changed = changed.prev
        }
        return 0.0, nil, buildShortestPath(changed.prev, changed, 1)
    }

    // check if the end vertex was reached and print path(s)
    v := m[end.GetPos()]
    if v.prev == nil {
        // no way found
        return 0.0, nil, nil
    }
    return v.distance, buildShortestPath(v, m[start.GetPos()], 1), nil

}

func buildShortestPath(v, until *shortestPathVertex, depth int) (path []graphLib.VertexInterface) {
    if v == until {
        path = make([]graphLib.VertexInterface, 0, depth)
    } else {
        path = buildShortestPath(v.prev, until, depth + 1)
    }
    path = append(path, v)
    return path
}