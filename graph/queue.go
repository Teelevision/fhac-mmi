package graph

import (
    "container/heap"
)

//***************************
// nearest vertex queue

// interface for a queue that returns the nearest vertices
type NearestVertexQueueInterface interface {
    PushVertex(VertexInterface, float64, VertexInterface) *NearestVertexQueueItem
    UpdatedVertex(*NearestVertexQueueItem)
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
    *this, item = old[0 : n - 1], old[n - 1]
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

// fix the vertex when updated
func (this *nearestVertexQueue) UpdatedVertex(v *NearestVertexQueueItem) {
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


//***************************
// cheapest edge queue

// interface for a queue that returns the cheapest edges
type CheapestEdgeQueueInterface interface {
    PushEdge(EdgeInterface) *CheapestEdgeQueueItem
    UpdatedEdge(*CheapestEdgeQueueItem)
    PopCheapestEdge() EdgeInterface
    Len() int
    IsEmpty() bool
}

// create a new queue and init it
func NewCheapestEdgeQueue(length uint) CheapestEdgeQueueInterface {
    pq := make(cheapestEdgeQueue, 0, length)
    heap.Init(&pq)
    return &pq
}

// item consists of the edge and the index
type CheapestEdgeQueueItem struct {
    EdgeInterface
    index int
}
type cheapestEdgeQueue []*CheapestEdgeQueueItem

// returns the length of the queue
func (this cheapestEdgeQueue) Len() int {
    return len(this)
}

func (this cheapestEdgeQueue) Less(i, j int) bool {
    return this[i].GetWeight() < this[j].GetWeight()
}

func (this cheapestEdgeQueue) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
    this[i].index = i
    this[j].index = j
}

func (this *cheapestEdgeQueue) Push(x interface{}) {
    x.(*CheapestEdgeQueueItem).index = len(*this)
    *this = append(*this, x.(*CheapestEdgeQueueItem))
}

func (this *cheapestEdgeQueue) Pop() (item interface{}) {
    n := len(*this)
    old := []*CheapestEdgeQueueItem(*this)
    *this, item = old[0 : n - 1], old[n - 1]
    item.(*CheapestEdgeQueueItem).index = -1
    return
}

// add edge
func (this *cheapestEdgeQueue) PushEdge(edge EdgeInterface) *CheapestEdgeQueueItem {
    v := &CheapestEdgeQueueItem{EdgeInterface: edge}
    heap.Push(this, v)
    return v
}

// fix the edge when updated
func (this *cheapestEdgeQueue) UpdatedEdge(e *CheapestEdgeQueueItem) {
    heap.Fix(this, e.index)
}

// returns one of the cheapest edges
func (this *cheapestEdgeQueue) PopCheapestEdge() EdgeInterface {
    if len(*this) > 0 {
        v := heap.Pop(this).(*CheapestEdgeQueueItem)
        return v.EdgeInterface
    }
    return nil
}

func (this cheapestEdgeQueue) IsEmpty() bool {
    return len(this) <= 0
}