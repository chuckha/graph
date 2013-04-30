# Graph Algorithms

Graph algorithm implementations written in Go.

# Reading a graph

This graph format assumes the first line is the number of verticies and the number of edges.
The following number-of-edges lines are edges from vertex `a` to vertex `b`:

### Example

```
4 4 // There are 4 verticies and 4 edges
1 2 // There is an edge from vertex 1 to vertex 2
2 3 // There is an edge from vertex 2 to vertex 3
3 4 // ...etc
4 1
```

## Usage

The `graph` package defines a reader to read this type of file which will be totally hidden.
All graphs will come from text files.
Graphs can be directed or undirected

    // This creates a new graph from the data in testdata.txt and says that it is a directed graph.
    // A false value will produce an undirected graph and add all inverse verticies automatically.
    g := graph.NewGraph("testdata.txt", true)
    fmt.Println(g)


## Test data

There is a test graph in `testdata.txt`.


