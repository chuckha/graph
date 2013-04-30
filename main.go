package main

import (
	"github.com/ChuckHa/graphs/graph"
	"fmt"
)

func main() {
	g := graph.NewGraph("testdata.txt", true)
	fmt.Println(g)
}

