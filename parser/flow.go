package parser

import (
    graphLib "github.com/teelevision/fhac-mmi/graph"
    "io"
    "os"
)

// parses an file containing a flow graph
func ParseFlowFile(file string) (*graphLib.Graph, error) {
    f, _ := os.Open(file)
    graph, err := ParseFlow(f)
    return graph, err
}

type FlowVertex struct {
    graphLib.VertexInterface
    Balance     float64
    FlowBalance float64
}

func (this FlowVertex) Clone() graphLib.VertexInterface {
    return &FlowVertex{
        VertexInterface: this.VertexInterface.Clone(),
        Balance: this.Balance,
    }
}

func (this FlowVertex) GetBalance() float64 {
    return this.Balance
}

func (this FlowVertex) GetFlowBalance() float64 {
    return this.FlowBalance
}

func (this *FlowVertex) SetFlowBalance(balance float64) {
    this.FlowBalance = balance
}

type FlowEdge struct {
    graphLib.EdgeInterface
    Cost float64
    Flow float64
}

func (this FlowEdge) Clone() graphLib.EdgeInterface {
    return &FlowEdge{
        EdgeInterface: this.EdgeInterface.Clone(),
        Cost: this.Cost,
        Flow: this.Flow,
    }
}

func (this FlowEdge) GetCapacity() float64 {
    return this.GetWeight()
}

func (this FlowEdge) GetCost() float64 {
    return this.Cost
}

func (this FlowEdge) GetFlow() float64 {
    return this.Flow
}

func (this *FlowEdge) SetFlow(flow float64) {
    this.Flow = flow;
}


// parses an file containing a flow graph
func ParseFlow(reader io.Reader) (*graphLib.Graph, error) {

    // parse vertices
    graph, vertices, scanner, err := parseHeader(reader)
    if err != nil {
        return graph, err
    }

    // add balance to vertices
    graph = graph.Transform(func(vertex graphLib.VertexInterface) graphLib.VertexInterface {
        balance, err := parseFloat(scanner)
        if err != nil {
            panic(err)
        }
        return &FlowVertex{
            VertexInterface: vertex,
            Balance: balance,
        }
    }, nil)


    // create edges
    costs := make([]float64, 0)
    for {

        // parse start vertex and test if input is empty
        start, err := parseInt(scanner)
        if err != nil && err.Error() == "EOF" {
            break
        } else if err != nil {
            return graph, err
        }

        // parse end vertex
        end, err := parseInt(scanner)
        if err != nil {
            return graph, err
        }

        // parse cost
        cost, err := parseFloat(scanner)
        if err != nil {
            return graph, err
        }
        costs = append(costs, cost)

        // parse capacity
        weight, err := parseFloat(scanner)
        if err != nil {
            return graph, err
        }

        // create edge
        graph.NewWeightedEdge(vertices[start], vertices[end], weight)
    }

    // add cost to edges
    graph = graph.Transform(nil, func(edge graphLib.EdgeInterface) graphLib.EdgeInterface {
        return &FlowEdge{
            EdgeInterface: edge,
            Cost: costs[edge.GetPos()],
        }
    })

    return graph, nil
}