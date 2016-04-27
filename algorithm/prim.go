package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
)

// simple wrapper
func (this Graph) Prim(start graphLib.VertexInterface) (float64, Graph, map[graphLib.VertexInterface]graphLib.VertexInterface) {
    return Prim(this, start)
}

// prim algorithm with result length
func Prim(graph Graph, start graphLib.VertexInterface) (float64, Graph, map[graphLib.VertexInterface]graphLib.VertexInterface) {

    // the number of vertices in our graph
    num := graph.GetVertices().Count()

    // map vertices to primVertex (holds the visited flag, the vertex of the new graph and the queue item)
    vertices := make(map[graphLib.VertexInterface]*graphLib.NearestVertexQueueItem, num)

    // create queue
    q := graphLib.NewNearestVertexQueue(num)

    // keep track of the length of the minimal spanning tree
    length := float64(0)

    // map source to result vertex
    vMap := make(map[graphLib.VertexInterface]graphLib.VertexInterface, num)

    // the result spanning tree
    result := graphLib.CreateNewGraphWithNumVerticesAndNumEdges(false, num, num - 1)
    for _, v := range graph.GetVertices().All() {
        vMap[v] = result.NewVertex()
    }

    // go through queue
    for v, d, n := start, float64(0), graphLib.VertexInterface(nil); v != nil; v, d, n = q.PopNearestVertex() {

        if n != nil {
            // add length
            length += d
            // add edge
            result.NewWeightedEdge(vMap[n], vMap[v], d)
        }

        // vertex is visited
        vertices[v] = nil

        // search all edges (ignore direction of the graph)
        for _, all := range v.GetEdgesFast() {
            for _, edge := range all() {

                // get the vertex that is not v
                neighbourVertex := edge.GetOtherVertex(v)

                // skip origin
                if neighbourVertex == n {
                    continue
                }

                // get the weight between v nd neighbourVertex
                weight := edge.GetWeight()

                // check if we know that vertex already
                if neighbour, ok := vertices[neighbourVertex]; ok {
                    if neighbour != nil && neighbour.Weight > weight {
                        // if it is not visited and we discovered a edge with lower weight
                        neighbour.Weight = weight
                        neighbour.From = v
                        q.UpdatedVertex(neighbour)
                    }
                } else {
                    // undiscovered vertex: add to queue
                    vertices[neighbourVertex] = q.PushVertex(neighbourVertex, weight, v)
                }
            }
        }
    }

    // return the length/weight of the minimal spanning tree and the minimal spanning tree itself
    return length, Graph{result}, vMap
}