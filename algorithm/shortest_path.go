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

// returns the length of the hamilton circle calculated by the nearest neighbour algorithm
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

    // function that prints the a path and it's distance
    printPath := func(v *shortestPathVertex) {
        fmt.Print("[", v.distance, "] ")
        for ; v.VertexInterface != start; v = v.prev {
            fmt.Print(v.GetId(), "-")
        }
        fmt.Println(start.GetId())
    }

    // print path to end vertex or every path if no end is defined
    if end != nil {
        printPath(m[end])
    } else {
        for _, v := range m {
            printPath(v)
        }
    }

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