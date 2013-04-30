package graph

import (
	"errors"
	"os"
	"fmt"
)

const maxVerticies = 100
const maxDegree = 50

type Graph struct {
	Edges              *[maxVerticies + 1][maxDegree]int
	Degree             *[maxVerticies + 1]int
	NVerticies, NEdges int
	Directed           bool
}

func NewGraph(filename string, directed bool) *Graph {
	f, _ := os.Open(filename)
	reader := NewReader(f)
	nVerticies, nEdges := reader.ReadLine()
	g := Graph{
		Edges:      new([maxVerticies + 1][maxDegree]int),
		Degree:     new([maxVerticies + 1]int),
		NVerticies: nVerticies,
		NEdges:     nEdges,
		Directed:   directed,
	}

	for i := 0; i < nEdges; i++ {
		left, right := reader.ReadLine()
		g.InsertEdge(left, right, g.Directed)
	}
	return &g
}

func (g *Graph) InsertEdge(x, y int, directed bool) error {
	if g.Degree[x] > maxDegree {
		return errors.New("Insertion exceeds maxdegree")
	}
	g.Edges[x][g.Degree[x]] = y
	g.Degree[x] += 1
	if !directed {
		g.InsertEdge(y, x, true)
	}
	return nil
}

func (g *Graph) String() string {
	s := ""
	// Count from one
	for i := 1; i <= g.NVerticies; i ++ {
		s += fmt.Sprintf("%d:", i)
		for j := 0; j < g.Degree[i]; j++ {
			s += fmt.Sprintf(" %d", g.Edges[i][j])
		}
		s += "\n"
	}
	return s
}
