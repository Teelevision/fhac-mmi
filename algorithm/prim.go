package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
)

// simple wrapper
func (this Graph) PrimLength(start graphLib.VertexInterface) float64 {
    return PrimLength(this, start)
}

// helper struct for the prim algorithm
type primVertex struct {
    // if the vertex was visited
    visited bool
    // the item in the queue
    item    *graphLib.NearestVertexQueueItem
}

// prim algorithm with result length
func PrimLength(graph Graph, start graphLib.VertexInterface) float64 {

    // the number of vertices in our graph
    num := graph.GetVertices().Count()

    // map vertices to primVertex (holds the visited flag, the vertex of the new graph and the queue item)
    vertices := make(map[graphLib.VertexInterface]*primVertex, num)

    // initialize with the start vertex
    vertices[start] = &primVertex{}

    // create queue
    q := graphLib.NewNearestVertexQueue(num)

    // keep track of the length of the minimal spanning tree
    length := float64(0)

    // go through queue
    for v, d := start, float64(0); v != nil; v, d, _ = q.PopNearestVertex() {

        // add length
        length += d

        // vertex is visited
        vv := vertices[v]
        vv.visited = true

        // search all edges (ignore direction of the graph)
        for _, all := range v.GetEdgesFast() {
            for _, edge := range all() {

                // get the vertex that is not v
                neighbourVertex := edge.GetOtherVertex(v)

                // get the weight between v nd neighbourVertex
                weight := edge.GetWeight()

                // check if we know that vertex already
                if neighbour, ok := vertices[neighbourVertex]; ok {
                    if !neighbour.visited && neighbour.item.Weight > weight {
                        // if it is not visited and we discovered a edge with lower weight
                        neighbour.item.Weight = weight
                        neighbour.item.From = v
                        q.UpdatedVertex(neighbour.item)
                    }
                } else {
                    // undiscovered vertex: add to map
                    vertices[neighbourVertex] = &primVertex{
                        // add to queue
                        item: q.PushVertex(neighbourVertex, weight, v),
                    }
                }
            }
        }
    }

    // return the length/weight of the minimal spanning tree
    return length
}