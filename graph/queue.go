package graph

import (
    "container/heap"
)

// interface for a queue that returns the nearest vertices
type NearestVertexQueueInterface interface {
    PushVertex(VertexInterface, float64, VertexInterface) *NearestVertexQueueItem
    UpdateVertex(*NearestVertexQueueItem, float64, VertexInterface)
    PopNearestVertex() (VertexInterface, float64, VertexInterface)
    Len() int
    IsEmpty() bool
}

// create a new queue and init it
func NewNearestVertexQueue(length uint) NearestVertexQueueInterface {
    pq := make(nearestVertexQueue, 0, length)
    heap.Init(&pq)
    return &pq
}

// item consists of the vertex and the distance
type NearestVertexQueueItem struct {
    Vertex VertexInterface
    Weight float64
    From   VertexInterface
    index  int
}
type nearestVertexQueue []*NearestVertexQueueItem

// returns the length of the queue
func (this nearestVertexQueue) Len() int {
    return len(this)
}

func (this nearestVertexQueue) Less(i, j int) bool {
    return this[i].Weight < this[j].Weight
}

func (this nearestVertexQueue) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
    this[i].index = i
    this[j].index = j
}

func (this *nearestVertexQueue) Push(x interface{}) {
    x.(*NearestVertexQueueItem).index = len(*this)
    *this = append(*this, x.(*NearestVertexQueueItem))
}

func (this *nearestVertexQueue) Pop() (item interface{}) {
    n := len(*this)
    old := []*NearestVertexQueueItem(*this)
    *this, item = old[0 : n-1], old[n - 1]
    item.(*NearestVertexQueueItem).index = -1
    return
}

// add vertex with the given distance
func (this *nearestVertexQueue) PushVertex(vertex VertexInterface, distance float64, nearest VertexInterface) *NearestVertexQueueItem {
    v := &NearestVertexQueueItem{
        Vertex: vertex,
        Weight: distance,
        From: nearest,
    }
    heap.Push(this, v)
    return v
}

// updates the vertex
func (this *nearestVertexQueue) UpdateVertex(v *NearestVertexQueueItem, distance float64, nearest VertexInterface) {
    v.Weight = distance
    v.From = nearest
    heap.Fix(this, v.index)
}

// returns one of the nearest vertices and its distance
func (this *nearestVertexQueue) PopNearestVertex() (VertexInterface, float64, VertexInterface) {
    if len(*this) > 0 {
        v := heap.Pop(this).(*NearestVertexQueueItem)
        return v.Vertex, v.Weight, v.From
    }
    return nil, -1.0, nil
}

func (this nearestVertexQueue) IsEmpty() bool {
    return len(this) <= 0
}