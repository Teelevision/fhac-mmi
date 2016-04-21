package main

import (
    "github.com/teelevision/fhac-mmi/parser"
    "github.com/teelevision/fhac-mmi/algorithm"
    "fmt"
    "time"
    "flag"
    "errors"
    "github.com/teelevision/fhac-mmi/graph"
    "os"
    "runtime/pprof"
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
    startVertex         *int
    showTime            *bool
    cpuProfile          *string
}

// inits the current config
func initConfig() {
    config.inputFormat = flag.String("f", "list", "input format (matrix|list)")
    config.weights = flag.Bool("w", false, "input list contains weights")
    config.directed = flag.Bool("d", false, "graph is directed")
    config.print = flag.Bool("print", true, "print info")
    config.breadthFirstSearch = flag.Bool("breadth", false, "breadth-first search")
    config.depthFirstSearch = flag.Bool("depth", false, "depth-first search")
    config.connectedComponents = flag.Bool("components", false, "connected components")
    config.prim = flag.Bool("prim", false, "prim minimal spanning tree length")
    config.kruskal = flag.Bool("kruskal", false, "kruskal minimal spanning tree length")
    config.startVertex = flag.Int("start", 0, "start vertex")
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
func parseFile(file string) (*graph.Graph, error) {
    switch *config.inputFormat {
    case "auto":
        fmt.Println("Input format \"auto\" is not yet implemented. Using \"list\".")
        fallthrough
    case "list":
        return parser.ParseEdgesFile(file, *config.weights)
    case "matrix":
        return parser.ParseAdjacencyMatrixFile(file)
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

        // get start vertex
        start := graph.GetVertices().Get(uint(*config.startVertex))

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
            length := graph.PrimLength(start)
            fmt.Println("Length of minimal spanning tree (Prim):", length)
        }

        // kruskal
        if *config.kruskal {
            length := graph.KruskalLength()
            fmt.Println("Length of minimal spanning tree (Kruskal):", length)
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
