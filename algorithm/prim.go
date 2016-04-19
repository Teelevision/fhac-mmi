package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
)

// simple wrapper
func (this Graph) Prim(start graphLib.VertexInterface) (float64, Graph) {
    return Prim(this, start)
}

// helper struct for the prim algorithm
type primVertex struct {
    // if the vertex was visited
    visited bool
    // the corresponding new vertex in the result graph
    new     graphLib.VertexInterface
    // the item in the queue
    item    *graphLib.NearestVertexQueueItem
}

func Prim(graph Graph, start graphLib.VertexInterface) (float64, Graph) {

    // the number of vertices in our graph
    num := graph.GetVertices().Count()

    // initialize result graph with num vertices and num - 1 edges
    result := graphLib.CreateNewGraphWithNumVerticesAndNumEdges(false, num, num - 1)

    // map vertices to primVertex (holds the visited flag, the vertex of the new graph and the queue item)
    vertices := make(map[graphLib.VertexInterface]*primVertex, num)

    // initialize with the start vertex
    vertices[start] = &primVertex{
        new: result.NewVertexWithId(start.GetId()),
    }

    // create queue
    q := graphLib.NewNearestVertexQueue(num)

    // keep track of the length of the minimal spanning tree
    length := float64(0)

    // go through queue
    for v, d, n := start, float64(0), graphLib.VertexInterface(nil); v != nil; v, d, n = q.PopNearestVertex() {

        // add length
        length += d

        // vertex is visited
        vv := vertices[v]
        vv.visited = true

        // search all edges (ignore direction of the graph)
        for _, edge := range v.GetEdges().All() {

            // get the vertex that is not v
            neighbourVertex := edge.GetOtherVertex(v)

            // get the weight between v nd neighbourVertex
            weight := edge.GetWeight()

            // check if we know that vertex already
            if neighbour, ok := vertices[neighbourVertex]; ok {
                if !neighbour.visited && neighbour.item.Weight > weight {
                    // if it is not visited and we discovered a edge with lower weight
                    q.UpdateVertex(neighbour.item, weight, v)
                }
            } else {
                // undiscovered vertex: add to map
                vertices[neighbourVertex] = &primVertex{
                    // add to result graph
                    new: result.NewVertexWithId(neighbourVertex.GetId()),
                    // add to queue
                    item: q.PushVertex(neighbourVertex, weight, v),
                }
            }
        }

        // add edge to minimal spanning tree / result graph
        if n != nil {
            result.NewWeightedEdge(vertices[n].new, vv.new, d)
        }
    }

    // return the length/weight of the minimal spanning tree and the resulting graph
    return length, Graph{result}
}