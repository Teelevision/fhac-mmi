package main

import "fmt"
import "github.com/teelevision/fhac-mmi/graph"

func main() {

    g := graph.NewGraph()

    vA := g.NewVertex()
    vB := g.NewVertex()
    eA := g.NewEdge(vA, vB)
    eB := g.NewEdge(vB, vA)
    eC := g.NewEdge(vA, vA)

    fmt.Println(eA)
    fmt.Println(eB)
    fmt.Println(eC)
    fmt.Printf("%v+\n", vA)
    fmt.Printf("%v+\n", vB)

}
