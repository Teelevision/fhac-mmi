package main

import (
    "github.com/teelevision/fhac-mmi/parser"
    "github.com/teelevision/fhac-mmi/algorithm"
    "fmt"
    "time"
    "flag"
    "errors"
    graphLib "github.com/teelevision/fhac-mmi/graph"
    "os"
    "runtime/pprof"
    "runtime"
)

var config struct {
    files               []string
    inputFormat         *string
    weights             *bool
    directed            *bool
    print               *bool
    breadthFirstSearch  *bool
    depthFirstSearch    *bool
    connectedComponents *bool
    prim                *bool
    kruskal             *bool
    nearestNeighbour    *bool
    doubleTree          *bool
    travelingSalesmanBF *bool
    travelingSalesmanBB *bool
    shortestPath        *string
    maxFlow             *bool
    optimalFlow         *string
    startVertex         *int
    endVertex           *int
    showTime            *bool
    cpuProfile          *string
}

// inits the current config
func initConfig() {
    config.inputFormat = flag.String("f", "list", "input format (matrix|list|flow)")
    config.weights = flag.Bool("w", false, "input list contains weights")
    config.directed = flag.Bool("d", false, "graph is directed")
    config.print = flag.Bool("print", true, "print info")
    config.breadthFirstSearch = flag.Bool("breadth", false, "breadth-first search")
    config.depthFirstSearch = flag.Bool("depth", false, "depth-first search")
    config.connectedComponents = flag.Bool("components", false, "connected components")
    config.prim = flag.Bool("prim", false, "prim minimal spanning tree length")
    config.kruskal = flag.Bool("kruskal", false, "kruskal minimal spanning tree length")
    config.nearestNeighbour = flag.Bool("nn", false, "nearest neighbour hamilton circle length")
    config.doubleTree = flag.Bool("dt", false, "double tree hamilton circle length")
    config.travelingSalesmanBF = flag.Bool("tsbf", false, "traveling salesman brute force")
    config.travelingSalesmanBB = flag.Bool("tsbb", false, "traveling salesman branch and bound")
    config.shortestPath = flag.String("sp", "", "shortest path (d|mbf)")
    config.maxFlow = flag.Bool("maxflow", false, "maximum flow")
    config.optimalFlow = flag.String("of", "", "optimal flow (cc|ssp)")
    config.startVertex = flag.Int("start", 0, "start vertex")
    config.endVertex = flag.Int("end", -1, "end vertex")
    config.showTime = flag.Bool("t", false, "show time")
    config.cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")

    flag.Parse()

    // the files are listed last
    config.files = flag.Args()
    if len(config.files) == 0 {
        panic(errors.New("No file given."))
    }
}

// parses the file and returns the graph
func parseFile(file string) (*graphLib.Graph, error) {
    switch *config.inputFormat {
    case "auto":
        fmt.Println("Input format \"auto\" is not yet implemented. Using \"list\".")
        fallthrough
    case "list":
        return parser.ParseEdgesFile(file, *config.weights)
    case "matrix":
        return parser.ParseAdjacencyMatrixFile(file)
    case "flow":
        return parser.ParseFlowFile(file)
    default:
        panic(errors.New(fmt.Sprintf("Unkown input format \"%s\".", *config.inputFormat)))
    }
}

func main() {

    initConfig()

    // cpu profile
    if *config.cpuProfile != "" {
        f, err := os.Create(*config.cpuProfile)
        if err != nil {
            panic(err)
        }
        runtime.SetCPUProfileRate(500)
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()
    }

    // do it all for every file
    for i, file := range config.files {

        // print file path
        fmt.Println("##################")
        fmt.Printf("%d. file: %s\n", i + 1, file)

        startTime := time.Now()

        // parse graph from file
        g, err := parseFile(file)
        if err != nil {
            panic(err)
        }
        g.SetDirected(*config.directed)
        graph := algorithm.Graph{g}

        initTime := time.Now()

        // get start/end vertex
        start := graph.GetVertices().Get(uint(*config.startVertex))
        end := graphLib.VertexInterface(nil)
        if *config.endVertex >= 0 {
            end = graph.GetVertices().Get(uint(*config.endVertex))
        }

        // print info
        if *config.print {
            fmt.Printf("Vertices: %d\n", graph.GetVertices().Count())
            fmt.Printf("Edges: %d\n", graph.GetEdges().Count())
        }

        // do breadth-first search
        if *config.breadthFirstSearch {
            fmt.Print("Breadth-first search:")
            for _, vertex := range graph.BreadthFirstSearch(start) {
                fmt.Printf(" %d", vertex.GetId())
            }
            fmt.Println()
        }

        // do depth-first search
        if *config.depthFirstSearch {
            fmt.Print("Depth-first search:")
            for _, vertex := range graph.DepthFirstSearch(start) {
                fmt.Printf(" %d", vertex.GetId())
            }
            fmt.Println()
        }

        // connected components
        if *config.connectedComponents {
            numB := algorithm.GetNumConnectedComponents(graph, algorithm.BreadthFirstSearch)
            numD := algorithm.GetNumConnectedComponents(graph, algorithm.DepthFirstSearch)
            fmt.Println("Connected components (via breadth-first search):", numB)
            fmt.Println("Connected components (via depth-first search):", numD)
        }

        // prim
        if *config.prim {
            length, _, _ := graph.Prim(start)
            fmt.Println("Length of minimal spanning tree (Prim):", length)
        }

        // kruskal
        if *config.kruskal {
            length := graph.KruskalLength()
            fmt.Println("Length of minimal spanning tree (Kruskal):", length)
        }

        // nearest neighbour
        if *config.nearestNeighbour {
            length := graph.NearestNeighbourHamiltonCircleLength(start)
            fmt.Println("Length of Hamilton circle (Nearest Neighbour):", length)
        }

        // double tree
        if *config.doubleTree {
            tour, length := graph.DoubleTreeHamiltonCircle(algorithm.Prim, start)
            fmt.Print("Length of Hamilton circle (Double Tree[Prim]): ", length, " [")
            for _, v := range tour {
                fmt.Print(" ", v.GetId())
            }
            fmt.Println(" ]")
        }

        // traveling salesman brute force
        if *config.travelingSalesmanBF {
            length := graph.TravelingSalesmanBruteForce(false)
            fmt.Println("Length of shortest Hamilton circle (brute force):", length)
        }

        // traveling salesman brute force
        if *config.travelingSalesmanBB {
            length := graph.TravelingSalesmanBruteForce(true)
            fmt.Println("Length of shortest Hamilton circle (branch and bound):", length)
        }

        // shortest paths
        switch *config.shortestPath {
        case "d":
            fmt.Println("Shortest paths (Dijkstra):")
            graph.ShortestPathsDijkstra(start, end)
        case "mbf":
            e := end
            if e == nil {
                e = graph.GetVertices().Get(graph.GetVertices().Count() - 1)
            }
            fmt.Println("Shortest paths (Moore-Bellman-Ford):")
            distance, path, circle := graph.ShortestPathsMBF(start, e)
            if path == nil && circle == nil {
                fmt.Printf("No way found from %d to %d.\n", start.GetPos(), e.GetPos())
            } else if circle != nil {
                fmt.Print("Negative circle found:")
                for _, v := range circle {
                    fmt.Printf(" %d", v.GetPos())
                }
                fmt.Println()
            } else {
                fmt.Printf("Result (length %f):", distance)
                for _, v := range path {
                    fmt.Printf(" %d", v.GetPos())
                }
                fmt.Println()
            }
        }

        // max flow Edmonds-Karp algorithm
        if *config.maxFlow {
            maxFlow, _ := graph.MaxFlowEdmondsKarp(start, end)
            fmt.Println("Maximum flow (Edmonds-Karp):", maxFlow)
        }

        // optimal flow
        if *config.optimalFlow != "" {
            var usage []float64
            switch *config.optimalFlow {
            case "cc":
                fmt.Println("Optimal flow (Cycle-Canceling):")
                usage = graph.OptimalFlowCycleCancelling()
            case "ssp":
                fmt.Println("Optimal flow (Successive Shortest Path):")
                usage = graph.OptimalFlowSuccessiveShortestPath()
            }
            for i, u := range usage {
                e := graph.GetEdges().GetPos(i)
                fmt.Printf("  %d -> %d: %f\n", e.GetStartVertex().GetPos(), e.GetEndVertex().GetPos(), u)
            }
        }

        endTime := time.Now()

        if *config.showTime {
            fmt.Printf("Duration: total %v | init %v | calc %v\n",
                endTime.Sub(startTime),
                initTime.Sub(startTime),
                endTime.Sub(initTime))
        }

    }
}
