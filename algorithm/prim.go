package algorithm

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
)

// simple wrapper
func (this Graph) Prim(start graphLib.VertexInterface) (float64, Graph) {
    return Prim(this, start)
}

type primVertex struct {
    visited bool
    new     graphLib.VertexInterface
    item    *graphLib.NearestVertexQueueItem
}

func Prim(graph Graph, start graphLib.VertexInterface) (float64, Graph) {

    result := graphLib.CreateNewGraph(false)
    num := graph.GetVertices().Count()
    vertices := make(map[graphLib.VertexInterface]*primVertex, num)
    vertices[start] = &primVertex{
        new: result.NewVertexWithId(start.GetId()),
    }
    q := graphLib.NewNearestVertexQueue(num)

    length := float64(0)

    for v, d, n := start, float64(0), graphLib.VertexInterface(nil); v != nil; v, d, n = q.PopNearestVertex() {

        length += d
        vv := vertices[v]
        vv.visited = true

        for _, edge := range v.GetEdges().All() {
            neighbourVertex := edge.GetOtherVertex(v)
            weight := edge.GetWeight()
            if neighbour, ok := vertices[neighbourVertex]; ok {
                if !neighbour.visited && neighbour.item.Weight > weight {
                    q.UpdateVertex(neighbour.item, weight, v)
                }
            } else {
                vertices[neighbourVertex] = &primVertex{
                    new: result.NewVertexWithId(neighbourVertex.GetId()),
                    item: q.PushVertex(neighbourVertex, weight, v),
                }
            }
        }

        if n != nil {
            result.NewWeightedEdge(vertices[n].new, vv.new, d)
        }
    }

    return length, Graph{result}
}